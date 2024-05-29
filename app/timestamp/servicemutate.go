// Package timestamp - Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Timestamp) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Timestamp, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Type)
	vals := lo.FlatMap(models, func(arg *Timestamp, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) CreateChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, logger util.Logger, models ...*Timestamp) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			count := ((idx + 1) * chunkSize) - 1
			if len(models) < count {
				count = len(models)
			}
			logger.Infof("creating timestamps [%d-%d]", idx*chunkSize, count)
		}
		if err := s.Create(ctx, tx, logger, chunk...); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Timestamp, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.ID, true, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original timestamp [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.TimeCurrentP()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", s.db.Type)
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Timestamp) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Timestamp, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columnsQuoted, s.db.Type)
	data := lo.FlatMap(models, func(model *Timestamp, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) SaveChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, logger util.Logger, models ...*Timestamp) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			count := ((idx + 1) * chunkSize) - 1
			if len(models) < count {
				count = len(models)
			}
			logger.Infof("saving timestamps [%d-%d]", idx*chunkSize, count)
		}
		if err := s.Save(ctx, tx, logger, chunk...); err != nil {
			return err
		}
	}
	return nil
}

// Delete doesn't actually delete, it only sets [deleted].
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) error {
	cols := []string{"deleted"}
	q := database.SQLUpdate(tableQuoted, cols, defaultWC(len(cols)), s.db.Type)
	_, err := s.db.Update(ctx, q, tx, 1, logger, util.TimeCurrent(), id)
	return err
}

// Delete doesn't actually delete, it only sets [deleted].
func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	cols := []string{"deleted"}
	q := database.SQLUpdate(tableQuoted, cols, wc, s.db.Type)
	_, err := s.db.Update(ctx, q, tx, expected, logger, append([]any{util.TimeCurrent()}, values...)...)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"deleted\" is null"
}
