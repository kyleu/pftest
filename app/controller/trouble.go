// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vtrouble"
)

func TroubleList(w http.ResponseWriter, r *http.Request) {
	Act("trouble.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("trouble", ps.Logger)
		ret, err := as.Services.Trouble.List(ps.Context, nil, prms, cutil.QueryStringBool(r, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Troubles", ret)
		page := &vtrouble.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "trouble")
	})
}

func TroubleDetail(w http.ResponseWriter, r *http.Request) {
	Act("trouble.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Trouble)", ret)

		return Render(r, as, &vtrouble.Detail{Model: ret}, ps, "trouble", ret.TitleString()+"**star")
	})
}

func TroubleCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("trouble.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &trouble.Trouble{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = trouble.Random()
		}
		ps.SetTitleAndData("Create [Trouble]", ret)
		ps.Data = ret
		return Render(r, as, &vtrouble.Edit{Model: ret, IsNew: true}, ps, "trouble", "Create")
	})
}

func TroubleRandom(w http.ResponseWriter, r *http.Request) {
	Act("trouble.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Trouble.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Trouble")
		}
		return ret.WebPath(), nil
	})
}

func TroubleCreate(w http.ResponseWriter, r *http.Request) {
	Act("trouble.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Trouble from form")
		}
		err = as.Services.Trouble.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Trouble")
		}
		msg := fmt.Sprintf("Trouble [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TroubleEditForm(w http.ResponseWriter, r *http.Request) {
	Act("trouble.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vtrouble.Edit{Model: ret}, ps, "trouble", ret.String())
	})
}

func TroubleEdit(w http.ResponseWriter, r *http.Request) {
	Act("trouble.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := troubleFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Trouble from form")
		}
		frm.From = ret.From
		frm.Where = ret.Where
		err = as.Services.Trouble.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Trouble [%s]", frm.String())
		}
		msg := fmt.Sprintf("Trouble [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TroubleDelete(w http.ResponseWriter, r *http.Request) {
	Act("trouble.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := troubleFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Trouble.Delete(ps.Context, nil, ret.From, ret.Where, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete trouble [%s]", ret.String())
		}
		msg := fmt.Sprintf("Trouble [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/troub/le", ps)
	})
}

func troubleFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*trouble.Trouble, error) {
	fromArg, err := cutil.PathString(r, "from", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [from] as a string argument")
	}
	whereArg, err := cutil.PathArray(r, "where")
	if err != nil {
		return nil, errors.Wrap(err, "must provide [where] as an comma-separated argument")
	}
	includeDeleted := cutil.QueryStringBool(r, "includeDeleted")
	return as.Services.Trouble.Get(ps.Context, nil, fromArg, whereArg, includeDeleted, ps.Logger)
}

func troubleFromForm(r *http.Request, b []byte, setPK bool) (*trouble.Trouble, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return trouble.FromMap(frm, setPK)
}
