// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *Trouble, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Trouble, logger util.Logger) error {
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"from\" = $7 and \"where\" = $8", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.From, model.Where)
	_, err := s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"from", "where"}, columnsQuoted, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *Trouble, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) SaveChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, logger util.Logger, models ...*Trouble) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			logger.Infof("saving troubles [%d-%d]", idx*chunkSize, ((idx+1)*chunkSize)-1)
		}
		if err := s.Save(ctx, tx, logger, chunk...); err != nil {
			return err
		}
	}
	return nil
}

// Delete doesn't actually delete, it only sets [delete].
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, from string, where []string, logger util.Logger) error {
	cols := []string{"delete"}
	q := database.SQLUpdate(tableQuoted, cols, defaultWC(len(cols)), s.db.Placeholder())
	_, err := s.db.Update(ctx, q, tx, 1, logger, util.TimeCurrent(), from, where)
	return err
}

// Delete doesn't actually delete, it only sets [delete].
func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	cols := []string{"delete"}
	q := database.SQLUpdate(tableQuoted, cols, wc, s.db.Placeholder())
	_, err := s.db.Update(ctx, q, tx, expected, logger, append([]any{util.TimeCurrent()}, values...)...)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"delete\" is null"
}
