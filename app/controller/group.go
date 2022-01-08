package controller

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/group"
	"github.com/kyleu/pftest/views/vgroup"
)

func GroupList(rc *fasthttp.RequestCtx) {
	act("group.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Groups"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Group.List(ps.Context, nil, params.Get("group", nil, ps.Logger))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vgroup.List{Models: ret, Params: params}, ps, "group")
	})
}

func GroupDetail(rc *fasthttp.RequestCtx) {
	act("group.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vgroup.Detail{Model: ret}, ps, "group", ret.String())
	})
}

func GroupCreateForm(rc *fasthttp.RequestCtx) {
	act("group.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &group.Group{}
		ps.Title = "Create [Group]"
		ps.Data = ret
		return render(rc, as, &vgroup.Edit{Model: ret, IsNew: true}, ps, "group", "Create")
	})
}

func GroupCreate(rc *fasthttp.RequestCtx) {
	act("group.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := groupFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		err = as.Services.Group.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Group")
		}
		msg := fmt.Sprintf("Group [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func GroupEditForm(rc *fasthttp.RequestCtx) {
	act("group.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vgroup.Edit{Model: ret}, ps, "group", ret.String())
	})
}

func GroupEdit(rc *fasthttp.RequestCtx) {
	act("group.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := groupFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		frm.ID = ret.ID
		err = as.Services.Group.Update(ps.Context, nil, frm)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Group [%s]", frm.String())
		}
		msg := fmt.Sprintf("Group [%s] updated", frm.String())
		return flashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func groupFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*group.Group, error) {
	idArg, err := rcRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	return as.Services.Group.Get(ps.Context, nil, idArg)
}

func groupFromForm(rc *fasthttp.RequestCtx, setPK bool) (*group.Group, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return group.FromMap(frm, setPK)
}
