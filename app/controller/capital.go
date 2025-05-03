package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vcapital"
)

func CapitalList(w http.ResponseWriter, r *http.Request) {
	Act("capital.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("capital", ps.Logger)
		ret, err := as.Services.Capital.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Capitals", ret)
		page := &vcapital.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "capital")
	})
}

func CapitalDetail(w http.ResponseWriter, r *http.Request) {
	Act("capital.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Capital)", ret)

		return Render(r, as, &vcapital.Detail{Model: ret}, ps, "capital", ret.TitleString()+"**star")
	})
}

func CapitalCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("capital.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &capital.Capital{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = capital.RandomCapital()
		}
		ps.SetTitleAndData("Create [Capital]", ret)
		return Render(r, as, &vcapital.Edit{Model: ret, IsNew: true}, ps, "capital", "Create")
	})
}

func CapitalRandom(w http.ResponseWriter, r *http.Request) {
	Act("capital.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Capital.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Capital")
		}
		return ret.WebPath(), nil
	})
}

func CapitalCreate(w http.ResponseWriter, r *http.Request) {
	Act("capital.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		err = as.Services.Capital.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Capital")
		}
		msg := fmt.Sprintf("Capital [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func CapitalEditForm(w http.ResponseWriter, r *http.Request) {
	Act("capital.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vcapital.Edit{Model: ret}, ps, "capital", ret.String())
	})
}

func CapitalEdit(w http.ResponseWriter, r *http.Request) {
	Act("capital.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := capitalFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Capital from form")
		}
		frm.ID = ret.ID
		err = as.Services.Capital.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Capital [%s]", frm.String())
		}
		msg := fmt.Sprintf("Capital [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func CapitalDelete(w http.ResponseWriter, r *http.Request) {
	Act("capital.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := capitalFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Capital.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete capital [%s]", ret.String())
		}
		msg := fmt.Sprintf("Capital [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/capital", ps)
	})
}

func capitalFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*capital.Capital, error) {
	idArg, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	return as.Services.Capital.Get(ps.Context, nil, idArg, ps.Logger)
}

func capitalFromForm(r *http.Request, b []byte, setPK bool) (*capital.Capital, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := capital.CapitalFromMap(frm, setPK)
	return ret, err
}
