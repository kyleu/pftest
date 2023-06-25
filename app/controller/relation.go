// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vrelation"
)

func RelationList(rc *fasthttp.RequestCtx) {
	Act("relation.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("relation", nil, ps.Logger).Sanitize("relation")
		var ret relation.Relations
		var err error
		if q == "" {
			ret, err = as.Services.Relation.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Relation.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Relations"
		ps.Data = ret
		basicIDsByBasicID := lo.Map(ret, func(x *relation.Relation, _ int) uuid.UUID {
			return x.BasicID
		})
		basicsByBasicID, err := as.Services.Basic.GetMultiple(ps.Context, nil, ps.Logger, basicIDsByBasicID...)
		if err != nil {
			return "", err
		}
		page := &vrelation.List{Models: ret, BasicsByBasicID: basicsByBasicID, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "relation")
	})
}

func RelationDetail(rc *fasthttp.RequestCtx) {
	Act("relation.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Relation)"
		ps.Data = ret

		basicByBasicID, _ := as.Services.Basic.Get(ps.Context, nil, ret.BasicID, ps.Logger)

		return Render(rc, as, &vrelation.Detail{Model: ret, BasicByBasicID: basicByBasicID}, ps, "relation", ret.String())
	})
}

func RelationCreateForm(rc *fasthttp.RequestCtx) {
	Act("relation.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &relation.Relation{}
		ps.Title = "Create [Relation]"
		ps.Data = ret
		return Render(rc, as, &vrelation.Edit{Model: ret, IsNew: true}, ps, "relation", "Create")
	})
}

func RelationCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("relation.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := relation.Random()
		ps.Title = "Create Random Relation"
		ps.Data = ret
		return Render(rc, as, &vrelation.Edit{Model: ret, IsNew: true}, ps, "relation", "Create")
	})
}

func RelationCreate(rc *fasthttp.RequestCtx) {
	Act("relation.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Relation from form")
		}
		err = as.Services.Relation.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Relation")
		}
		msg := fmt.Sprintf("Relation [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func RelationEditForm(rc *fasthttp.RequestCtx) {
	Act("relation.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vrelation.Edit{Model: ret}, ps, "relation", ret.String())
	})
}

func RelationEdit(rc *fasthttp.RequestCtx) {
	Act("relation.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := relationFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Relation from form")
		}
		frm.ID = ret.ID
		err = as.Services.Relation.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Relation [%s]", frm.String())
		}
		msg := fmt.Sprintf("Relation [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func RelationDelete(rc *fasthttp.RequestCtx) {
	Act("relation.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Relation.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete relation [%s]", ret.String())
		}
		msg := fmt.Sprintf("Relation [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/relation", rc, ps)
	})
}

func relationFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*relation.Relation, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Relation.Get(ps.Context, nil, idArg, ps.Logger)
}

func relationFromForm(rc *fasthttp.RequestCtx, setPK bool) (*relation.Relation, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return relation.FromMap(frm, setPK)
}
