// Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/util"
)

func buildDefaultAppState(flags *Flags, logger util.Logger) (*app.State, error) {
	f := filesystem.NewFileSystem(flags.ConfigDir, logger)
	telemetryDisabled := util.GetEnvBool("disable_telemetry", false)
	st, err := app.NewState(flags.Debug, _buildInfo, f, !telemetryDisabled, logger)
	if err != nil {
		return nil, err
	}

	ctx, span, logger := telemetry.StartSpan(context.Background(), "app:init", logger)
	defer span.Complete()

	db, err := database.OpenDefaultPostgres(ctx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open database")
	}
	st.DB = db

	svcs, err := app.NewServices(ctx, st)
	if err != nil {
		return nil, errors.Wrap(err, "error creating services")
	}
	st.Services = svcs

	return st, nil
}
