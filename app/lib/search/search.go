// Content managed by Project Forge, see [projectforge.md] for details.
package search

import (
	"context"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/search/result"
)

type Provider func(context.Context, *app.State, *Params) (result.Results, error)

func Search(ctx context.Context, as *app.State, params *Params) (result.Results, []error) {
	if params.Q == "" {
		return nil, nil
	}
	var allProviders []Provider
	// $PF_SECTION_START(search_functions)$
	// $PF_SECTION_END(search_functions)$
	// $PF_INJECT_START(codegen)$
	basicFunc := func(ctx context.Context, as *app.State, params *Params) (result.Results, error) {
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
	auditedFunc := func(ctx context.Context, as *app.State, params *Params) (result.Results, error) {
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
	allProviders = append(allProviders, basicFunc, auditedFunc)
	// $PF_INJECT_END(codegen)$

	if len(allProviders) == 0 {
		return nil, []error{errors.New("no search providers configured")}
	}

	ret := result.Results{}
	var errs []error
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(allProviders))
	params.Q = strings.TrimSpace(params.Q)

	for _, p := range allProviders {
		f := p
		go func() {
			res, err := f(ctx, as, params)
			mu.Lock()
			if err != nil {
				errs = append(errs, err)
			}
			ret = append(ret, res...)
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	ret.Sort()
	return ret, errs
}
