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

func HistoryList(rc *fasthttp.RequestCtx) {
	Act("history.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("history", nil, ps.Logger).Sanitize("history")
		ret, err := as.Services.History.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Histories"
		ps.Data = ret
		return Render(rc, as, &vhistory.List{Models: ret, Params: params}, ps, "history")
	})
}

func HistoryDetail(rc *fasthttp.RequestCtx) {
	Act("history.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		hist, err := as.Services.History.GetHistories(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (History)"
		ps.Data = ret
		return Render(rc, as, &vhistory.Detail{
			Model:     ret,
			Params:    params,
			Histories: hist,
		}, ps, "history", ret.String())
	})
}

func HistoryCreateForm(rc *fasthttp.RequestCtx) {
	Act("history.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &history.History{}
		ps.Title = "Create [History]"
		ps.Data = ret
		return Render(rc, as, &vhistory.Edit{Model: ret, IsNew: true}, ps, "history", "Create")
	})
}

func HistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("history.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := history.Random()
		ps.Title = "Create Random History"
		ps.Data = ret
		return Render(rc, as, &vhistory.Edit{Model: ret, IsNew: true}, ps, "history", "Create")
	})
}

func HistoryCreate(rc *fasthttp.RequestCtx) {
	Act("history.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse History from form")
		}
		err = as.Services.History.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created History")
		}
		msg := fmt.Sprintf("History [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func HistoryEditForm(rc *fasthttp.RequestCtx) {
	Act("history.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vhistory.Edit{Model: ret}, ps, "history", ret.String())
	})
}

func HistoryEdit(rc *fasthttp.RequestCtx) {
	Act("history.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := historyFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse History from form")
		}
		frm.ID = ret.ID
		err = as.Services.History.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update History [%s]", frm.String())
		}
		msg := fmt.Sprintf("History [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func HistoryDelete(rc *fasthttp.RequestCtx) {
	Act("history.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.History.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("History [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/history", rc, ps)
	})
}

func HistoryHistory(rc *fasthttp.RequestCtx) {
	Act("history.history", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := historyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		histID, err := cutil.RCRequiredUUID(rc, "historyID")
		if err != nil {
			return "", errors.Wrap(err, "must provide [historyID] as an argument")
		}
		hist, err := as.Services.History.GetHistory(ps.Context, nil, *histID, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = hist.ID.String()
		ps.Data = hist
		return Render(rc, as, &vhistory.History{Model: ret, History: hist}, ps, "history", ret.String(), hist.ID.String())
	})
}

func historyFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*history.History, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.History.Get(ps.Context, nil, idArg, ps.Logger)
}

func historyFromForm(rc *fasthttp.RequestCtx, setPK bool) (*history.History, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return history.FromMap(frm, setPK)
}
