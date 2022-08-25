// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/seed"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vseed"
)

func SeedList(rc *fasthttp.RequestCtx) {
	Act("seed.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("seed", nil, ps.Logger).Sanitize("seed")
		ret, err := as.Services.Seed.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Seeds"
		ps.Data = ret
		return Render(rc, as, &vseed.List{Models: ret, Params: params}, ps, "seed")
	})
}

func SeedDetail(rc *fasthttp.RequestCtx) {
	Act("seed.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Seed)"
		ps.Data = ret
		return Render(rc, as, &vseed.Detail{Model: ret}, ps, "seed", ret.String())
	})
}

func SeedCreateForm(rc *fasthttp.RequestCtx) {
	Act("seed.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &seed.Seed{}
		ps.Title = "Create [Seed]"
		ps.Data = ret
		return Render(rc, as, &vseed.Edit{Model: ret, IsNew: true}, ps, "seed", "Create")
	})
}

func SeedCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("seed.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := seed.Random()
		ps.Title = "Create Random Seed"
		ps.Data = ret
		return Render(rc, as, &vseed.Edit{Model: ret, IsNew: true}, ps, "seed", "Create")
	})
}

func SeedCreate(rc *fasthttp.RequestCtx) {
	Act("seed.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Seed from form")
		}
		err = as.Services.Seed.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Seed")
		}
		msg := fmt.Sprintf("Seed [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SeedEditForm(rc *fasthttp.RequestCtx) {
	Act("seed.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vseed.Edit{Model: ret}, ps, "seed", ret.String())
	})
}

func SeedEdit(rc *fasthttp.RequestCtx) {
	Act("seed.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := seedFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Seed from form")
		}
		frm.ID = ret.ID
		err = as.Services.Seed.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Seed [%s]", frm.String())
		}
		msg := fmt.Sprintf("Seed [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SeedDelete(rc *fasthttp.RequestCtx) {
	Act("seed.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := seedFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Seed.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete seed [%s]", ret.String())
		}
		msg := fmt.Sprintf("Seed [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/seed", rc, ps)
	})
}

func seedFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*seed.Seed, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func seedFromForm(rc *fasthttp.RequestCtx, setPK bool) (*seed.Seed, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return seed.FromMap(frm, setPK)
}
