// Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"
	"fmt"
	"time"

	"github.com/kyleu/pftest/app/lib/auth"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/theme"
	"github.com/kyleu/pftest/app/util"
)

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func (b *BuildInfo) String() string {
	if b.Date == "unknown" {
	} else {
		d, _ := util.TimeFromJS(b.Date)
		return fmt.Sprintf("%s (%s)", b.Version, util.TimeToYMD(d))
	}
	return b.Version
}

type State struct {
	Debug     bool
	BuildInfo *BuildInfo
	Files     filesystem.FileLoader
	Auth      *auth.Service
	DB        *database.Service
	Themes    *theme.Service
	Logger    util.Logger
	Services  *Services
	Started   time.Time
}

func (s State) Close(ctx context.Context, logger util.Logger) error {
	if err := s.DB.Close(); err != nil {
		logger.Errorf("error closing database: %+v", err)
	}
	return s.Services.Close(ctx, logger)
}

func NewState(debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, logger util.Logger) (*State, error) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	time.Local = loc

	_ = telemetry.InitializeIfNeeded(enableTelemetry, bi.Version, logger)
	as := auth.NewService("", logger)
	ts := theme.NewService(f)

	return &State{
		Debug:     debug,
		BuildInfo: bi,
		Files:     f,
		Auth:      as,
		Themes:    ts,
		Logger:    logger,
		Started:   time.Now(),
	}, nil
}
