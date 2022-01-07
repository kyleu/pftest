package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/kyleu/pftest/app/util"
)

const keyAll = "all"

func allCmd() *cobra.Command {
	short := fmt.Sprintf("Starts the main http server on port %d and the marketing site on port %d", util.AppPort, util.AppPort+1)
	f := func(*cobra.Command, []string) error { return allF() }
	ret := &cobra.Command{Use: keyAll, Short: short, RunE: f}
	return ret
}

func allF() error {
	if err := initIfNeeded(); err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	go func() {
		if err := startSite(_flags.Clone(_flags.Port + 1)); err != nil {
			_logger.Errorf("unable to start marketing site: %+v", err)
		}
	}()
	return startServer(_flags)
}
