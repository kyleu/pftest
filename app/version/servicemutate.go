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

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Version) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentRevisions(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Revision = revs[model.String()] + 1
		model.Updated = util.NowPointer()
	}

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Version, logger util.Logger) error {
	revs, err := s.getCurrentRevisions(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	model.Revision = revs[model.String()] + 1
	curr, err := s.Get(ctx, tx, model.ID, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original version [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()

	err = s.upsertCore(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Version) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentRevisions(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Revision = revs[model.String()] + 1
		model.Updated = util.NowPointer()
	}

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertRevision(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) upsertCore(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Version) error {
	conflicts := util.StringArrayQuoted([]string{"id"})
	q := database.SQLUpsert(tableQuoted, columnsCore, len(models), conflicts, columnsCore, "")
	data := make([]any, 0, len(columnsCore)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataCore()...)
	}
	_, err := s.db.Update(ctx, q, tx, 1, logger, data...)
	return err
}

func (s *Service) insertRevision(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Version) error {
	q := database.SQLInsert(tableRevisionQuoted, columnsRevision, len(models), "")
	data := make([]any, 0, len(columnsRevision)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataRevision()...)
	}
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0))
	_, err := s.db.Delete(ctx, q, tx, 1, logger, id)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc)
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
