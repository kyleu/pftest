package search

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/search/result"
	"github.com/kyleu/pftest/app/util"
)

func generatedSearch() []Provider {
	auditedFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("audited", logger).WithLimit(5)
		return as.Services.Audited.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	basicFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("basic", logger).WithLimit(5)
		return as.Services.Basic.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	pathFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("path", logger).WithLimit(5)
		return as.Services.Path.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	referenceFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("reference", logger).WithLimit(5)
		return as.Services.Reference.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	relationFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("relation", logger).WithLimit(5)
		return as.Services.Relation.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	return []Provider{auditedFunc, basicFunc, pathFunc, referenceFunc, relationFunc}
}
