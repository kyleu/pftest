package version

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

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
		model.Revision = revs[model.String()] + 1
		model.Created = time.Now()
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
		model.Revision = revs[model.String()] + 1
		model.Created = time.Now()
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
	data := make([]interface{}, 0, len(columnsCore)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataCore()...)
	}
	_, err := s.db.Update(ctx, q, tx, 1, data...)
	return err
}

func (s *Service) insertRevision(ctx context.Context, tx *sqlx.Tx, models ...*Version) error {
	q := database.SQLInsert(tableRevisionQuoted, columnsRevision, len(models), "")
	data := make([]interface{}, 0, len(columnsRevision)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataRevision()...)
	}
	return s.db.Insert(ctx, q, tx, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(tableQuoted, "\"id\" = $1")
	_, err := s.db.Delete(ctx, q, tx, 1, id)
	return err
}
