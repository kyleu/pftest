// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views/vcapital"
)

const capitalDefaultTitle = "Capitals"

func CapitalList(rc *fasthttp.RequestCtx) {
	act("capital.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = capitalDefaultTitle
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Capital.List(ps.Context, nil, params.Get("capital", nil, ps.Logger))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vcapital.List{Models: ret, Params: params}, ps, "capital")
	})
}

func CapitalDetail(rc *fasthttp.RequestCtx) {
	act("capital.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		versions, err := as.Services.Capital.GetAllVersions(ps.Context, nil, ret.ID, params.Get("capital", nil, ps.Logger), false)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vcapital.Detail{Model: ret, Versions: versions}, ps, "capital", ret.String())
	})
}

func CapitalVersion(rc *fasthttp.RequestCtx) {
	act("capital.Version", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		latest, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		version, err := RCRequiredInt(rc, "version")
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Capital.GetVersion(ps.Context, nil, latest.ID, version)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vcapital.Detail{Model: ret}, ps, "capital", ret.String())
	})
}

func CapitalCreateForm(rc *fasthttp.RequestCtx) {
	act("capital.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &capital.Capital{}
		ps.Title = "Create [Capital]"
		ps.Data = ret
		return render(rc, as, &vcapital.Edit{Model: ret, IsNew: true}, ps, "capital", "Create")
	})
}

func CapitalCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("capital.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := capital.Random()
		ps.Title = "Create Random [Capital]"
		ps.Data = ret
		return render(rc, as, &vcapital.Edit{Model: ret, IsNew: true}, ps, "capital", "Create")
	})
}

func CapitalCreate(rc *fasthttp.RequestCtx) {
	act("capital.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		err = as.Services.Capital.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Capital")
		}
		msg := fmt.Sprintf("Capital [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func CapitalEditForm(rc *fasthttp.RequestCtx) {
	act("capital.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vcapital.Edit{Model: ret}, ps, "capital", ret.String())
	})
}

func CapitalEdit(rc *fasthttp.RequestCtx) {
	act("capital.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := capitalFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		frm.ID = ret.ID
		err = as.Services.Capital.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Capital [%s]", frm.String())
		}
		msg := fmt.Sprintf("Capital [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func CapitalDelete(rc *fasthttp.RequestCtx) {
	act("capital.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Capital.Delete(ps.Context, nil, ret.ID)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete capital [%s]", ret.String())
		}
		msg := fmt.Sprintf("Capital [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/capital", rc, ps)
	})
}

func capitalFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*capital.Capital, error) {
	idArg, err := RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.Capital.Get(ps.Context, nil, idArg)
}

func capitalFromForm(rc *fasthttp.RequestCtx, setPK bool) (*capital.Capital, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return capital.FromMap(frm, setPK)
}
