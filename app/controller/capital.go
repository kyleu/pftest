// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vcapital"
)

func CapitalList(rc *fasthttp.RequestCtx) {
	Act("capital.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("capital", nil, ps.Logger).Sanitize("capital")
		ret, err := as.Services.Capital.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Capitals", ret)
		page := &vcapital.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "capital")
	})
}

func CapitalDetail(rc *fasthttp.RequestCtx) {
	Act("capital.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Capital)", ret)

		return Render(rc, as, &vcapital.Detail{Model: ret}, ps, "capital", ret.TitleString()+"**star")
	})
}

func CapitalCreateForm(rc *fasthttp.RequestCtx) {
	Act("capital.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &capital.Capital{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = capital.Random()
		}
		ps.SetTitleAndData("Create [Capital]", ret)
		ps.Data = ret
		return Render(rc, as, &vcapital.Edit{Model: ret, IsNew: true}, ps, "capital", "Create")
	})
}

func CapitalRandom(rc *fasthttp.RequestCtx) {
	Act("capital.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Capital.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Capital")
		}
		return ret.WebPath(), nil
	})
}

func CapitalCreate(rc *fasthttp.RequestCtx) {
	Act("capital.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		err = as.Services.Capital.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Capital")
		}
		msg := fmt.Sprintf("Capital [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func CapitalEditForm(rc *fasthttp.RequestCtx) {
	Act("capital.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vcapital.Edit{Model: ret}, ps, "capital", ret.String())
	})
}

func CapitalEdit(rc *fasthttp.RequestCtx) {
	Act("capital.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := capitalFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		frm.ID = ret.ID
		err = as.Services.Capital.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Capital [%s]", frm.String())
		}
		msg := fmt.Sprintf("Capital [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func CapitalDelete(rc *fasthttp.RequestCtx) {
	Act("capital.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Capital.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete capital [%s]", ret.String())
		}
		msg := fmt.Sprintf("Capital [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/capital", rc, ps)
	})
}

func capitalFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*capital.Capital, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	return as.Services.Capital.Get(ps.Context, nil, idArg, ps.Logger)
}

func capitalFromForm(rc *fasthttp.RequestCtx, setPK bool) (*capital.Capital, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return capital.FromMap(frm, setPK)
}
