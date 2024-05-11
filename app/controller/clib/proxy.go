// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views"
)

func ProxyIndex(w http.ResponseWriter, r *http.Request) {
	controller.Act("proxy.index", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := as.Services.Proxy.List()
		ps.SetTitleAndData("Proxy", ret)
		return controller.Render(r, as, &views.Debug{}, ps)
	})
}

func ProxyHandle(w http.ResponseWriter, r *http.Request) {
	controller.Act("proxy.handle", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		svc, _ := cutil.PathString(r, "svc", true)
		pth, _ := cutil.PathString(r, "path", true)
		return "", as.Services.Proxy.Handle(ps.Context, svc, w, r, pth, ps.Logger)
	})
}
