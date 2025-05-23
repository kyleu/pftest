package app

import (
	"context"

	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/lib/git"
	"github.com/kyleu/pftest/app/lib/grep"
	"github.com/kyleu/pftest/app/lib/proxy"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	CoreServices
	GeneratedServices

	Proxy *proxy.Service
	Git   *git.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		logger.Errorf("unable to migrate database: %v", err)
	}

	aud := audit.NewService(st.DB, logger)
	prx := proxy.NewService("/proxy", nil)
	g := git.NewService(util.AppKey, ".")
	_ = grep.NewRequest("", "", true) // to typecheck package

	core := initCoreServices(ctx, st, aud, logger)
	gen := initGeneratedServices(ctx, st, aud, logger)

	return &Services{CoreServices: core, GeneratedServices: gen, Proxy: prx, Git: g}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
