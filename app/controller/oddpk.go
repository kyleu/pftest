package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/oddpk"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/voddpk"
)

func OddPKList(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("oddpk", ps.Logger)
		var ret oddpk.OddPKs
		var err error
		if q == "" {
			ret, err = as.Services.OddPK.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.OddPK.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Odd PKs", ret)
		page := &voddpk.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "oddpk")
	})
}

func OddPKDetail(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddpkFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Odd PK)", ret)

		return Render(r, as, &voddpk.Detail{Model: ret}, ps, "oddpk", ret.TitleString()+"**star")
	})
}

func OddPKCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &oddpk.OddPK{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = oddpk.RandomOddPK()
		}
		ps.SetTitleAndData("Create [OddPK]", ret)
		ps.Data = ret
		return Render(r, as, &voddpk.Edit{Model: ret, IsNew: true}, ps, "oddpk", "Create")
	})
}

func OddPKRandom(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.OddPK.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random OddPK")
		}
		return ret.WebPath(), nil
	})
}

func OddPKCreate(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddpkFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse OddPK from form")
		}
		err = as.Services.OddPK.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created OddPK")
		}
		msg := fmt.Sprintf("OddPK [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func OddPKEditForm(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddpkFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &voddpk.Edit{Model: ret}, ps, "oddpk", ret.String())
	})
}

func OddPKEdit(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddpkFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := oddpkFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse OddPK from form")
		}
		frm.Project = ret.Project
		frm.Path = ret.Path
		err = as.Services.OddPK.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update OddPK [%s]", frm.String())
		}
		msg := fmt.Sprintf("OddPK [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func OddPKDelete(w http.ResponseWriter, r *http.Request) {
	Act("oddpk.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddpkFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.OddPK.Delete(ps.Context, nil, ret.Project, ret.Path, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete odd pk [%s]", ret.String())
		}
		msg := fmt.Sprintf("OddPK [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/oddpk", ps)
	})
}

func oddpkFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*oddpk.OddPK, error) {
	projectArgStr, err := cutil.PathString(r, "project", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [project] as an argument")
	}
	projectArgP := util.UUIDFromString(projectArgStr)
	if projectArgP == nil {
		return nil, errors.Errorf("argument [project] (%s) is not a valid UUID", projectArgStr)
	}
	projectArg := *projectArgP
	pathArg, err := cutil.PathString(r, "path", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [path] as a string argument")
	}
	return as.Services.OddPK.Get(ps.Context, nil, projectArg, pathArg, ps.Logger)
}

func oddpkFromForm(r *http.Request, b []byte, setPK bool) (*oddpk.OddPK, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := oddpk.OddPKFromMap(frm, setPK)
	return ret, err
}
