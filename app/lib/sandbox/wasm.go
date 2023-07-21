package sandbox

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/util"
)

var wasm = &Sandbox{Key: "wasm", Title: "WASM", Icon: "gift", Run: onWASM}

func onWASM(_ context.Context, _ *app.State, _ util.Logger) (any, error) {
	ret := util.ValueMap{"status": "ok"}
	return ret, nil
}
