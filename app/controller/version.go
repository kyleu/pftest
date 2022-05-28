// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/version"
	"github.com/kyleu/pftest/views/vversion"
)

const versionDefaultTitle = "Versions"

func VersionList(rc *fasthttp.RequestCtx) {
	act("version.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = versionDefaultTitle
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("version", nil, ps.Logger).Sanitize("version")
		ret, err := as.Services.Version.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Versions"
		ps.Data = ret
		return render(rc, as, &vversion.List{Models: ret, Params: params}, ps, "version")
	})
}

func VersionDetail(rc *fasthttp.RequestCtx) {
	act("version.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := versionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		prms := params.Get("version", nil, ps.Logger).Sanitize("version")
		revisions, err := as.Services.Version.GetAllRevisions(ps.Context, nil, ret.ID, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Version)"
		ps.Data = ret
		return render(rc, as, &vversion.Detail{
			Model:     ret,
			Params:    params,
			Revisions: revisions,
		}, ps, "version", ret.String())
	})
}

func VersionRevision(rc *fasthttp.RequestCtx) {
	act("version.revision", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		latest, err := versionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		revision, err := cutil.RCRequiredInt(rc, "revision")
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Version.GetRevision(ps.Context, nil, latest.ID, revision, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vversion.Detail{Model: ret}, ps, "version", ret.String())
	})
}

func VersionCreateForm(rc *fasthttp.RequestCtx) {
	act("version.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &version.Version{}
		ps.Title = "Create [Version]"
		ps.Data = ret
		return render(rc, as, &vversion.Edit{Model: ret, IsNew: true}, ps, "version", "Create")
	})
}

func VersionCreateFormRandom(rc *fasthttp.RequestCtx) {
	act("version.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := version.Random()
		ps.Title = "Create Random Version"
		ps.Data = ret
		return render(rc, as, &vversion.Edit{Model: ret, IsNew: true}, ps, "version", "Create")
	})
}

func VersionCreate(rc *fasthttp.RequestCtx) {
	act("version.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := versionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Version from form")
		}
		err = as.Services.Version.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Version")
		}
		msg := fmt.Sprintf("Version [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func VersionEditForm(rc *fasthttp.RequestCtx) {
	act("version.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := versionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return render(rc, as, &vversion.Edit{Model: ret}, ps, "version", ret.String())
	})
}

func VersionEdit(rc *fasthttp.RequestCtx) {
	act("version.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		rc.SetUserValue("includeDeleted", true)
		ret, err := versionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := versionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Version from form")
		}
		frm.ID = ret.ID
		err = as.Services.Version.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Version [%s]", frm.String())
		}
		msg := fmt.Sprintf("Version [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func VersionDelete(rc *fasthttp.RequestCtx) {
	act("version.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := versionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Version.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete version [%s]", ret.String())
		}
		msg := fmt.Sprintf("Version [%s] deleted", ret.String())
		return flashAndRedir(true, msg, "/version", rc, ps)
	})
}

func versionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*version.Version, error) {
	idArg, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.Version.Get(ps.Context, nil, idArg, ps.Logger)
}

func versionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*version.Version, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return version.FromMap(frm, setPK)
}
