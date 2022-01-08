package group

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

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
	q := database.SQLInsert(table, columns, len(models), "")
	vals := make([]interface{}, 0, len(models)*len(columns))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Group) error {
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(table, columns, "id = $7", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, ret := s.db.Update(ctx, q, tx, 1, data...)
	return ret
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Group) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLUpsert(table, columns, len(models), []string{"id"}, columns, "")
	var data []interface{}
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(table, "id = $1")
	_, err := s.db.Delete(ctx, q, tx, 1, id)
	return err
}
