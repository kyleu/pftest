// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vreference"
)

func ReferenceList(rc *fasthttp.RequestCtx) {
	Act("reference.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("reference", nil, ps.Logger).Sanitize("reference")
		ret, err := as.Services.Reference.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "References"
		ps.Data = ret
		return Render(rc, as, &vreference.List{Models: ret, Params: params}, ps, "reference")
	})
}

func ReferenceDetail(rc *fasthttp.RequestCtx) {
	Act("reference.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Reference)"
		ps.Data = ret
		return Render(rc, as, &vreference.Detail{Model: ret}, ps, "reference", ret.String())
	})
}

func ReferenceCreateForm(rc *fasthttp.RequestCtx) {
	Act("reference.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &reference.Reference{}
		ps.Title = "Create [Reference]"
		ps.Data = ret
		return Render(rc, as, &vreference.Edit{Model: ret, IsNew: true}, ps, "reference", "Create")
	})
}

func ReferenceCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("reference.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := reference.Random()
		ps.Title = "Create Random Reference"
		ps.Data = ret
		return Render(rc, as, &vreference.Edit{Model: ret, IsNew: true}, ps, "reference", "Create")
	})
}

func ReferenceCreate(rc *fasthttp.RequestCtx) {
	Act("reference.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Reference from form")
		}
		err = as.Services.Reference.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Reference")
		}
		msg := fmt.Sprintf("Reference [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func ReferenceEditForm(rc *fasthttp.RequestCtx) {
	Act("reference.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vreference.Edit{Model: ret}, ps, "reference", ret.String())
	})
}

func ReferenceEdit(rc *fasthttp.RequestCtx) {
	Act("reference.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := referenceFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Reference from form")
		}
		frm.ID = ret.ID
		err = as.Services.Reference.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Reference [%s]", frm.String())
		}
		msg := fmt.Sprintf("Reference [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func ReferenceDelete(rc *fasthttp.RequestCtx) {
	Act("reference.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := referenceFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Reference.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete reference [%s]", ret.String())
		}
		msg := fmt.Sprintf("Reference [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/reference", rc, ps)
	})
}

func referenceFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*reference.Reference, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func referenceFromForm(rc *fasthttp.RequestCtx, setPK bool) (*reference.Reference, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return reference.FromMap(frm, setPK)
}
