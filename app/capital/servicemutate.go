// Package capital - Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentVersions(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	lo.ForEach(models, func(model *Capital, _ int) {
		model.Version = revs[model.String()] + 1
	})

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Capital, logger util.Logger) error {
	revs, err := s.getCurrentVersions(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	model.Version = revs[model.String()] + 1

	err = s.upsertCore(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateIfNeeded(ctx context.Context, tx *sqlx.Tx, model *Capital, logger util.Logger) error {
	revs, err := s.getCurrentVersions(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	model.Version = revs[model.String()] + 1

	err = s.upsertCore(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, logger, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) error {
	if len(models) == 0 {
		return nil
	}
	revs, err := s.getCurrentVersions(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	lo.ForEach(models, func(model *Capital, _ int) {
		model.Version = revs[model.String()] + 1
	})

	err = s.upsertCore(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	err = s.insertVersion(ctx, tx, logger, models...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) upsertCore(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) error {
	conflicts := util.StringArrayQuoted([]string{"ID"})
	q := database.SQLUpsert(tableQuoted, columnsCore, len(models), conflicts, columnsCore, s.db.Placeholder())
	data := lo.FlatMap(models, func(model *Capital, _ int) []any {
		return model.ToDataCore()
	})
	_, err := s.db.Update(ctx, q, tx, 1, logger, data...)
	return err
}

func (s *Service) insertVersion(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) error {
	q := database.SQLInsert(tableVersionQuoted, columnsVersion, len(models), s.db.Placeholder())
	data := lo.FlatMap(models, func(model *Capital, _ int) []any {
		return model.ToDataVersion()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, 1, logger, id)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Placeholder())
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
