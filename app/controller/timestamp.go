// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/views/vtimestamp"
)

const timestampDefaultTitle = "Timestamps"

func TimestampList(rc *fasthttp.RequestCtx) {
	Act("timestamp.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = timestampDefaultTitle
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("timestamp", nil, ps.Logger).Sanitize("timestamp")
		ret, err := as.Services.Timestamp.List(ps.Context, nil, prms, cutil.QueryStringBool(rc, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Timestamps"
		ps.Data = ret
		return Render(rc, as, &vtimestamp.List{Models: ret, Params: params}, ps, "timestamp")
	})
}

func TimestampDetail(rc *fasthttp.RequestCtx) {
	Act("timestamp.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Timestamp)"
		ps.Data = ret
		return Render(rc, as, &vtimestamp.Detail{Model: ret}, ps, "timestamp", ret.String())
	})
}

func TimestampCreateForm(rc *fasthttp.RequestCtx) {
	Act("timestamp.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &timestamp.Timestamp{}
		ps.Title = "Create [Timestamp]"
		ps.Data = ret
		return Render(rc, as, &vtimestamp.Edit{Model: ret, IsNew: true}, ps, "timestamp", "Create")
	})
}

func TimestampCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("timestamp.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := timestamp.Random()
		ps.Title = "Create Random Timestamp"
		ps.Data = ret
		return Render(rc, as, &vtimestamp.Edit{Model: ret, IsNew: true}, ps, "timestamp", "Create")
	})
}

func TimestampCreate(rc *fasthttp.RequestCtx) {
	Act("timestamp.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		err = as.Services.Timestamp.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Timestamp")
		}
		msg := fmt.Sprintf("Timestamp [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TimestampEditForm(rc *fasthttp.RequestCtx) {
	Act("timestamp.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vtimestamp.Edit{Model: ret}, ps, "timestamp", ret.String())
	})
}

func TimestampEdit(rc *fasthttp.RequestCtx) {
	Act("timestamp.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := timestampFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		frm.ID = ret.ID
		err = as.Services.Timestamp.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Timestamp [%s]", frm.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TimestampDelete(rc *fasthttp.RequestCtx) {
	Act("timestamp.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Timestamp.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete timestamp [%s]", ret.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/timestamp", rc, ps)
	})
}

func timestampFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*timestamp.Timestamp, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.QueryStringBool(rc, "includeDeleted")
	return as.Services.Timestamp.Get(ps.Context, nil, idArg, includeDeleted, ps.Logger)
}

func timestampFromForm(rc *fasthttp.RequestCtx, setPK bool) (*timestamp.Timestamp, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return timestamp.FromMap(frm, setPK)
}
