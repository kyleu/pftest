// Content managed by Project Forge, see [projectforge.md] for details.
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

func GroupChildList(rc *fasthttp.RequestCtx) {
	Act("group.child.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "[Groups] by child"
		ret, err := as.Services.Group.GetChildren(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Data = ret
		return Render(rc, as, &vgroup.Children{Children: ret}, ps, "group", "child")
	})
}

func GroupListByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		prms := ps.Params.Get("group", nil, ps.Logger).Sanitize("group")
		ret, err := as.Services.Group.GetByChild(ps.Context, nil, childArg, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Groups"
		ps.Data = ret
		page := &vgroup.List{Models: ret, Params: ps.Params}
		return Render(rc, as, page, ps, "group", "child")
	})
}

func GroupDetailByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		ps.Title = ret.TitleString() + " (Group)"
		ps.Data = ret
		return Render(rc, as, &vgroup.Detail{Model: ret}, ps, "group", "child", ret.String())
	})
}

func GroupCreateFormByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret := &group.Group{Child: childArg}
		ps.Title = fmt.Sprintf("Create [Group] for child [%s]", childArg)
		ps.Data = ret
		return Render(rc, as, &vgroup.Edit{Model: ret, IsNew: true}, ps, "group", "child", "Create")
	})
}

func GroupCreateByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret, err := groupFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		if ret.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		err = as.Services.Group.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Group")
		}
		msg := fmt.Sprintf("Group [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func GroupEditFormByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vgroup.Edit{Model: ret}, ps, "group", "child", ret.String())
	})
}

func GroupEditByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		frm, err := groupFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Group from form")
		}
		if frm.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		frm.ID = ret.ID
		err = as.Services.Group.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Group [%s]", frm.String())
		}
		msg := fmt.Sprintf("Group [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func GroupDeleteByChild(rc *fasthttp.RequestCtx) {
	Act("group.child.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		childArg, err := cutil.RCRequiredString(rc, "child", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [child] as an argument")
		}
		ret, err := groupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		if ret.Child != childArg {
			return "", errors.New("unauthorized: incorrect [child]")
		}
		err = as.Services.Group.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete group [%s]", ret.String())
		}
		msg := fmt.Sprintf("Group [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/group", rc, ps)
	})
}
