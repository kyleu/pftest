// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vmixedcase"
)

func MixedCaseList(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("mixedcase", ps.Logger)
		ret, err := as.Services.MixedCase.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Mixed Cases", ret)
		page := &vmixedcase.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "mixedcase")
	})
}

func MixedCaseDetail(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Mixed Case)", ret)

		return Render(r, as, &vmixedcase.Detail{Model: ret}, ps, "mixedcase", ret.TitleString()+"**star")
	})
}

func MixedCaseCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &mixedcase.MixedCase{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = mixedcase.Random()
		}
		ps.SetTitleAndData("Create [MixedCase]", ret)
		ps.Data = ret
		return Render(r, as, &vmixedcase.Edit{Model: ret, IsNew: true}, ps, "mixedcase", "Create")
	})
}

func MixedCaseRandom(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.MixedCase.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random MixedCase")
		}
		return ret.WebPath(), nil
	})
}

func MixedCaseCreate(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		err = as.Services.MixedCase.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created MixedCase")
		}
		msg := fmt.Sprintf("MixedCase [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func MixedCaseEditForm(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vmixedcase.Edit{Model: ret}, ps, "mixedcase", ret.String())
	})
}

func MixedCaseEdit(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := mixedcaseFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse MixedCase from form")
		}
		frm.ID = ret.ID
		err = as.Services.MixedCase.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update MixedCase [%s]", frm.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func MixedCaseDelete(w http.ResponseWriter, r *http.Request) {
	Act("mixedcase.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := mixedcaseFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.MixedCase.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete mixed case [%s]", ret.String())
		}
		msg := fmt.Sprintf("MixedCase [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/mixedcase", ps)
	})
}

func mixedcaseFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*mixedcase.MixedCase, error) {
	idArg, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	return as.Services.MixedCase.Get(ps.Context, nil, idArg, ps.Logger)
}

func mixedcaseFromForm(r *http.Request, b []byte, setPK bool) (*mixedcase.MixedCase, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return mixedcase.FromMap(frm, setPK)
}
