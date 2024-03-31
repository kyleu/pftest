// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vbasic"
)

func BasicList(w http.ResponseWriter, r *http.Request) {
	Act("basic.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("basic", ps.Logger)
		var ret basic.Basics
		var err error
		if q == "" {
			ret, err = as.Services.Basic.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Basic.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), w, ps)
			}
		}
		ps.SetTitleAndData("Basics", ret)
		page := &vbasic.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(w, r, as, page, ps, "basic")
	})
}

func BasicDetail(w http.ResponseWriter, r *http.Request) {
	Act("basic.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Basic)", ret)

		relRelationsByBasicIDPrms := ps.Params.Sanitized("relation", ps.Logger)
		relRelationsByBasicID, err := as.Services.Relation.GetByBasicID(ps.Context, nil, ret.ID, relRelationsByBasicIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child relations")
		}
		return Render(w, r, as, &vbasic.Detail{
			Model:  ret,
			Params: ps.Params,

			RelRelationsByBasicID: relRelationsByBasicID,
		}, ps, "basic", ret.TitleString()+"**star")
	})
}

func BasicCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("basic.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &basic.Basic{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = basic.Random()
		}
		ps.SetTitleAndData("Create [Basic]", ret)
		ps.Data = ret
		return Render(w, r, as, &vbasic.Edit{Model: ret, IsNew: true}, ps, "basic", "Create")
	})
}

func BasicRandom(w http.ResponseWriter, r *http.Request) {
	Act("basic.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Basic.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Basic")
		}
		return ret.WebPath(), nil
	})
}

func BasicCreate(w http.ResponseWriter, r *http.Request) {
	Act("basic.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		err = as.Services.Basic.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Basic")
		}
		msg := fmt.Sprintf("Basic [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func BasicEditForm(w http.ResponseWriter, r *http.Request) {
	Act("basic.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(w, r, as, &vbasic.Edit{Model: ret}, ps, "basic", ret.String())
	})
}

func BasicEdit(w http.ResponseWriter, r *http.Request) {
	Act("basic.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := basicFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Basic from form")
		}
		frm.ID = ret.ID
		err = as.Services.Basic.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Basic [%s]", frm.String())
		}
		msg := fmt.Sprintf("Basic [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func BasicDelete(w http.ResponseWriter, r *http.Request) {
	Act("basic.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := basicFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Basic.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete basic [%s]", ret.String())
		}
		msg := fmt.Sprintf("Basic [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/basic", w, ps)
	})
}

func basicFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*basic.Basic, error) {
	idArgStr, err := cutil.RCRequiredString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Basic.Get(ps.Context, nil, idArg, ps.Logger)
}

func basicFromForm(r *http.Request, b []byte, setPK bool) (*basic.Basic, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return basic.FromMap(frm, setPK)
}
