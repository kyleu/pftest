package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vsoftdel"
)

func SoftdelList(w http.ResponseWriter, r *http.Request) {
	Act("softdel.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("softdel", ps.Logger)
		ret, err := as.Services.Softdel.List(ps.Context, nil, prms, cutil.QueryStringBool(r, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Softdels", ret)
		page := &vsoftdel.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "softdel")
	})
}

func SoftdelDetail(w http.ResponseWriter, r *http.Request) {
	Act("softdel.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Softdel)", ret)

		return Render(r, as, &vsoftdel.Detail{Model: ret}, ps, "softdel", ret.TitleString()+"**star")
	})
}

func SoftdelCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("softdel.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &softdel.Softdel{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = softdel.Random()
		}
		ps.SetTitleAndData("Create [Softdel]", ret)
		ps.Data = ret
		return Render(r, as, &vsoftdel.Edit{Model: ret, IsNew: true}, ps, "softdel", "Create")
	})
}

func SoftdelRandom(w http.ResponseWriter, r *http.Request) {
	Act("softdel.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Softdel.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Softdel")
		}
		return ret.WebPath(), nil
	})
}

func SoftdelCreate(w http.ResponseWriter, r *http.Request) {
	Act("softdel.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		err = as.Services.Softdel.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Softdel")
		}
		msg := fmt.Sprintf("Softdel [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SoftdelEditForm(w http.ResponseWriter, r *http.Request) {
	Act("softdel.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vsoftdel.Edit{Model: ret}, ps, "softdel", ret.String())
	})
}

func SoftdelEdit(w http.ResponseWriter, r *http.Request) {
	Act("softdel.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := softdelFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		frm.ID = ret.ID
		err = as.Services.Softdel.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Softdel [%s]", frm.String())
		}
		msg := fmt.Sprintf("Softdel [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SoftdelDelete(w http.ResponseWriter, r *http.Request) {
	Act("softdel.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Softdel.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete softdel [%s]", ret.String())
		}
		msg := fmt.Sprintf("Softdel [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/softdel", ps)
	})
}

func softdelFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*softdel.Softdel, error) {
	idArg, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	includeDeleted := cutil.QueryStringBool(r, "includeDeleted")
	return as.Services.Softdel.Get(ps.Context, nil, idArg, includeDeleted, ps.Logger)
}

func softdelFromForm(r *http.Request, b []byte, setPK bool) (*softdel.Softdel, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := softdel.FromMap(frm, setPK)
	return ret, err
}
