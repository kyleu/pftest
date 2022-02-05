// Package app - $PF_IGNORE$
package app

import (
	"context"

	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/lib/audit"

	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/group"
	"github.com/kyleu/pftest/app/history"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/app/version"
	"github.com/kyleu/pftest/queries/migrations"
	"github.com/pkg/errors"
)

type Services struct {
	Basic     *basic.Service
	Audited   *audited.Service
	Timestamp *timestamp.Service
	Softdel   *softdel.Service
	History   *history.Service
	Version   *version.Service
	Group     *group.Service
	MixedCase *mixedcase.Service
	Trouble   *trouble.Service
	Capital   *capital.Service
	Audit     *audit.Service
}

func NewServices(ctx context.Context, st *State) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, st.Logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to migrate database")
	}

	aud := audit.NewService(st.DB, st.Logger)

	return &Services{
		Basic:     basic.NewService(st.DB, st.Logger),
		Audited:   audited.NewService(st.DB, aud, st.Logger),
		Timestamp: timestamp.NewService(st.DB, st.Logger),
		Softdel:   softdel.NewService(st.DB, st.Logger),
		History:   history.NewService(st.DB, st.Logger),
		Version:   version.NewService(st.DB, st.Logger),
		Group:     group.NewService(st.DB, st.Logger),
		MixedCase: mixedcase.NewService(st.DB, st.Logger),
		Trouble:   trouble.NewService(st.DB, st.Logger),
		Capital:   capital.NewService(st.DB, st.Logger),
		Audit:     audit.NewService(st.DB, st.Logger),
	}, nil
}
