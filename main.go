package main // import github.com/kyleu/pftest

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/cmd"
)

var (
	version = "0.0.0" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
