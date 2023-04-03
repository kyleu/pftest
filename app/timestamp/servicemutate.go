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

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Timestamp) error {
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

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Timestamp, logger util.Logger) error {
	curr, err := s.Get(ctx, tx, model.ID, true, logger)
	if err != nil {
		return errors.Wrapf(err, "can't get original timestamp [%s]", model.String())
	}
	model.Created = curr.Created
	model.Updated = util.NowPointer()
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $5", s.db.Placeholder())
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

// Delete doesn't actually delete, it only sets [deleted].
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) error {
	cols := []string{"deleted"}
	q := database.SQLUpdate(tableQuoted, cols, defaultWC(len(cols)), s.db.Placeholder())
	_, err := s.db.Update(ctx, q, tx, 1, logger, time.Now(), id)
	return err
}

// Delete doesn't actually delete, it only sets [deleted].
func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	cols := []string{"deleted"}
	q := database.SQLUpdate(tableQuoted, cols, wc, s.db.Placeholder())
	_, err := s.db.Update(ctx, q, tx, expected, logger, append([]any{time.Now()}, values...)...)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"deleted\" is null"
}
