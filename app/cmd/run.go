package cmd

import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return _logger, err
	}
	return _logger, nil
}
