package coddpk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/oddpk/oddrel"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/voddpk/voddrel"
)

func OddrelList(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(cutil.QueryStringString(r, "q"))
		prms := ps.Params.Sanitized("oddrel", ps.Logger)
		var ret oddrel.Oddrels
		var err error
		if q == "" {
			ret, err = as.Services.Oddrel.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Oddrel.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return controller.FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Oddrels", ret)
		page := &voddrel.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return controller.Render(r, as, page, ps, "oddpk", "oddrel")
	})
}

func OddrelDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddrelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Oddrel)", ret)

		return controller.Render(r, as, &voddrel.Detail{Model: ret}, ps, "oddpk", "oddrel", ret.TitleString()+"**star")
	})
}

func OddrelCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &oddrel.Oddrel{}
		if cutil.QueryStringString(r, "prototype") == util.KeyRandom {
			ret = oddrel.RandomOddrel()
		}
		ps.SetTitleAndData("Create [Oddrel]", ret)
		return controller.Render(r, as, &voddrel.Edit{Model: ret, IsNew: true}, ps, "oddpk", "oddrel", "Create")
	})
}

func OddrelRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Oddrel.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Oddrel")
		}
		return ret.WebPath(), nil
	})
}

func OddrelCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddrelFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Oddrel from form")
		}
		err = as.Services.Oddrel.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Oddrel")
		}
		msg := fmt.Sprintf("Oddrel [%s] created", ret.TitleString())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func OddrelEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddrelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &voddrel.Edit{Model: ret}, ps, "oddpk", "oddrel", ret.String())
	})
}

func OddrelEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddrelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := oddrelFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Oddrel from form")
		}
		frm.ID = ret.ID
		err = as.Services.Oddrel.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Oddrel [%s]", frm.String())
		}
		msg := fmt.Sprintf("Oddrel [%s] updated", frm.TitleString())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func OddrelDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("oddrel.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := oddrelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Oddrel.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete oddrel [%s]", ret.String())
		}
		msg := fmt.Sprintf("Oddrel [%s] deleted", ret.TitleString())
		return controller.FlashAndRedir(true, msg, "/oddpk/oddrel", ps)
	})
}

func oddrelFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*oddrel.Oddrel, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Oddrel.Get(ps.Context, nil, idArg, ps.Logger)
}

func oddrelFromForm(r *http.Request, b []byte, setPK bool) (*oddrel.Oddrel, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := oddrel.OddrelFromMap(frm, setPK)
	return ret, err
}
