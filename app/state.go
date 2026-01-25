package app

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/auth"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/lib/graphql"
	"github.com/kyleu/pftest/app/lib/log"
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

func NewState(ctx context.Context, debug bool, bi *BuildInfo, f filesystem.FileLoader, enableTelemetry bool, port uint16, logger util.Logger) (*State, error) {
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

	_ = telemetry.InitializeIfNeeded(ctx, enableTelemetry, bi.Version, logger)

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

func (s *State) Close(ctx context.Context, logger util.Logger) error {
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

func (s *State) User(ctx context.Context, id uuid.UUID, logger util.Logger) (*user.User, error) {
	if s == nil || s.Services == nil || s.Services.User == nil {
		return nil, nil
	}
	return s.Services.User.Get(ctx, nil, id, logger)
}

func Bootstrap(ctx context.Context, bi *BuildInfo, cfgDir string, port uint16, debug bool, logger util.Logger) (*State, error) {
	fs, err := filesystem.NewFileSystem(cfgDir, false, "")
	if err != nil {
		return nil, err
	}

	telemetryDisabled := util.GetEnvBoolAny(false, "disable_telemetry", "telemetry_disabled")
	st, err := NewState(ctx, debug, bi, fs, !telemetryDisabled, port, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(ctx, "app:init", logger)
	defer span.Complete()
	t := util.TimerStart()

	db, err := database.OpenDefaultPostgres(ctx, logger)
	if err != nil {
		logger.Errorf("unable to open default database: %+v", err)
	}
	st.DB = db
	roSuffix := "_readonly"
	rKey := util.AppKey + roSuffix
	if x := util.GetEnv("read_db_host", ""); x != "" {
		paramsR := database.PostgresParamsFromEnv(rKey, rKey, "read_")
		logger.Infof("using [%s:%s] for read-only database pool", paramsR.Host, paramsR.Database)
		st.DBRead, err = database.OpenPostgresDatabase(ctx, rKey, paramsR, logger)
	} else {
		paramsR := database.PostgresParamsFromEnv(rKey, util.AppKey, "")
		if strings.HasSuffix(paramsR.Database, roSuffix) {
			paramsR.Database = util.AppKey
		}
		logger.Infof("using default database as read-only database pool")
		st.DBRead, err = database.OpenPostgresDatabase(ctx, rKey, paramsR, logger)
	}
	if err != nil {
		logger.Errorf("unable to open default read-only database: %v", err)
	}
	st.DBRead.ReadOnly = true
	svcs, err := NewServices(ctx, st, logger)
	if err != nil {
		return nil, err
	}
	logger.Debugf("created app state in [%s]", util.MicrosToMillis(t.End()))
	st.Services = svcs

	return st, nil
}

func BootstrapRunDefault[T any](ctx context.Context, bi *BuildInfo, fn func(as *State, logger util.Logger) (T, error)) (T, error) {
	logger, _ := log.InitLogging(false)
	as, err := Bootstrap(ctx, bi, util.ConfigDir, 0, false, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	ret, err := fn(as, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	err = as.Close(ctx, logger)
	if err != nil {
		var dflt T
		return dflt, err
	}
	return ret, nil
}
