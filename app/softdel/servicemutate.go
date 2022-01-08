package softdel

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Softdel) error {
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

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Softdel) error {
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(table, columns, "id = $5", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, ret := s.db.Update(ctx, q, tx, 1, data...)
	return ret
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Softdel) error {
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

// Delete doesn't actually delete, it only sets [deleted]
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLUpdate(table, []string{"deleted"}, "id = $2", "")
	_, err := s.db.Update(ctx, q, tx, 1, time.Now(), id)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"deleted\" is null"
}
