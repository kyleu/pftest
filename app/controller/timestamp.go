package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vtimestamp"
)

func TimestampList(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("timestamp", ps.Logger)
		ret, err := as.Services.Timestamp.List(ps.Context, nil, prms, cutil.QueryStringBool(r, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Timestamps", ret)
		page := &vtimestamp.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "timestamp")
	})
}

func TimestampDetail(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Timestamp)", ret)

		return Render(r, as, &vtimestamp.Detail{Model: ret}, ps, "timestamp", ret.TitleString()+"**star")
	})
}

func TimestampCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &timestamp.Timestamp{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = timestamp.RandomTimestamp()
		}
		ps.SetTitleAndData("Create [Timestamp]", ret)
		return Render(r, as, &vtimestamp.Edit{Model: ret, IsNew: true}, ps, "timestamp", "Create")
	})
}

func TimestampRandom(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Timestamp.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Timestamp")
		}
		return ret.WebPath(), nil
	})
}

func TimestampCreate(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		err = as.Services.Timestamp.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Timestamp")
		}
		msg := fmt.Sprintf("Timestamp [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TimestampEditForm(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vtimestamp.Edit{Model: ret}, ps, "timestamp", ret.String())
	})
}

func TimestampEdit(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := timestampFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Timestamp from form")
		}
		frm.ID = ret.ID
		err = as.Services.Timestamp.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Timestamp [%s]", frm.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TimestampDelete(w http.ResponseWriter, r *http.Request) {
	Act("timestamp.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := timestampFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Timestamp.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete timestamp [%s]", ret.String())
		}
		msg := fmt.Sprintf("Timestamp [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/timestamp", ps)
	})
}

func timestampFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*timestamp.Timestamp, error) {
	idArg, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	includeDeleted := cutil.QueryStringBool(r, "includeDeleted")
	return as.Services.Timestamp.Get(ps.Context, nil, idArg, includeDeleted, ps.Logger)
}

func timestampFromForm(r *http.Request, b []byte, setPK bool) (*timestamp.Timestamp, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := timestamp.TimestampFromMap(frm, setPK)
	return ret, err
}
