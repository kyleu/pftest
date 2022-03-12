// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/history"
	"github.com/kyleu/pftest/views/vhistory"
)

const historyDefaultTitle = "Histories"

func HistoryList(rc *fasthttp.RequestCtx) {
	act("history.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = historyDefaultTitle
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("history", nil, ps.Logger)
		ret, err := as.Services.History.List(ps.Context, nil, prms)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vhistory.List{Models: ret, Params: params}, ps, "history")
	})
}

func HistoryDetail(rc *fasthttp.RequestCtx) {
	act("history.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		hist, err := as.Services.History.GetHistories(ps.Context, nil, ret.ID)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vhistory.Detail{
			Model:     ret,
			Params:    params,
			Histories: hist,
		}, ps, "history", ret.String())
	})
}

func HistoryCreateForm(rc *fasthttp.RequestCtx) {
	act("history.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &history.History{}
		ps.Title = "Create [History]"
		ps.Data = ret
		return render(rc, as, &vhistory.Edit{Model: ret, IsNew: true}, ps, "history", "Create")
	})
}

func HistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("history.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := history.Random()
		ps.Title = "Create Random [History]"
		ps.Data = ret
		return render(rc, as, &vhistory.Edit{Model: ret, IsNew: true}, ps, "history", "Create")
	})
}

func HistoryCreate(rc *fasthttp.RequestCtx) {
	act("history.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse History from form")
		}
		err = as.Services.History.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created History")
		}
		msg := fmt.Sprintf("History [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func HistoryEditForm(rc *fasthttp.RequestCtx) {
	act("history.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vhistory.Edit{Model: ret}, ps, "history", ret.String())
	})
}

func HistoryEdit(rc *fasthttp.RequestCtx) {
	act("history.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := historyFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse History from form")
		}
		frm.ID = ret.ID
		err = as.Services.History.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update History [%s]", frm.String())
		}
		msg := fmt.Sprintf("History [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func HistoryDelete(rc *fasthttp.RequestCtx) {
	act("history.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.History.Delete(ps.Context, nil, ret.ID)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("History [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/history", rc, ps)
	})
}

func HistoryHistory(rc *fasthttp.RequestCtx) {
	act("history.history", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		histID, err := RCRequiredUUID(rc, "historyID")
		if err != nil {
			return "", errors.Wrap(err, "must provide [historyID] as an argument")
		}
		hist, err := as.Services.History.GetHistory(ps.Context, nil, *histID)
		if err != nil {
			return "", err
		}
		ps.Title = hist.ID.String()
		ps.Data = hist
		return render(rc, as, &vhistory.History{Model: ret, History: hist}, ps, "history", ret.String(), hist.ID.String())
	})
}

func historyFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*history.History, error) {
	idArg, err := RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.History.Get(ps.Context, nil, idArg)
}

func historyFromForm(rc *fasthttp.RequestCtx, setPK bool) (*history.History, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return history.FromMap(frm, setPK)
}
