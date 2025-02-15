package git

import (
	"context"
	"strings"

	"github.com/kyleu/pftest/app/util"
)

func gitBranch(ctx context.Context, path string, logger util.Logger) string {
	out, err := gitCmd(ctx, "branch --show-current", path, logger)
	if err != nil {
		if isNotRepo(err) {
			return "norepo"
		}
		return "error: " + err.Error()
	}
	return strings.TrimSpace(out)
}
