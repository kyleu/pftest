package basic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/pftest/app/lib/database"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Basic) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), "")
	vals := make([]interface{}, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Basic) error {
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $4", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, ret := s.db.Update(ctx, q, tx, 1, data...)
	return ret
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Basic) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columns, "")
	var data []interface{}
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error {
	q := database.SQLDelete(tableQuoted, "\"id\" = $1")
	_, err := s.db.Delete(ctx, q, tx, 1, id)
	return err
}
