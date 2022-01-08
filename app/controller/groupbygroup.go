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

func GroupGroupList(rc *fasthttp.RequestCtx) {
	act("group.group.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Groups"
		ret, err := as.Services.Group.GetGroups(ps.Context, nil)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vgroup.Groups{Groups: ret}, ps, "group", "group")
	})
}

func GroupListByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ps.Title = "Groups"
		params := cutil.ParamSetFromRequest(rc)
		ret, err := as.Services.Group.GetByGroup(ps.Context, nil, groupArg, params.Get("group", nil, ps.Logger))
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return render(rc, as, &vgroup.List{Models: ret, Params: params}, ps, "group", "group")
	})
}

func GroupDetailByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Group != groupArg {
			return "", errors.New("unauthorized: incorrect [group]")
		}
		ps.Title = ret.String()
		ps.Data = ret
		return render(rc, as, &vgroup.Detail{Model: ret}, ps, "group", "group", ret.String())
	})
}

func GroupCreateFormByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ret := &group.Group{Group: groupArg}
		ps.Title = "Create [Group]"
		ps.Data = ret
		return render(rc, as, &vgroup.Edit{Model: ret, IsNew: true}, ps, "group", "group", "Create")
	})
}

func GroupCreateByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ret, err := groupFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		if ret.Group != groupArg {
			return "", errors.New("unauthorized: incorrect [group]")
		}
		err = as.Services.Group.Create(ps.Context, nil, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Group")
		}
		msg := fmt.Sprintf("Group [%s] created", ret.String())
		return flashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func GroupEditFormByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Group != groupArg {
			return "", errors.New("unauthorized: incorrect [group]")
		}
		ps.Title = "Edit [" + ret.String() + "]"
		ps.Data = ret
		return render(rc, as, &vgroup.Edit{Model: ret}, ps, "group", "group", ret.String())
	})
}

func GroupEditByGroup(rc *fasthttp.RequestCtx) {
	act("group.group.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		groupArg, err := rcRequiredString(rc, "group", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [group] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Group != groupArg {
			return "", errors.New("unauthorized: incorrect [group]")
		}
		frm, err := groupFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		if frm.Group != groupArg {
			return "", errors.New("unauthorized: incorrect [group]")
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
