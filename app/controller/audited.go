// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vaudited"
)

func AuditedList(w http.ResponseWriter, r *http.Request) {
	Act("audited.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
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
		ps.SetTitleAndData("Auditeds", ret)
		page := &vaudited.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(w, r, as, page, ps, "audited")
	})
}

func AuditedDetail(w http.ResponseWriter, r *http.Request) {
	Act("audited.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Audited)", ret)

		relatedAuditRecords, err := as.Services.Audit.RecordsForModel(ps.Context, nil, "audited", ret.ID.String(), nil, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to retrieve related audit records")
		}

		return Render(w, r, as, &vaudited.Detail{Model: ret, AuditRecords: relatedAuditRecords}, ps, "audited", ret.TitleString()+"**star")
	})
}

func AuditedCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("audited.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &audited.Audited{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = audited.Random()
		}
		ps.SetTitleAndData("Create [Audited]", ret)
		ps.Data = ret
		return Render(w, r, as, &vaudited.Edit{Model: ret, IsNew: true}, ps, "audited", "Create")
	})
}

func AuditedRandom(w http.ResponseWriter, r *http.Request) {
	Act("audited.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Audited.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Audited")
		}
		return ret.WebPath(), nil
	})
}

func AuditedCreate(w http.ResponseWriter, r *http.Request) {
	Act("audited.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Audited from form")
		}
		err = as.Services.Audited.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Audited")
		}
		msg := fmt.Sprintf("Audited [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func AuditedEditForm(w http.ResponseWriter, r *http.Request) {
	Act("audited.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(w, r, as, &vaudited.Edit{Model: ret}, ps, "audited", ret.String())
	})
}

func AuditedEdit(w http.ResponseWriter, r *http.Request) {
	Act("audited.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := auditedFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Audited from form")
		}
		frm.ID = ret.ID
		err = as.Services.Audited.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Audited [%s]", frm.String())
		}
		msg := fmt.Sprintf("Audited [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func AuditedDelete(w http.ResponseWriter, r *http.Request) {
	Act("audited.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := auditedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Audited.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete audited [%s]", ret.String())
		}
		msg := fmt.Sprintf("Audited [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/audited", w, ps)
	})
}

func auditedFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*audited.Audited, error) {
	idArgStr, err := cutil.RCRequiredString(r, "id", false)
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

func auditedFromForm(r *http.Request, b []byte, setPK bool) (*audited.Audited, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return audited.FromMap(frm, setPK)
}
