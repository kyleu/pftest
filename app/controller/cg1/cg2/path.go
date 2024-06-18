package cg2

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vg1/vg2/vpath"
)

func PathList(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("path", ps.Logger)
		var ret path.Paths
		var err error
		if q == "" {
			ret, err = as.Services.Path.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Path.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return controller.FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Paths", ret)
		page := &vpath.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return controller.Render(r, as, page, ps, "g1", "g2", "path")
	})
}

func PathDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Path)", ret)

		return controller.Render(r, as, &vpath.Detail{Model: ret}, ps, "g1", "g2", "path", ret.TitleString()+"**star")
	})
}

func PathCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &path.Path{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = path.Random()
		}
		ps.SetTitleAndData("Create [Path]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vpath.Edit{Model: ret, IsNew: true}, ps, "g1", "g2", "path", "Create")
	})
}

func PathRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Path.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Path")
		}
		return ret.WebPath(), nil
	})
}

func PathCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Path from form")
		}
		err = as.Services.Path.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Path")
		}
		msg := fmt.Sprintf("Path [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func PathEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vpath.Edit{Model: ret}, ps, "g1", "g2", "path", ret.String())
	})
}

func PathEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := pathFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Path from form")
		}
		frm.ID = ret.ID
		err = as.Services.Path.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Path [%s]", frm.String())
		}
		msg := fmt.Sprintf("Path [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func PathDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("path.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Path.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete path [%s]", ret.String())
		}
		msg := fmt.Sprintf("Path [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/g1/g2/path", ps)
	})
}

func pathFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*path.Path, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Path.Get(ps.Context, nil, idArg, ps.Logger)
}

func pathFromForm(r *http.Request, b []byte, setPK bool) (*path.Path, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := path.FromMap(frm, setPK)
	return ret, err
}
