// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentSelectcols(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Selectcol = revs[model.String()] + 1
	}

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertSelectcol(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Trouble, logger util.Logger) error {
	revs, err := s.getCurrentSelectcols(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	model.Selectcol = revs[model.String()] + 1

	err = s.upsertCore(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	err = s.insertSelectcol(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateIfNeeded(ctx context.Context, tx *sqlx.Tx, model *Trouble, logger util.Logger) error {
	revs, err := s.getCurrentSelectcols(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	model.Selectcol = revs[model.String()] + 1

	err = s.upsertCore(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	err = s.insertSelectcol(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentSelectcols(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Selectcol = revs[model.String()] + 1
	}

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertSelectcol(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) upsertCore(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	conflicts := util.StringArrayQuoted([]string{"from", "where"})
	q := database.SQLUpsert(tableQuoted, columnsCore, len(models), conflicts, columnsCore, "")
	data := make([]any, 0, len(columnsCore)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataCore()...)
	}
	_, err := s.db.Update(ctx, q, tx, 1, logger, data...)
	return err
}

func (s *Service) insertSelectcol(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) error {
	q := database.SQLInsert(tableSelectcolQuoted, columnsSelectcol, len(models), "")
	data := make([]any, 0, len(columnsSelectcol)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataSelectcol()...)
	}
	return s.db.Insert(ctx, q, tx, logger, data...)
}

// Delete doesn't actually delete, it only sets [delete].
func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, from string, where int, logger util.Logger) error {
	cols := []string{"delete"}
	q := database.SQLUpdate(tableQuoted, cols, defaultWC(len(cols)), "")
	_, err := s.db.Update(ctx, q, tx, 1, logger, time.Now(), from, where)
	return err
}

// Delete doesn't actually delete, it only sets [delete].
func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	cols := []string{"delete"}
	q := database.SQLUpdate(tableQuoted, cols, wc, "")
	_, err := s.db.Update(ctx, q, tx, expected, logger, append([]any{time.Now()}, values...)...)
	return err
}

func addDeletedClause(wc string, includeDeleted bool) string {
	if includeDeleted {
		return wc
	}
	return wc + " and \"delete\" is null"
}
