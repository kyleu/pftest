// Content managed by Project Forge, see [projectforge.md] for details.
package migrate

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

const (
	migrationTable    = "migration"
	migrationTableSQL = `create table if not exists "migration" (
  "idx" int not null primary key,
  "title" varchar(1024) not null,
  "src" text not null,
  "created" timestamp not null default now()
);`
)

func ListMigrations(ctx context.Context, s *database.Service, params *filter.Params, logger *zap.SugaredLogger) Migrations {
	params = filter.ParamsWithDefaultOrdering(migrationTable, params, &filter.Ordering{Column: "created", Asc: false})
	var dtos []migrationDTO
	q := database.SQLSelect("*", migrationTable, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.Select(ctx, &dtos, q, nil)
	if err != nil {
		logger.Errorf("error retrieving migrations: %+v", err)
		return nil
	}
	return toMigrations(dtos)
}

func createMigrationTableIfNeeded(ctx context.Context, s *database.Service, logger *zap.SugaredLogger) error {
	q := database.SQLSelectSimple("count(*) as x", migrationTable)
	_, err := s.SingleInt(ctx, q, nil)
	if err != nil {
		logger.Info("first run, creating migration table")
		_, err = s.Exec(ctx, migrationTableSQL, nil, -1)
		if err != nil {
			return errors.Wrapf(err, "error creating migration table: %+v", err)
		}
	}
	return nil
}

func getMigrationByIdx(ctx context.Context, s *database.Service, idx int, logger *zap.SugaredLogger) *Migration {
	dto := &migrationDTO{}
	q := database.SQLSelectSimple("*", "migration", "idx = $1")
	err := s.Get(ctx, dto, q, nil, idx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		logger.Errorf("error getting migration by idx [%v]: %+v", idx, err)
		return nil
	}
	return dto.toMigration()
}

func removeMigrationByIdx(ctx context.Context, s *database.Service, idx int) error {
	q := database.SQLDelete("migration", "idx = $1")
	_, err := s.Delete(ctx, q, nil, 1, idx)
	if err != nil {
		return errors.Wrap(err, "error removing migration")
	}
	return nil
}

func newMigration(ctx context.Context, s *database.Service, e *Migration) error {
	q := database.SQLInsert("migration", []string{"idx", "title", "src"}, 1, s.Type.Placeholder)
	return s.Insert(ctx, q, nil, e.Idx, e.Title, e.Src)
}

func maxMigrationIdx(ctx context.Context, s *database.Service, logger *zap.SugaredLogger) int {
	q := database.SQLSelectSimple("max(idx) as x", "migration")
	max, err := s.SingleInt(ctx, q, nil)
	if err != nil {
		logger.Errorf("error getting migrations: %+v", err)
		return -1
	}
	return int(max)
}
