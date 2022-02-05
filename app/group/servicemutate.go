// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Group) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), "")
	vals := make([]interface{}, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Group) error {
	curr, err := s.Get(ctx, tx, model.ID)
	if err != nil {
		return errors.Wrapf(err, "can't get original group [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $7", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Group) error {
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
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columns, "")
	var data []interface{}
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(tableQuoted, defaultWC)
	_, err := s.db.Delete(ctx, q, tx, 1, id)
	return err
}
