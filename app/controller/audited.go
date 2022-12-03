// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vaudited"
)

func AuditedList(rc *fasthttp.RequestCtx) {
	Act("audited.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("audited", nil, ps.Logger).Sanitize("audited")
		var ret audited.Auditeds
		var err error
		if q == "" {
			ret, err = as.Services.Audited.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Audited.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Auditeds"
		ps.Data = ret
		page := &vaudited.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "audited")
	})
}

func AuditedDetail(rc *fasthttp.RequestCtx) {
	Act("audited.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Audited)"
		ps.Data = ret
		return Render(rc, as, &vaudited.Detail{Model: ret}, ps, "audited", ret.String())
	})
}

func AuditedCreateForm(rc *fasthttp.RequestCtx) {
	Act("audited.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &audited.Audited{}
		ps.Title = "Create [Audited]"
		ps.Data = ret
		return Render(rc, as, &vaudited.Edit{Model: ret, IsNew: true}, ps, "audited", "Create")
	})
}

func AuditedCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("audited.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := audited.Random()
		ps.Title = "Create Random Audited"
		ps.Data = ret
		return Render(rc, as, &vaudited.Edit{Model: ret, IsNew: true}, ps, "audited", "Create")
	})
}

func AuditedCreate(rc *fasthttp.RequestCtx) {
	Act("audited.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Audited from form")
		}
		err = as.Services.Audited.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Audited")
		}
		msg := fmt.Sprintf("Audited [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func AuditedEditForm(rc *fasthttp.RequestCtx) {
	Act("audited.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vaudited.Edit{Model: ret}, ps, "audited", ret.String())
	})
}

func AuditedEdit(rc *fasthttp.RequestCtx) {
	Act("audited.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := auditedFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Audited from form")
		}
		frm.ID = ret.ID
		err = as.Services.Audited.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Audited [%s]", frm.String())
		}
		msg := fmt.Sprintf("Audited [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func AuditedDelete(rc *fasthttp.RequestCtx) {
	Act("audited.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Audited.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete audited [%s]", ret.String())
		}
		msg := fmt.Sprintf("Audited [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/audited", rc, ps)
	})
}

func auditedFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*audited.Audited, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Audited.Get(ps.Context, nil, idArg, ps.Logger)
}

func auditedFromForm(rc *fasthttp.RequestCtx, setPK bool) (*audited.Audited, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return audited.FromMap(frm, setPK)
}
