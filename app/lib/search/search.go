// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/search/result"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/util"
)

type Provider func(context.Context, *app.State, *Params, *zap.SugaredLogger) (result.Results, error)

func Search(ctx context.Context, as *app.State, params *Params, logger *zap.SugaredLogger) (result.Results, []error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "search", logger)
	defer span.Complete()

	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	// $PF_SECTION_END(search_functions)$
	// $PF_INJECT_START(codegen)$
	basicFunc := func(ctx context.Context, as *app.State, params *Params, logger *zap.SugaredLogger) (result.Results, error) {
		models, err := as.Services.Basic.Search(ctx, params.Q, nil, params.PS.Get("basic", nil, as.Logger))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("basic", m.String(), m.WebPath(), m.String(), "star", m, params.Q))
		}
		return res, nil
	}
	referenceFunc := func(ctx context.Context, as *app.State, params *Params, logger *zap.SugaredLogger) (result.Results, error) {
		models, err := as.Services.Reference.Search(ctx, params.Q, nil, params.PS.Get("reference", nil, as.Logger))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("reference", m.String(), m.WebPath(), m.String(), "star", m, params.Q))
		}
		return res, nil
	}
	auditedFunc := func(ctx context.Context, as *app.State, params *Params, logger *zap.SugaredLogger) (result.Results, error) {
		models, err := as.Services.Audited.Search(ctx, params.Q, nil, params.PS.Get("audited", nil, as.Logger))
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("audited", m.String(), m.WebPath(), m.String(), "star", m, params.Q))
		}
		return res, nil
	}
	allProviders = append(allProviders, basicFunc, referenceFunc, auditedFunc)
	// $PF_INJECT_END(codegen)$
	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	params.Q = strings.TrimSpace(params.Q)

	results, errs := util.AsyncCollect(allProviders, func(item Provider) (result.Results, error) {
		return item(ctx, as, params, logger)
	})

	ret := make(result.Results, 0, len(results)*len(results))
	for _, x := range results {
		ret = append(ret, x...)
	}

	ret.Sort()
	return ret, errs
}
