package controller

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/seed"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vseed"
)

func SeedList(w http.ResponseWriter, r *http.Request) {
	Act("seed.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("seed", ps.Logger)
		ret, err := as.Services.Seed.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Seeds", ret)
		page := &vseed.List{Models: ret, Params: ps.Params}
		return Render(r, as, page, ps, "seed")
	})
}

func SeedDetail(w http.ResponseWriter, r *http.Request) {
	Act("seed.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Seed)", ret)

		return Render(r, as, &vseed.Detail{Model: ret}, ps, "seed", ret.TitleString()+"**star")
	})
}

func SeedCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("seed.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &seed.Seed{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = seed.RandomSeed()
		}
		ps.SetTitleAndData("Create [Seed]", ret)
		return Render(r, as, &vseed.Edit{Model: ret, IsNew: true}, ps, "seed", "Create")
	})
}

func SeedRandom(w http.ResponseWriter, r *http.Request) {
	Act("seed.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Seed.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Seed")
		}
		return ret.WebPath(), nil
	})
}

func SeedCreate(w http.ResponseWriter, r *http.Request) {
	Act("seed.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Seed from form")
		}
		err = as.Services.Seed.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Seed")
		}
		msg := fmt.Sprintf("Seed [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SeedEditForm(w http.ResponseWriter, r *http.Request) {
	Act("seed.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vseed.Edit{Model: ret}, ps, "seed", ret.String())
	})
}

func SeedEdit(w http.ResponseWriter, r *http.Request) {
	Act("seed.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := seedFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Seed from form")
		}
		frm.ID = ret.ID
		err = as.Services.Seed.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Seed [%s]", frm.String())
		}
		msg := fmt.Sprintf("Seed [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SeedDelete(w http.ResponseWriter, r *http.Request) {
	Act("seed.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Seed.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete seed [%s]", ret.String())
		}
		msg := fmt.Sprintf("Seed [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/seed", ps)
	})
}

func seedFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*seed.Seed, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Seed.Get(ps.Context, nil, idArg, ps.Logger)
}

func seedFromForm(r *http.Request, b []byte, setPK bool) (*seed.Seed, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := seed.SeedFromMap(frm, setPK)
	return ret, err
}
