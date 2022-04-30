// Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Timestamp) error {
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

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Timestamp) error {
	curr, err := s.Get(ctx, tx, model.ID, true)
	if err != nil {
		return errors.Wrapf(err, "can't get original timestamp [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, err = s.db.Update(ctx, q, tx, 1, s.logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Timestamp) error {
	if len(models) == 0 {
		return nil
	}
	for _, model := range models {
		model.Created = time.Now()
		model.Updated = util.NowPointer()
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columnsQuoted, "")
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, data...)
}

// Delete doesn't actually delete, it only sets [deleted].
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	cols := []string{"deleted"}
	q := database.SQLUpdate(tableQuoted, cols, defaultWC(len(cols)), "")
	_, err := s.db.Update(ctx, q, tx, 1, s.logger, time.Now(), id)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"deleted\" is null"
}
