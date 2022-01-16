package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/mixed_case"
	"github.com/kyleu/pftest/views/vmixed_case"
)

func MixedCaseList(rc *fasthttp.RequestCtx) {
	act("mixed_case.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "MixedCases"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.MixedCase.List(ps.Context, nil, params.Get("mixed_case", nil, ps.Logger))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vmixed_case.List{Models: ret, Params: params}, ps, "mixed_case")
	})
}

func MixedCaseDetail(rc *fasthttp.RequestCtx) {
	act("mixed_case.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixed_caseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vmixed_case.Detail{Model: ret}, ps, "mixed_case", ret.String())
	})
}

func MixedCaseCreateForm(rc *fasthttp.RequestCtx) {
	act("mixed_case.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &mixed_case.MixedCase{}
		ps.Title = "Create [MixedCase]"
		ps.Data = ret
		return render(rc, as, &vmixed_case.Edit{Model: ret, IsNew: true}, ps, "mixed_case", "Create")
	})
}

func MixedCaseCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("mixed_case.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := mixed_case.Random()
		ps.Title = "Create Random [MixedCase]"
		ps.Data = ret
		return render(rc, as, &vmixed_case.Edit{Model: ret, IsNew: true}, ps, "mixed_case", "Create")
	})
}

func MixedCaseCreate(rc *fasthttp.RequestCtx) {
	act("mixed_case.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixed_caseFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		err = as.Services.MixedCase.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created MixedCase")
		}
		msg := fmt.Sprintf("MixedCase [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func MixedCaseEditForm(rc *fasthttp.RequestCtx) {
	act("mixed_case.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixed_caseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vmixed_case.Edit{Model: ret}, ps, "mixed_case", ret.String())
	})
}

func MixedCaseEdit(rc *fasthttp.RequestCtx) {
	act("mixed_case.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixed_caseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := mixed_caseFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		frm.ID = ret.ID
		err = as.Services.MixedCase.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update MixedCase [%s]", frm.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func MixedCaseDelete(rc *fasthttp.RequestCtx) {
	act("mixed_case.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixed_caseFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.MixedCase.Delete(ps.Context, nil, ret.ID)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete mixed case [%s]", ret.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/mixedCase", rc, ps)
	})
}

func mixed_caseFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*mixed_case.MixedCase, error) {
	idArg, err := RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.MixedCase.Get(ps.Context, nil, idArg)
}

func mixed_caseFromForm(rc *fasthttp.RequestCtx, setPK bool) (*mixed_case.MixedCase, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return mixed_case.FromMap(frm, setPK)
}
