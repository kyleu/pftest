// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vmixedcase"
)

func MixedCaseList(rc *fasthttp.RequestCtx) {
	Act("mixedcase.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("mixedcase", nil, ps.Logger).Sanitize("mixedcase")
		ret, err := as.Services.MixedCase.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Mixed Cases", ret)
		page := &vmixedcase.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "mixedcase")
	})
}

func MixedCaseDetail(rc *fasthttp.RequestCtx) {
	Act("mixedcase.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Mixed Case)", ret)

		return Render(rc, as, &vmixedcase.Detail{Model: ret}, ps, "mixedcase", ret.TitleString()+"**star")
	})
}

func MixedCaseCreateForm(rc *fasthttp.RequestCtx) {
	Act("mixedcase.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &mixedcase.MixedCase{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = mixedcase.Random()
		}
		ps.SetTitleAndData("Create [MixedCase]", ret)
		ps.Data = ret
		return Render(rc, as, &vmixedcase.Edit{Model: ret, IsNew: true}, ps, "mixedcase", "Create")
	})
}

func MixedCaseRandom(rc *fasthttp.RequestCtx) {
	Act("mixedcase.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.MixedCase.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random MixedCase")
		}
		return ret.WebPath(), nil
	})
}

func MixedCaseCreate(rc *fasthttp.RequestCtx) {
	Act("mixedcase.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		err = as.Services.MixedCase.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created MixedCase")
		}
		msg := fmt.Sprintf("MixedCase [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func MixedCaseEditForm(rc *fasthttp.RequestCtx) {
	Act("mixedcase.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vmixedcase.Edit{Model: ret}, ps, "mixedcase", ret.String())
	})
}

func MixedCaseEdit(rc *fasthttp.RequestCtx) {
	Act("mixedcase.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := mixedcaseFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		frm.ID = ret.ID
		err = as.Services.MixedCase.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update MixedCase [%s]", frm.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func MixedCaseDelete(rc *fasthttp.RequestCtx) {
	Act("mixedcase.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.MixedCase.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete mixed case [%s]", ret.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/mixedcase", rc, ps)
	})
}

func mixedcaseFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*mixedcase.MixedCase, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	return as.Services.MixedCase.Get(ps.Context, nil, idArg, ps.Logger)
}

func mixedcaseFromForm(rc *fasthttp.RequestCtx, setPK bool) (*mixedcase.MixedCase, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return mixedcase.FromMap(frm, setPK)
}
