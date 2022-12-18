// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/search/result"
	"github.com/kyleu/pftest/app/util"
)

//nolint:gocognit
func generatedSearch() []Provider {
	auditedFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Audited.Search(ctx, params.Q, nil, params.PS.Get("audited", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("audited", m.String(), m.WebPath(), m.String(), "star", m, m, params.Q))
		}
		return res, nil
	}
	basicFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Basic.Search(ctx, params.Q, nil, params.PS.Get("basic", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("basic", m.String(), m.WebPath(), m.String(), "star", m, m, params.Q))
		}
		return res, nil
	}
	pathFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Path.Search(ctx, params.Q, nil, params.PS.Get("path", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("path", m.String(), m.WebPath(), m.String(), "star", m, m, params.Q))
		}
		return res, nil
	}
	referenceFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Reference.Search(ctx, params.Q, nil, params.PS.Get("reference", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("reference", m.String(), m.WebPath(), m.String(), "star", m, m, params.Q))
		}
		return res, nil
	}
	relationFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		models, err := as.Services.Relation.Search(ctx, params.Q, nil, params.PS.Get("relation", nil, logger), logger)
		if err != nil {
			return nil, err
		}
		res := make(result.Results, 0, len(models))
		for _, m := range models {
			res = append(res, result.NewResult("relation", m.String(), m.WebPath(), m.String(), "star", m, m, params.Q))
		}
		return res, nil
	}
	return []Provider{auditedFunc, basicFunc, pathFunc, referenceFunc, relationFunc}
}
