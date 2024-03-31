// Package search - Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/lib/search/result"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/util"
)

func generatedSearch() []Provider {
	auditedFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("audited", logger).WithLimit(5)
		models, err := as.Services.Audited.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *audited.Audited, _ int) *result.Result {
			return result.NewResult("audited", m.String(), m.WebPath(), m.TitleString(), "star", m, m, params.Q)
		}), nil
	}
	basicFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("basic", logger).WithLimit(5)
		models, err := as.Services.Basic.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *basic.Basic, _ int) *result.Result {
			return result.NewResult("basic", m.String(), m.WebPath(), m.TitleString(), "star", m, m, params.Q)
		}), nil
	}
	pathFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("path", logger).WithLimit(5)
		models, err := as.Services.Path.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *path.Path, _ int) *result.Result {
			return result.NewResult("path", m.String(), m.WebPath(), m.TitleString(), "star", m, m, params.Q)
		}), nil
	}
	referenceFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("reference", logger).WithLimit(5)
		models, err := as.Services.Reference.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *reference.Reference, _ int) *result.Result {
			return result.NewResult("reference", m.String(), m.WebPath(), m.TitleString(), "star", m, m, params.Q)
		}), nil
	}
	relationFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		if !page.Admin {
			return nil, nil
		}
		prm := params.PS.Sanitized("relation", logger).WithLimit(5)
		models, err := as.Services.Relation.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *relation.Relation, _ int) *result.Result {
			return result.NewResult("relation", m.String(), m.WebPath(), m.TitleString(), "star", m, m, params.Q)
		}), nil
	}
	return []Provider{auditedFunc, basicFunc, pathFunc, referenceFunc, relationFunc}
}
