package controller

import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
)

// Initialize app-specific system dependencies.
func initApp(_ *app.State, _ util.Logger) {
}

// Configure app-specific data for each request.
func initAppRequest(*app.State, *cutil.PageState) error {
	return nil
}
