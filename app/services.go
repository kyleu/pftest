package app

import (
	"context"
	"encoding/json"

	"github.com/kyleu/pftest/app/gql"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database/migrate"
	"github.com/kyleu/pftest/app/lib/help"
	"github.com/kyleu/pftest/app/lib/notebook"
	"github.com/kyleu/pftest/app/lib/scripting"
	"github.com/kyleu/pftest/app/lib/websocket"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/queries/migrations"
)

type Services struct {
	CoreServices
	GeneratedServices

	Script   *scripting.Service
	Help     *help.Service
	Socket   *websocket.Service
	Notebook *notebook.Service
	Schema   *gql.Schema
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		logger.Errorf("unable to migrate database: %v", err)
	}

	aud := audit.NewService(st.DB, logger)

	core := initCoreServices(ctx, st, aud, logger)
	core.Socket.Close()
	core.Socket = websocket.NewService(nil, socketHandler, nil)
	gen := initGeneratedServices(ctx, st, aud, logger)

	schema, err := gql.NewSchema(st.GraphQL)
	if err != nil {
		return nil, err
	}

	return &Services{CoreServices: core, GeneratedServices: gen, Schema: schema}, nil
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
