package app

import (
	"context"

	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	CoreServices
	GeneratedServices
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		logger.Errorf("unable to migrate database: %v", err)
	}

	aud := audit.NewService(st.DB, logger)

	core := initCoreServices(ctx, st, aud, logger)
	gen := initGeneratedServices(ctx, st, aud, logger)

	return &Services{CoreServices: core, GeneratedServices: gen}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
