package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vrelation"
)

func RelationList(w http.ResponseWriter, r *http.Request) {
	Act("relation.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("relation", ps.Logger)
		var ret relation.Relations
		var err error
		if q == "" {
			ret, err = as.Services.Relation.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Relation.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Relations", ret)
		basicIDsByBasicID := lo.Map(ret, func(x *relation.Relation, _ int) uuid.UUID {
			return x.BasicID
		})
		basicsByBasicID, err := as.Services.Basic.GetMultiple(ps.Context, nil, nil, ps.Logger, basicIDsByBasicID...)
		if err != nil {
			return "", err
		}
		page := &vrelation.List{Models: ret, BasicsByBasicID: basicsByBasicID, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "relation")
	})
}

func RelationDetail(w http.ResponseWriter, r *http.Request) {
	Act("relation.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Relation)", ret)

		basicByBasicID, _ := as.Services.Basic.Get(ps.Context, nil, ret.BasicID, ps.Logger)

		return Render(r, as, &vrelation.Detail{Model: ret, BasicByBasicID: basicByBasicID}, ps, "relation", ret.TitleString()+"**star")
	})
}

func RelationCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("relation.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &relation.Relation{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = relation.Random()
			randomBasic, err := as.Services.Basic.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomBasic != nil {
				ret.BasicID = randomBasic.ID
			}
		}
		ps.SetTitleAndData("Create [Relation]", ret)
		ps.Data = ret
		return Render(r, as, &vrelation.Edit{Model: ret, IsNew: true}, ps, "relation", "Create")
	})
}

func RelationRandom(w http.ResponseWriter, r *http.Request) {
	Act("relation.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Relation.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Relation")
		}
		return ret.WebPath(), nil
	})
}

func RelationCreate(w http.ResponseWriter, r *http.Request) {
	Act("relation.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Relation from form")
		}
		err = as.Services.Relation.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Relation")
		}
		msg := fmt.Sprintf("Relation [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func RelationEditForm(w http.ResponseWriter, r *http.Request) {
	Act("relation.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vrelation.Edit{Model: ret}, ps, "relation", ret.String())
	})
}

func RelationEdit(w http.ResponseWriter, r *http.Request) {
	Act("relation.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := relationFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Relation from form")
		}
		frm.ID = ret.ID
		err = as.Services.Relation.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Relation [%s]", frm.String())
		}
		msg := fmt.Sprintf("Relation [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func RelationDelete(w http.ResponseWriter, r *http.Request) {
	Act("relation.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := relationFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Relation.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete relation [%s]", ret.String())
		}
		msg := fmt.Sprintf("Relation [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/relation", ps)
	})
}

func relationFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*relation.Relation, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func relationFromForm(r *http.Request, b []byte, setPK bool) (*relation.Relation, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := relation.FromMap(frm, setPK)
	return ret, err
}
