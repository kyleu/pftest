// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/hist"
	"github.com/kyleu/pftest/views/vhist"
)

func HistList(rc *fasthttp.RequestCtx) {
	Act("hist.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("hist", nil, ps.Logger).Sanitize("hist")
		ret, err := as.Services.Hist.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Hists"
		ps.Data = ret
		page := &vhist.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "hist")
	})
}

func HistDetail(rc *fasthttp.RequestCtx) {
	Act("hist.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		hist, err := as.Services.Hist.GetHistories(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Hist)"
		ps.Data = ret

		return Render(rc, as, &vhist.Detail{
			Model:     ret,
			Params:    ps.Params,
			Histories: hist,
		}, ps, "hist", ret.String())
	})
}

func HistCreateForm(rc *fasthttp.RequestCtx) {
	Act("hist.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &hist.Hist{}
		ps.Title = "Create [Hist]"
		ps.Data = ret
		return Render(rc, as, &vhist.Edit{Model: ret, IsNew: true}, ps, "hist", "Create")
	})
}

func HistCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("hist.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := hist.Random()
		ps.Title = "Create Random Hist"
		ps.Data = ret
		return Render(rc, as, &vhist.Edit{Model: ret, IsNew: true}, ps, "hist", "Create")
	})
}

func HistCreate(rc *fasthttp.RequestCtx) {
	Act("hist.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Hist from form")
		}
		err = as.Services.Hist.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Hist")
		}
		msg := fmt.Sprintf("Hist [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func HistEditForm(rc *fasthttp.RequestCtx) {
	Act("hist.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vhist.Edit{Model: ret}, ps, "hist", ret.String())
	})
}

func HistEdit(rc *fasthttp.RequestCtx) {
	Act("hist.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := histFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Hist from form")
		}
		frm.ID = ret.ID
		err = as.Services.Hist.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Hist [%s]", frm.String())
		}
		msg := fmt.Sprintf("Hist [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func HistDelete(rc *fasthttp.RequestCtx) {
	Act("hist.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Hist.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete hist [%s]", ret.String())
		}
		msg := fmt.Sprintf("Hist [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/hist", rc, ps)
	})
}

func HistHistory(rc *fasthttp.RequestCtx) {
	Act("hist.history", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := histFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		histID, err := cutil.RCRequiredUUID(rc, "historyID")
		if err != nil {
			return "", errors.Wrap(err, "must provide [historyID] as an argument")
		}
		hist, err := as.Services.Hist.GetHistory(ps.Context, nil, *histID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = hist.ID.String()
		ps.Data = hist
		return Render(rc, as, &vhist.History{Model: ret, History: hist}, ps, "hist", ret.String(), hist.ID.String())
	})
}

func histFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*hist.Hist, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	return as.Services.Hist.Get(ps.Context, nil, idArg, ps.Logger)
}

func histFromForm(rc *fasthttp.RequestCtx, setPK bool) (*hist.Hist, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return hist.FromMap(frm, setPK)
}
