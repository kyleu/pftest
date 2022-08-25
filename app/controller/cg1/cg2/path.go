// Content managed by Project Forge, see [projectforge.md] for details.
package cg2

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vg1/vg2/vpath"
)

func PathList(rc *fasthttp.RequestCtx) {
	controller.Act("path.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("path", nil, ps.Logger).Sanitize("path")
		ret, err := as.Services.Path.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Paths"
		ps.Data = ret
		return controller.Render(rc, as, &vpath.List{Models: ret, Params: params}, ps, "g1", "g2", "path")
	})
}

func PathDetail(rc *fasthttp.RequestCtx) {
	controller.Act("path.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Path)"
		ps.Data = ret
		return controller.Render(rc, as, &vpath.Detail{Model: ret}, ps, "g1", "g2", "path", ret.String())
	})
}

func PathCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("path.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &path.Path{}
		ps.Title = "Create [Path]"
		ps.Data = ret
		return controller.Render(rc, as, &vpath.Edit{Model: ret, IsNew: true}, ps, "g1", "g2", "path", "Create")
	})
}

func PathCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("path.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := path.Random()
		ps.Title = "Create Random Path"
		ps.Data = ret
		return controller.Render(rc, as, &vpath.Edit{Model: ret, IsNew: true}, ps, "g1", "g2", "path", "Create")
	})
}

func PathCreate(rc *fasthttp.RequestCtx) {
	controller.Act("path.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Path from form")
		}
		err = as.Services.Path.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Path")
		}
		msg := fmt.Sprintf("Path [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func PathEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("path.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vpath.Edit{Model: ret}, ps, "g1", "g2", "path", ret.String())
	})
}

func PathEdit(rc *fasthttp.RequestCtx) {
	controller.Act("path.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := pathFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Path from form")
		}
		frm.ID = ret.ID
		err = as.Services.Path.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Path [%s]", frm.String())
		}
		msg := fmt.Sprintf("Path [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func PathDelete(rc *fasthttp.RequestCtx) {
	controller.Act("path.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := pathFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Path.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete path [%s]", ret.String())
		}
		msg := fmt.Sprintf("Path [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/path", rc, ps)
	})
}

func pathFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*path.Path, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func pathFromForm(rc *fasthttp.RequestCtx, setPK bool) (*path.Path, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return path.FromMap(frm, setPK)
}
