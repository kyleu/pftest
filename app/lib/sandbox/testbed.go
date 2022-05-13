// Package sandbox $PF_IGNORE$
package sandbox

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/util"
)

var testbed = &Sandbox{Key: "testbed", Title: "Testbed", Icon: "star", Run: onTestbed}

func onTestbed(ctx context.Context, st *app.State, logger util.Logger) (any, error) {
	ret := util.ValueMap{"status": "ok"}
	return ret, nil
}
