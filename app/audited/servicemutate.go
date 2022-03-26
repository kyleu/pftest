// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/pftest/app/lib/database"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Audited) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), "")
	vals := make([]any, 0, len(models)*len(columnsQuoted))
	for _, arg := range models {
		vals = append(vals, arg.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Audited) error {
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"id\" = $3", "")
	data := model.ToData()
	data = append(data, model.ID)
	_, err := s.db.Update(ctx, q, tx, 1, s.logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Audited) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columns, "")
	var data []any
	for _, model := range models {
		data = append(data, model.ToData()...)
	}
	return s.db.Insert(ctx, q, tx, s.logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error {
	q := database.SQLDelete(tableQuoted, defaultWC)
	_, err := s.db.Delete(ctx, q, tx, 1, s.logger, id)
	return err
}
