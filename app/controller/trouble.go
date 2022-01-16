package controller

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/views/vtrouble"
)

func TroubleList(rc *fasthttp.RequestCtx) {
	act("trouble.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Troubles"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Trouble.List(ps.Context, nil, params.Get("trouble", nil, ps.Logger), cutil.RequestCtxBool(rc, "includeDeleted"))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vtrouble.List{Models: ret, Params: params}, ps, "trouble")
	})
}

func TroubleDetail(rc *fasthttp.RequestCtx) {
	act("trouble.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		params := cutil.ParamSetFromRequest(rc)
		selectcols, err := as.Services.Trouble.GetAllSelectcols(ps.Context, nil, ret.From, ret.Where, params.Get("trouble", nil, ps.Logger), false)
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vtrouble.Detail{Model: ret, Selectcols: selectcols, Params: params}, ps, "trouble", ret.String())
	})
}

func TroubleSelectcol(rc *fasthttp.RequestCtx) {
	act("trouble.selectcol", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		latest, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		selectcol, err := RCRequiredInt(rc, "selectcol")
		ret, err := as.Services.Trouble.GetSelectcol(ps.Context, nil, latest.From, latest.Where, selectcol)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vtrouble.Detail{Model: ret}, ps, "trouble", ret.String())
	})
}

func TroubleCreateForm(rc *fasthttp.RequestCtx) {
	act("trouble.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &trouble.Trouble{}
		ps.Title = "Create [Trouble]"
		ps.Data = ret
		return render(rc, as, &vtrouble.Edit{Model: ret, IsNew: true}, ps, "trouble", "Create")
	})
}

func TroubleCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("trouble.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := trouble.Random()
		ps.Title = "Create Random [Trouble]"
		ps.Data = ret
		return render(rc, as, &vtrouble.Edit{Model: ret, IsNew: true}, ps, "trouble", "Create")
	})
}

func TroubleCreate(rc *fasthttp.RequestCtx) {
	act("trouble.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Trouble from form")
		}
		err = as.Services.Trouble.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Trouble")
		}
		msg := fmt.Sprintf("Trouble [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TroubleEditForm(rc *fasthttp.RequestCtx) {
	act("trouble.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vtrouble.Edit{Model: ret}, ps, "trouble", ret.String())
	})
}

func TroubleEdit(rc *fasthttp.RequestCtx) {
	act("trouble.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := troubleFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Trouble from form")
		}
		frm.From = ret.From
		frm.Where = ret.Where
		err = as.Services.Trouble.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Trouble [%s]", frm.String())
		}
		msg := fmt.Sprintf("Trouble [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TroubleDelete(rc *fasthttp.RequestCtx) {
	act("trouble.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Trouble.Delete(ps.Context, nil, ret.From, ret.Where)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete trouble [%s]", ret.String())
		}
		msg := fmt.Sprintf("Trouble [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/trouble", rc, ps)
	})
}

func troubleFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*trouble.Trouble, error) {
	fromArg, err := RCRequiredString(rc, "from", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [from] as an argument")
	}
	whereArgStr, err := RCRequiredString(rc, "where", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [where] as an argument")
	}
	whereArg, err := strconv.Atoi(whereArgStr)
	if err != nil {
		return nil, errors.Wrap(err, "field [where] must be a valid a valid integer")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.RequestCtxBool(rc, "includeDeleted")
	return as.Services.Trouble.Get(ps.Context, nil, fromArg, whereArg, includeDeleted)
}

func troubleFromForm(rc *fasthttp.RequestCtx, setPK bool) (*trouble.Trouble, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return trouble.FromMap(frm, setPK)
}
