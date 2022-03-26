// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Version) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentRevisions(ctx, tx, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Revision = revs[model.String()] + 1
		model.Updated = util.NowPointer()
	}

	err = s.upsertCore(ctx, tx, models...)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Version) error {
	revs, err := s.getCurrentRevisions(ctx, tx, model)
	if err != nil {
		return err
	}
	model.Revision = revs[model.String()] + 1
	curr, err := s.Get(ctx, tx, model.ID)
	if err != nil {
		return errors.Wrapf(err, "can't get original version [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()

	err = s.upsertCore(ctx, tx, model)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Version) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentRevisions(ctx, tx, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		curr, e := s.Get(ctx, tx, model.ID)
		if e == nil && curr != nil {
			model.Created = curr.Created
		} else {
			model.Created = time.Now()
		}
		model.Revision = revs[model.String()] + 1
		model.Updated = util.NowPointer()
	}

	err = s.upsertCore(ctx, tx, models...)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) upsertCore(ctx context.Context, tx *sqlx.Tx, models ...*Version) error {
	conflicts := util.StringArrayQuoted([]string{"id"})
	q := database.SQLUpsert(tableQuoted, columnsCore, len(models), conflicts, columnsCore, "")
	data := make([]any, 0, len(columnsCore)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataCore()...)
	}
	_, err := s.db.Update(ctx, q, tx, 1, s.logger, data...)
	return err
}

func (s *Service) insertRevision(ctx context.Context, tx *sqlx.Tx, models ...*Version) error {
	q := database.SQLInsert(tableRevisionQuoted, columnsRevision, len(models), "")
	data := make([]any, 0, len(columnsRevision)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataRevision()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(tableQuoted, defaultWC)
	_, err := s.db.Delete(ctx, q, tx, 1, s.logger, id)
	return err
}
