package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/auth"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/lib/graphql"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/theme"
	"github.com/kyleu/pftest/app/user"
	"github.com/kyleu/pftest/app/util"
)

var once sync.Once

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == util.KeyUnknown {
		return b.Version
	}
	d, _ := util.TimeFromJS(b.Date)
	return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Auth      *auth.Service
	DB        *database.Service
	DBRead    *database.Service
	GraphQL   *graphql.Service
	Themes    *theme.Service
	Services  *Services
	Started   time.Time
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, port uint16, logger util.Logger) (*State, error) {
	var loadLocationError error
	once.Do(func() {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			loadLocationError = err
			return
		}
		time.Local = loc
	})
	if loadLocationError != nil {
		return nil, loadLocationError
	}

	_ = telemetry.InitializeIfNeeded(enableTelemetry, bi.Version, logger)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      auth.NewService("", port, logger),
		GraphQL:   graphql.NewService(),
		Themes:    theme.NewService(f),
		Started:   util.TimeCurrent(),
	}, nil
}

func (s State) Close(ctx context.Context, logger util.Logger) error {
	defer func() { _ = telemetry.Close(ctx) }()
	if err := s.DB.Close(); err != nil {
		logger.Errorf("error closing database: %+v", err)
	}
	if err := s.DBRead.Close(); err != nil {
		logger.Errorf("error closing read-only database: %+v", err)
	}
	if err := s.GraphQL.Close(); err != nil {
		logger.Errorf("error closing GraphQL service: %+v", err)
	}
	return s.Services.Close(ctx, logger)
}

func (s State) User(ctx context.Context, id uuid.UUID, logger util.Logger) (*user.User, error) {
	if s.Services == nil || s.Services.User == nil {
		return nil, nil
	}
	return s.Services.User.Get(ctx, nil, id, logger)
}
