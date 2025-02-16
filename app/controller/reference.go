package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vreference"
)

func ReferenceList(w http.ResponseWriter, r *http.Request) {
	Act("reference.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("reference", ps.Logger)
		var ret reference.References
		var err error
		if q == "" {
			ret, err = as.Services.Reference.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Reference.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("References", ret)
		page := &vreference.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "reference")
	})
}

func ReferenceDetail(w http.ResponseWriter, r *http.Request) {
	Act("reference.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Reference)", ret)

		return Render(r, as, &vreference.Detail{Model: ret}, ps, "reference", ret.TitleString()+"**star")
	})
}

func ReferenceCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("reference.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &reference.Reference{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = reference.RandomReference()
		}
		ps.SetTitleAndData("Create [Reference]", ret)
		ps.Data = ret
		return Render(r, as, &vreference.Edit{Model: ret, IsNew: true}, ps, "reference", "Create")
	})
}

func ReferenceRandom(w http.ResponseWriter, r *http.Request) {
	Act("reference.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Reference.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Reference")
		}
		return ret.WebPath(), nil
	})
}

func ReferenceCreate(w http.ResponseWriter, r *http.Request) {
	Act("reference.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Reference from form")
		}
		err = as.Services.Reference.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Reference")
		}
		msg := fmt.Sprintf("Reference [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func ReferenceEditForm(w http.ResponseWriter, r *http.Request) {
	Act("reference.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vreference.Edit{Model: ret}, ps, "reference", ret.String())
	})
}

func ReferenceEdit(w http.ResponseWriter, r *http.Request) {
	Act("reference.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := referenceFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Reference from form")
		}
		frm.ID = ret.ID
		err = as.Services.Reference.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Reference [%s]", frm.String())
		}
		msg := fmt.Sprintf("Reference [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func ReferenceDelete(w http.ResponseWriter, r *http.Request) {
	Act("reference.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Reference.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete reference [%s]", ret.String())
		}
		msg := fmt.Sprintf("Reference [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/reference", ps)
	})
}

func referenceFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*reference.Reference, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Reference.Get(ps.Context, nil, idArg, ps.Logger)
}

func referenceFromForm(r *http.Request, b []byte, setPK bool) (*reference.Reference, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := reference.ReferenceFromMap(frm, setPK)
	return ret, err
}
