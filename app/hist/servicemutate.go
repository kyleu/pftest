// Content managed by Project Forge, see [projectforge.md] for details.
package hist

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Hist) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Placeholder())
	vals := make([]any, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Hist, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.ID, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original hist [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()

	_, hErr := s.SaveHistory(ctx, tx, curr, model, logger)
	if hErr != nil {
		return errors.Wrap(hErr, "unable to save history")
	}
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateIfNeeded(ctx context.Context, tx *sqlx.Tx, model *Hist, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.ID, logger)
	if curr == nil || err != nil {
		return s.Create(ctx, tx, logger, model)
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()

	h, hErr := s.SaveHistory(ctx, tx, curr, model, logger)
	if hErr != nil {
		return errors.Wrap(hErr, "unable to save history")
	}
	if h == nil || len(h.Changes) == 0 {
		return nil
	}
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", s.db.Placeholder())
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Hist) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columnsQuoted, s.db.Placeholder())
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
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
