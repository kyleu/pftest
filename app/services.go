// Package app - $PF_IGNORE$
package app

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/group"
	"github.com/kyleu/pftest/app/history"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/app/version"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	Basic     *basic.Service
	Relation  *relation.Service
	Reference *reference.Service
	Audited   *audited.Service
	Timestamp *timestamp.Service
	Softdel   *softdel.Service
	History   *history.Service
	Version   *version.Service
	Group     *group.Service
	MixedCase *mixedcase.Service
	Trouble   *trouble.Service
	Capital   *capital.Service
	Path      *path.Service
	Audit     *audit.Service
}

type x struct{}

func (_ *x) Hello() string {
	return "Hi!"
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to migrate database")
	}

	aud := audit.NewService(st.DB, logger)

	sch := `type Query {
		hello: String!
	}`
	resolver := &x{}

	err = st.GraphQL.RegisterStringSchema("foo", "Foo", sch, resolver)
	if err != nil {
		return nil, errors.Wrap(err, "unable to register GraphQL schema")
	}
	err = st.GraphQL.RegisterStringSchema("bar", "Bar", sch, resolver)
	if err != nil {
		return nil, errors.Wrap(err, "unable to register GraphQL schema")
	}

	return &Services{
		Basic:     basic.NewService(st.DB),
		Relation:  relation.NewService(st.DB),
		Reference: reference.NewService(st.DB),
		Audited:   audited.NewService(st.DB, aud),
		Timestamp: timestamp.NewService(st.DB),
		Softdel:   softdel.NewService(st.DB),
		History:   history.NewService(st.DB),
		Version:   version.NewService(st.DB),
		Group:     group.NewService(st.DB),
		MixedCase: mixedcase.NewService(st.DB),
		Trouble:   trouble.NewService(st.DB),
		Capital:   capital.NewService(st.DB),
		Path:      path.NewService(st.DB),
		Audit:     aud,
	}, nil
}

func (s *Services) Close(_ context.Context, logger util.Logger) error {
	return nil
}
