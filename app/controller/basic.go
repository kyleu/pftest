// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vbasic"
)

const basicDefaultTitle = "Basics"

func BasicList(rc *fasthttp.RequestCtx) {
	act("basic.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = basicDefaultTitle
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("basic", nil, ps.Logger).Sanitize("basic")
		ret, err := as.Services.Basic.List(ps.Context, nil, prms)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vbasic.List{Models: ret, Params: params}, ps, "basic")
	})
}

func BasicDetail(rc *fasthttp.RequestCtx) {
	act("basic.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		relationPrms := params.Get("relation", nil, ps.Logger).Sanitize("relation")
		relationsByBasicID, err := as.Services.Relation.GetByBasicID(ps.Context, nil, ret.ID, relationPrms)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child relations")
		}
		return render(rc, as, &vbasic.Detail{
			Model:              ret,
			Params:             params,
			RelationsByBasicID: relationsByBasicID,
		}, ps, "basic", ret.String())
	})
}

func BasicCreateForm(rc *fasthttp.RequestCtx) {
	act("basic.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &basic.Basic{}
		ps.Title = "Create [Basic]"
		ps.Data = ret
		return render(rc, as, &vbasic.Edit{Model: ret, IsNew: true}, ps, "basic", "Create")
	})
}

func BasicCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("basic.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := basic.Random()
		ps.Title = "Create Random [Basic]"
		ps.Data = ret
		return render(rc, as, &vbasic.Edit{Model: ret, IsNew: true}, ps, "basic", "Create")
	})
}

func BasicCreate(rc *fasthttp.RequestCtx) {
	act("basic.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		err = as.Services.Basic.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Basic")
		}
		msg := fmt.Sprintf("Basic [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func BasicEditForm(rc *fasthttp.RequestCtx) {
	act("basic.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vbasic.Edit{Model: ret}, ps, "basic", ret.String())
	})
}

func BasicEdit(rc *fasthttp.RequestCtx) {
	act("basic.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := basicFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		frm.ID = ret.ID
		err = as.Services.Basic.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Basic [%s]", frm.String())
		}
		msg := fmt.Sprintf("Basic [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func BasicDelete(rc *fasthttp.RequestCtx) {
	act("basic.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Basic.Delete(ps.Context, nil, ret.ID)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete basic [%s]", ret.String())
		}
		msg := fmt.Sprintf("Basic [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/basic", rc, ps)
	})
}

func basicFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*basic.Basic, error) {
	idArgStr, err := RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Basic.Get(ps.Context, nil, idArg)
}

func basicFromForm(rc *fasthttp.RequestCtx, setPK bool) (*basic.Basic, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return basic.FromMap(frm, setPK)
}
