// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Group) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Group, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *Group, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Group, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.ID, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original group [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.TimeCurrentP()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $7", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Group) error {
	if len(models) == 0 {
		return nil
	}
	lo.ForEach(models, func(model *Group, _ int) {
		model.Created = util.TimeCurrent()
		model.Updated = util.TimeCurrentP()
	})
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columnsQuoted, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *Group, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, 1, logger, id)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
