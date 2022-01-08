package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/views/vsoftdel"
)

func SoftdelList(rc *fasthttp.RequestCtx) {
	act("softdel.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Softdels"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Softdel.List(ps.Context, nil, params.Get("softdel", nil, ps.Logger), cutil.RequestCtxBool(rc, "includeDeleted"))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vsoftdel.List{Models: ret, Params: params}, ps, "softdel")
	})
}

func SoftdelDetail(rc *fasthttp.RequestCtx) {
	act("softdel.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vsoftdel.Detail{Model: ret}, ps, "softdel", ret.String())
	})
}

func SoftdelCreateForm(rc *fasthttp.RequestCtx) {
	act("softdel.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &softdel.Softdel{}
		ps.Title = "Create [Softdel]"
		ps.Data = ret
		return render(rc, as, &vsoftdel.Edit{Model: ret, IsNew: true}, ps, "softdel", "Create")
	})
}

func SoftdelCreate(rc *fasthttp.RequestCtx) {
	act("softdel.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		err = as.Services.Softdel.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Softdel")
		}
		msg := fmt.Sprintf("Softdel [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SoftdelEditForm(rc *fasthttp.RequestCtx) {
	act("softdel.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vsoftdel.Edit{Model: ret}, ps, "softdel", ret.String())
	})
}

func SoftdelEdit(rc *fasthttp.RequestCtx) {
	act("softdel.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := softdelFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		frm.ID = ret.ID
		err = as.Services.Softdel.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Softdel [%s]", frm.String())
		}
		msg := fmt.Sprintf("Softdel [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func softdelFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*softdel.Softdel, error) {
	idArg, err := rcRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.RequestCtxBool(rc, "includeDeleted")
	return as.Services.Softdel.Get(ps.Context, nil, idArg, includeDeleted)
}

func softdelFromForm(rc *fasthttp.RequestCtx, setPK bool) (*softdel.Softdel, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return softdel.FromMap(frm, setPK)
}
