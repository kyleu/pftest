package git

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/util"
)

func (s *Service) gitStash(ctx context.Context, path string, logger util.Logger) (string, error) {
	out, err := gitCmd(ctx, "stash", path, logger)
	if err != nil {
		if isNoRepo(err) {
			return "", nil
		}
		return "", errors.Wrap(err, "unable to apply stash")
	}
	return out, nil
}

func (s *Service) gitStashPop(ctx context.Context, path string, logger util.Logger) (string, error) {
	out, err := gitCmd(ctx, "stash pop", path, logger)
	if err != nil {
		if isNoRepo(err) {
			return "", nil
		}
		return "", errors.Wrap(err, "unable to pop stash")
	}
	return out, nil
}