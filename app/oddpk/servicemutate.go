package oddpk

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*OddPK) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Type)
	vals := lo.FlatMap(models, func(arg *OddPK, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) CreateChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, progress *util.Progress, logger util.Logger, models ...*OddPK) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			count := ((idx + 1) * chunkSize) - 1
			if len(models) < count {
				count = len(models)
			}
			logger.Infof("creating odd pks [%d-%d]", idx*chunkSize, count)
		}
		if err := s.Create(ctx, tx, logger, chunk...); err != nil {
			return err
		}
		progress.Increment(len(chunk), logger)
	}
	return nil
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *OddPK, logger util.Logger) error {
	q := database.SQLUpdate(tableQuoted, columnsQuoted, "\"project\" = $4 and \"path\" = $5", s.db.Type)
	data := model.ToData()
	data = append(data, model.Project, model.Path)
	_, err := s.db.Update(ctx, q, tx, 1, logger, data...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*OddPK) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"project", "path"}, columnsQuoted, s.db.Type)
	data := lo.FlatMap(models, func(model *OddPK, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) SaveChunked(ctx context.Context, tx *sqlx.Tx, chunkSize int, progress *util.Progress, logger util.Logger, models ...*OddPK) error {
	for idx, chunk := range lo.Chunk(models, chunkSize) {
		if logger != nil {
			count := ((idx + 1) * chunkSize) - 1
			if len(models) < count {
				count = len(models)
			}
			logger.Infof("saving odd pks [%d-%d]", idx*chunkSize, count)
		}
		if err := s.Save(ctx, tx, logger, chunk...); err != nil {
			return err
		}
		progress.Increment(len(chunk), logger)
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, project uuid.UUID, path string, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC(0), s.db.Type)
	_, err := s.db.Delete(ctx, q, tx, 1, logger, project, path)
	return err
}

func (s *Service) DeleteWhere(ctx context.Context, tx *sqlx.Tx, wc string, expected int, logger util.Logger, values ...any) error {
	q := database.SQLDelete(tableQuoted, wc, s.db.Type)
	_, err := s.db.Delete(ctx, q, tx, expected, logger, values...)
	return err
}
