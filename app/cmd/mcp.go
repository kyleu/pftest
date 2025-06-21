package cmd

import (
	"context"

	"github.com/muesli/coral"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/log"
	"github.com/kyleu/pftest/app/lib/mcpserver"
	"github.com/kyleu/pftest/app/util"
)

func mcpCmd() *coral.Command {
	f := func(*coral.Command, []string) error { return runMCP(context.Background()) }
	ret := &coral.Command{Use: "mcp", Short: "Handles Model Context Protocol requests", RunE: f}
	return ret
}

func runMCP(ctx context.Context) error {
	l, err := log.InitDevLogging(log.GetLevel(zap.FatalLevel))
	if err != nil {
		return errors.Wrap(err, "error initializing logging")
	}
	util.RootLogger = l.Sugar()
	if err = initIfNeeded(); err != nil {
		return errors.Wrap(err, "error initializing application")
	}
	st, err := app.Bootstrap(_buildInfo, _flags.ConfigDir, _flags.Port, false, util.RootLogger)
	if err != nil {
		return err
	}
	mcp, err := mcpserver.NewServer(ctx, st, util.RootLogger)
	if err != nil {
		return err
	}
	if err := mcp.ServeCLI(ctx); err != nil {
		return err
	}
	return nil
}
