// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vbasic"
)

func BasicList(rc *fasthttp.RequestCtx) {
	Act("basic.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("basic", nil, ps.Logger).Sanitize("basic")
		var ret basic.Basics
		var err error
		if q == "" {
			ret, err = as.Services.Basic.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Basic.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Basics"
		ps.Data = ret
		page := &vbasic.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "basic")
	})
}

func BasicDetail(rc *fasthttp.RequestCtx) {
	Act("basic.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Basic)"
		ps.Data = ret

		relRelationsByBasicIDPrms := ps.Params.Get("relation", nil, ps.Logger).Sanitize("relation")
		relRelationsByBasicID, err := as.Services.Relation.GetByBasicID(ps.Context, nil, ret.ID, relRelationsByBasicIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child relations")
		}
		return Render(rc, as, &vbasic.Detail{
			Model:  ret,
			Params: ps.Params,

			RelRelationsByBasicID: relRelationsByBasicID,
		}, ps, "basic", ret.String())
	})
}

func BasicCreateForm(rc *fasthttp.RequestCtx) {
	Act("basic.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &basic.Basic{}
		ps.Title = "Create [Basic]"
		ps.Data = ret
		return Render(rc, as, &vbasic.Edit{Model: ret, IsNew: true}, ps, "basic", "Create")
	})
}

func BasicCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("basic.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := basic.Random()
		ps.Title = "Create Random Basic"
		ps.Data = ret
		return Render(rc, as, &vbasic.Edit{Model: ret, IsNew: true}, ps, "basic", "Create")
	})
}

func BasicCreate(rc *fasthttp.RequestCtx) {
	Act("basic.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		err = as.Services.Basic.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Basic")
		}
		msg := fmt.Sprintf("Basic [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func BasicEditForm(rc *fasthttp.RequestCtx) {
	Act("basic.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vbasic.Edit{Model: ret}, ps, "basic", ret.String())
	})
}

func BasicEdit(rc *fasthttp.RequestCtx) {
	Act("basic.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := basicFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		frm.ID = ret.ID
		err = as.Services.Basic.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Basic [%s]", frm.String())
		}
		msg := fmt.Sprintf("Basic [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func BasicDelete(rc *fasthttp.RequestCtx) {
	Act("basic.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Basic.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete basic [%s]", ret.String())
		}
		msg := fmt.Sprintf("Basic [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/basic", rc, ps)
	})
}

func basicFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*basic.Basic, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Basic.Get(ps.Context, nil, idArg, ps.Logger)
}

func basicFromForm(rc *fasthttp.RequestCtx, setPK bool) (*basic.Basic, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return basic.FromMap(frm, setPK)
}
