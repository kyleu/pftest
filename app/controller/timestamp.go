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

func TimestampList(rc *fasthttp.RequestCtx) {
	act("timestamp.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Timestamps"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Timestamp.List(ps.Context, nil, params.Get("timestamp", nil, ps.Logger), cutil.RequestCtxBool(rc, "includeDeleted"))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vtimestamp.List{Models: ret, Params: params}, ps, "timestamp")
	})
}

func TimestampDetail(rc *fasthttp.RequestCtx) {
	act("timestamp.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vtimestamp.Detail{Model: ret}, ps, "timestamp", ret.String())
	})
}

func TimestampCreateForm(rc *fasthttp.RequestCtx) {
	act("timestamp.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &timestamp.Timestamp{}
		ps.Title = "Create [Timestamp]"
		ps.Data = ret
		return render(rc, as, &vtimestamp.Edit{Model: ret, IsNew: true}, ps, "timestamp", "Create")
	})
}

func TimestampCreate(rc *fasthttp.RequestCtx) {
	act("timestamp.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		err = as.Services.Timestamp.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Timestamp")
		}
		msg := fmt.Sprintf("Timestamp [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TimestampEditForm(rc *fasthttp.RequestCtx) {
	act("timestamp.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vtimestamp.Edit{Model: ret}, ps, "timestamp", ret.String())
	})
}

func TimestampEdit(rc *fasthttp.RequestCtx) {
	act("timestamp.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := timestampFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		frm.ID = ret.ID
		err = as.Services.Timestamp.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Timestamp [%s]", frm.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TimestampDelete(rc *fasthttp.RequestCtx) {
	act("timestamp.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Timestamp.Delete(ps.Context, nil, ret.ID)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete timestamp [%s]", ret.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/timestamp", rc, ps)
	})
}

func timestampFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*timestamp.Timestamp, error) {
	idArg, err := RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.RequestCtxBool(rc, "includeDeleted")
	return as.Services.Timestamp.Get(ps.Context, nil, idArg, includeDeleted)
}

func timestampFromForm(rc *fasthttp.RequestCtx, setPK bool) (*timestamp.Timestamp, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return timestamp.FromMap(frm, setPK)
}
