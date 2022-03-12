// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/pftest

import (
	"os"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/cmd"
	"github.com/kyleu/pftest/app/lib/log"
)

var (
	version = "0.0.0" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	logger, err := cmd.Run(&app.BuildInfo{Version: version, Commit: commit, Date: date})
	if err != nil {
		const msg = "exiting due to error"
		if logger == nil {
			println(log.Red.Add(err.Error())) //nolint
			println(log.Red.Add(msg))         //nolint
		} else {
			logger.Error(err)
			logger.Error(msg)
		}
		os.Exit(1)
	}
}
