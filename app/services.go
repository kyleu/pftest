// Package app $PF_IGNORE$
package app

import (
	"context"

	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/queries/migrations"
	"github.com/pkg/errors"
)

type Services struct {
	Basic *basic.Service
}

func NewServices(ctx context.Context, st *State) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, st.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to migrate database")
	}

	b := basic.NewService(st.DB, st.Logger)
	return &Services{Basic: b}, nil
}
