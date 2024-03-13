package app

import (
	"context"
	"encoding/json"

	"github.com/kyleu/pftest/app/gql"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/lib/exec"
	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/lib/help"
	"github.com/kyleu/pftest/app/lib/notebook"
	"github.com/kyleu/pftest/app/lib/scripting"
	"github.com/kyleu/pftest/app/lib/websocket"
	"github.com/kyleu/pftest/app/user"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	GeneratedServices

	Audit    *audit.Service
	Exec     *exec.Service
	Script   *scripting.Service
	User     *user.Service
	Help     *help.Service
	Socket   *websocket.Service
	Notebook *notebook.Service
	Schema   *gql.Schema
	Har      *har.Service
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
		GeneratedServices: initGeneratedServices(ctx, st.DB, st.DBRead, aud, logger),

		Audit:    aud,
		Exec:     exec.NewService(),
		Script:   scripting.NewService(st.Files, "scripts"),
		User:     user.NewService(st.Files, logger),
		Help:     help.NewService(logger),
		Socket:   sock,
		Notebook: notebook.NewService(),
		Schema:   schema,
		Har:      har.NewService(st.Files),
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
