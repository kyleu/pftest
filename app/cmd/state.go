// Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/util"
)

func buildDefaultAppState(flags *Flags, logger util.Logger) (*app.State, error) {
	fs, err := filesystem.NewFileSystem(flags.ConfigDir, false, "")
	if err != nil {
		return nil, err
	}

	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := app.NewState(flags.Debug, _buildInfo, fs, !telemetryDisabled, flags.Port, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(context.Background(), "app:init", logger)
	defer span.Complete()
	t := util.TimerStart()

	db, err := database.OpenDefaultPostgres(ctx, logger)
	if err != nil {
		return nil, err
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
		return nil, errors.Wrap(err, "unable to open read-only database")
	}
	st.DBRead.ReadOnly = true
	svcs, err := app.NewServices(ctx, st, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error creating services")
	}
	logger.Debugf("created app state in [%s]", util.MicrosToMillis(t.End()))
	st.Services = svcs

	return st, nil
}
