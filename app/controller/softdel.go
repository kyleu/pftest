// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vsoftdel"
)

func SoftdelList(rc *fasthttp.RequestCtx) {
	Act("softdel.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("softdel", nil, ps.Logger).Sanitize("softdel")
		ret, err := as.Services.Softdel.List(ps.Context, nil, prms, cutil.QueryStringBool(rc, "includeDeleted"), ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Softdels", ret)
		page := &vsoftdel.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "softdel")
	})
}

func SoftdelDetail(rc *fasthttp.RequestCtx) {
	Act("softdel.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Softdel)", ret)

		return Render(rc, as, &vsoftdel.Detail{Model: ret}, ps, "softdel", ret.TitleString()+"**star")
	})
}

func SoftdelCreateForm(rc *fasthttp.RequestCtx) {
	Act("softdel.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &softdel.Softdel{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = softdel.Random()
		}
		ps.SetTitleAndData("Create [Softdel]", ret)
		ps.Data = ret
		return Render(rc, as, &vsoftdel.Edit{Model: ret, IsNew: true}, ps, "softdel", "Create")
	})
}

func SoftdelRandom(rc *fasthttp.RequestCtx) {
	Act("softdel.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Softdel.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Softdel")
		}
		return ret.WebPath(), nil
	})
}

func SoftdelCreate(rc *fasthttp.RequestCtx) {
	Act("softdel.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		err = as.Services.Softdel.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Softdel")
		}
		msg := fmt.Sprintf("Softdel [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SoftdelEditForm(rc *fasthttp.RequestCtx) {
	Act("softdel.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vsoftdel.Edit{Model: ret}, ps, "softdel", ret.String())
	})
}

func SoftdelEdit(rc *fasthttp.RequestCtx) {
	Act("softdel.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := softdelFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Softdel from form")
		}
		frm.ID = ret.ID
		err = as.Services.Softdel.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Softdel [%s]", frm.String())
		}
		msg := fmt.Sprintf("Softdel [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SoftdelDelete(rc *fasthttp.RequestCtx) {
	Act("softdel.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := softdelFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Softdel.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete softdel [%s]", ret.String())
		}
		msg := fmt.Sprintf("Softdel [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/softdel", rc, ps)
	})
}

func softdelFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*softdel.Softdel, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as a string argument")
	}
	includeDeleted := rc.UserValue("includeDeleted") != nil || cutil.QueryStringBool(rc, "includeDeleted")
	return as.Services.Softdel.Get(ps.Context, nil, idArg, includeDeleted, ps.Logger)
}

func softdelFromForm(rc *fasthttp.RequestCtx, setPK bool) (*softdel.Softdel, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return softdel.FromMap(frm, setPK)
}
