// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views/vscripting"
)

func ScriptingList(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Scripting"
		ret := as.Services.Script.ListScripts(ps.Logger)
		ps.Data = ret
		return controller.Render(rc, as, &vscripting.List{Scripts: ret}, ps, "scripting")
	})
}

func ScriptingDetail(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		src, err := as.Services.Script.LoadScript(key, ps.Logger)
		if err != nil {
			return "", err
		}
		vm, err := as.Services.Script.LoadVM(src)
		if err != nil {
			return "", err
		}
		res, err := as.Services.Script.RunExamples(vm)
		if err != nil {
			return "", err
		}
		ps.Title = "Scripting"
		ps.Data = map[string]any{"script": src, "results": res}
		return controller.Render(rc, as, &vscripting.Detail{Path: key, Script: src, Results: res}, ps, "scripting", key)
	})
}

func ScriptingNew(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.new", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "New Script"
		return controller.Render(rc, as, &vscripting.Form{}, ps, "scripting", "New")
	})
}

func ScriptingCreate(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		pth := frm.GetStringOpt("path")
		if pth == "" {
			return "", errors.New("must provide path")
		}
		content := frm.GetStringOpt("content")
		if content == "" {
			return "", errors.New("content is required")
		}
		err = as.Services.Script.SaveScript(pth, content, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Scripting created", "/admin/scripting/"+url.QueryEscape(pth), rc, ps)
	})
}

func ScriptingForm(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		sc, err := as.Services.Script.LoadScript(key, ps.Logger)
		ps.Title = "Edit [" + key + "]"
		ps.Data = sc
		return controller.Render(rc, as, &vscripting.Form{Path: key, Content: sc}, ps, "scripting", key, "Edit")
	})
}

func ScriptingSave(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		content := frm.GetStringOpt("content")
		if content == "" {
			return "", errors.New("content is required")
		}
		err = as.Services.Script.SaveScript(key, content, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Scripting saved", "/admin/scripting/"+url.QueryEscape(key), rc, ps)
	})
}

func ScriptingDelete(rc *fasthttp.RequestCtx) {
	controller.Act("scripting.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", true)
		if err != nil {
			return "", err
		}
		err = as.Services.Script.DeleteScript(key, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Script deleted", "/admin/scripting", rc, ps)
	})
}
