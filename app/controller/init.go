package controller

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/gql"
	"github.com/kyleu/pftest/app/lib/mcpserver"
	"github.com/kyleu/pftest/app/util"
)

// Initialize app-specific system dependencies.
func initApp(ctx context.Context, as *app.State, logger util.Logger) error {
	_, err := gql.NewSchema(as, logger)
	if err != nil {
		return err
	}
	_, err = mcpserver.NewServer(ctx, as, logger)
	if err != nil {
		logger.Errorf("unable to start MCP server: %v", err)
	}
	return nil
}

// Configure app-specific data for each request.
func initAppRequest(*app.State, *cutil.PageState) error {
	return nil
}
