// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*History) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), "")
	vals := make([]any, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *History) error {
	curr, err := s.Get(ctx, tx, model.ID)
	if err != nil {
		return errors.Wrapf(err, "can't get original history [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()

	_, hErr := s.SaveHistory(ctx, tx, curr, model)
	if hErr != nil {
		return errors.Wrap(hErr, "unable to save history")
	}
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, s.logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*History) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		curr, e := s.Get(ctx, tx, model.ID)
		if e == nil && curr != nil {
			model.Created = curr.Created
		} else {
			model.Created = time.Now()
		}
		model.Updated = util.NowPointer()

		_, hErr := s.SaveHistory(ctx, tx, curr, model)
		if hErr != nil {
			return errors.Wrap(hErr, "unable to save history")
		}
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columns, "")
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0))
	_, err := s.db.Delete(ctx, q, tx, 1, s.logger, id)
	return err
}
