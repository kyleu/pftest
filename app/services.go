package app

import (
	"context"
	"encoding/json"

	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/gql"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/lib/exec"
	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/lib/help"
	"github.com/kyleu/pftest/app/lib/scripting"
	"github.com/kyleu/pftest/app/lib/websocket"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/seed"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/app/user"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	Basic     *basic.Service
	Relation  *relation.Service
	Reference *reference.Service
	Audited   *audited.Service
	Seed      *seed.Service
	Timestamp *timestamp.Service
	Softdel   *softdel.Service
	MixedCase *mixedcase.Service
	Trouble   *trouble.Service
	Capital   *capital.Service
	Path      *path.Service
	Audit     *audit.Service
	Exec      *exec.Service
	Script    *scripting.Service
	User      *user.Service
	Help      *help.Service
	Socket    *websocket.Service
	Schema    *gql.Schema
	Har       *har.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		logger.Errorf("unable to migrate database: %v", err)
	}

	aud := audit.NewService(st.DB, logger)
	sock := websocket.NewService(nil, socketHandler, nil)
	schema, err := gql.NewSchema(st.GraphQL)
	if err != nil {
		return nil, err
	}

	return &Services{
		Basic:     basic.NewService(st.DB, st.DBRead),
		Relation:  relation.NewService(st.DB, st.DBRead),
		Reference: reference.NewService(st.DB, st.DBRead),
		Audited:   audited.NewService(st.DB, aud),
		Seed:      seed.NewService(st.DB, st.DBRead),
		Timestamp: timestamp.NewService(st.DB, st.DBRead),
		Softdel:   softdel.NewService(st.DB, st.DBRead),
		MixedCase: mixedcase.NewService(st.DB, st.DBRead),
		Trouble:   trouble.NewService(st.DB, st.DBRead),
		Capital:   capital.NewService(st.DB, st.DBRead),
		Path:      path.NewService(st.DB, st.DBRead),
		Audit:     aud,
		Exec:      exec.NewService(),
		Script:    scripting.NewService(st.Files, "scripts"),
		User:      user.NewService(st.Files, logger),
		Help:      help.NewService(logger),
		Socket:    sock,
		Schema:    schema,
		Har:       har.NewService(st.Files),
	}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}

func socketHandler(_ context.Context, s *websocket.Service, c *websocket.Connection, _ string, cmd string, _ json.RawMessage, logger util.Logger) error {
	switch cmd {
	case "connect":
		_, err := s.Join(c.ID, "tap", logger)
		if err != nil {
			return err
		}
	default:
		logger.Error("unhandled command [" + cmd + "]")
	}
	return nil
}
