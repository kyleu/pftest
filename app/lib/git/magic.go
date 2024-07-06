package git

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Magic(ctx context.Context, message string, dryRun bool, logger util.Logger) (*Result, error) {
	args, err := s.magicArgsFor(ctx, message, dryRun, logger)
	if err != nil {
		return nil, err
	}

	var logs []string
	add := func(msg string, args ...any) {
		logs = append(logs, fmt.Sprintf(msg, args...))
	}

	switch {
	case args.Ahead == 0 && args.Behind == 0:
		err := s.onClean(args, add)
		if err != nil {
			return args.Result, err
		}
	case args.Ahead == 0 && args.Behind > 0:
		err := s.onBehind(args, add)
		if err != nil {
			return args.Result, err
		}
	case args.Ahead > 0 && args.Behind == 0:
		err := s.onAhead(args, add)
		if err != nil {
			return args.Result, err
		}
	case args.Ahead > 0 && args.Behind > 0:
		args.Result.Status = "conflicting commits"
		return args.Result, errors.New("encountered conflicting commits")
	default:
		return args.Result, errors.New("invalid git state")
	}

	args.Result.Data["logs"] = logs
	return args.Result, nil
}
