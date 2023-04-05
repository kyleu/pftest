// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/views/vtrouble"
)

func TroubleList(rc *fasthttp.RequestCtx) {
	Act("trouble.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("trouble", nil, ps.Logger).Sanitize("trouble")
		ret, err := as.Services.Trouble.List(ps.Context, nil, prms, cutil.QueryStringBool(rc, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Troubles"
		ps.Data = ret
		page := &vtrouble.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "trouble")
	})
}

func TroubleDetail(rc *fasthttp.RequestCtx) {
	Act("trouble.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		prms := ps.Params.Get("trouble", nil, ps.Logger).Sanitize("trouble")
		selectcols, err := as.Services.Trouble.GetAllSelectcols(ps.Context, nil, ret.From, ret.Where, prms, false, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Trouble)"
		ps.Data = ret

		return Render(rc, as, &vtrouble.Detail{
			Model:      ret,
			Params:     ps.Params,
			Selectcols: selectcols,
		}, ps, "trouble", ret.String())
	})
}

func TroubleSelectcol(rc *fasthttp.RequestCtx) {
	Act("trouble.selectcol", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		latest, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		selectcol, err := cutil.RCRequiredInt(rc, "selectcol")
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Trouble.GetSelectcol(ps.Context, nil, latest.From, latest.Where, selectcol, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return Render(rc, as, &vtrouble.Detail{Model: ret}, ps, "trouble", ret.String())
	})
}

func TroubleCreateForm(rc *fasthttp.RequestCtx) {
	Act("trouble.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &trouble.Trouble{}
		ps.Title = "Create [Trouble]"
		ps.Data = ret
		return Render(rc, as, &vtrouble.Edit{Model: ret, IsNew: true}, ps, "trouble", "Create")
	})
}

func TroubleCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("trouble.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := trouble.Random()
		ps.Title = "Create Random Trouble"
		ps.Data = ret
		return Render(rc, as, &vtrouble.Edit{Model: ret, IsNew: true}, ps, "trouble", "Create")
	})
}

func TroubleCreate(rc *fasthttp.RequestCtx) {
	Act("trouble.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Trouble from form")
		}
		err = as.Services.Trouble.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Trouble")
		}
		msg := fmt.Sprintf("Trouble [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TroubleEditForm(rc *fasthttp.RequestCtx) {
	Act("trouble.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vtrouble.Edit{Model: ret}, ps, "trouble", ret.String())
	})
}

func TroubleEdit(rc *fasthttp.RequestCtx) {
	Act("trouble.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
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
		err = as.Services.Trouble.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Trouble [%s]", frm.String())
		}
		msg := fmt.Sprintf("Trouble [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TroubleDelete(rc *fasthttp.RequestCtx) {
	Act("trouble.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Trouble.Delete(ps.Context, nil, ret.From, ret.Where, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete trouble [%s]", ret.String())
		}
		msg := fmt.Sprintf("Trouble [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/trouble", rc, ps)
	})
}

func troubleFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*trouble.Trouble, error) {
	fromArg, err := cutil.RCRequiredString(rc, "from", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [from] as an argument")
	}
	whereArg, err := cutil.RCRequiredArray(rc, "where")
	if err != nil {
		return nil, errors.Wrap(err, "must provide [where] as an comma-separated argument")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.QueryStringBool(rc, "includeDeleted")
	return as.Services.Trouble.Get(ps.Context, nil, fromArg, whereArg, includeDeleted, ps.Logger)
}

func troubleFromForm(rc *fasthttp.RequestCtx, setPK bool) (*trouble.Trouble, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return trouble.FromMap(frm, setPK)
}
