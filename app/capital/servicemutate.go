// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, models ...*Capital) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentVersions(ctx, tx, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Version = revs[model.String()] + 1
	}

	err = s.upsertCore(ctx, tx, models...)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Capital) error {
	revs, err := s.getCurrentVersions(ctx, tx, model)
	if err != nil {
		return err
	}
	model.Version = revs[model.String()] + 1

	err = s.upsertCore(ctx, tx, model)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, models ...*Capital) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentVersions(ctx, tx, models...)
	if err != nil {
		return err
	}
	for _, model := range models {
		model.Version = revs[model.String()] + 1
	}

	err = s.upsertCore(ctx, tx, models...)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) upsertCore(ctx context.Context, tx *sqlx.Tx, models ...*Capital) error {
	conflicts := util.StringArrayQuoted([]string{"ID"})
	q := database.SQLUpsert(tableQuoted, columnsCore, len(models), conflicts, columnsCore, "")
	data := make([]any, 0, len(columnsCore)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataCore()...)
	}
	_, err := s.db.Update(ctx, q, tx, 1, data...)
	return err
}

func (s *Service) insertVersion(ctx context.Context, tx *sqlx.Tx, models ...*Capital) error {
	q := database.SQLInsert(tableVersionQuoted, columnsVersion, len(models), "")
	data := make([]any, 0, len(columnsVersion)*len(models))
	for _, model := range models {
		data = append(data, model.ToDataVersion()...)
	}
	return s.db.Insert(ctx, q, tx, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string) error {
	q := database.SQLDelete(tableQuoted, defaultWC)
	_, err := s.db.Delete(ctx, q, tx, 1, id)
	return err
}
