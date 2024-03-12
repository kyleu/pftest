// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/notebook"
	"github.com/kyleu/pftest/views/vnotebook"
)

var notebookSvc *notebook.Service

func Notebook(rc *fasthttp.RequestCtx) {
	controller.Act("notebook", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		if notebookSvc == nil {
			notebookSvc = notebook.NewService()
		}
		status := notebookSvc.Status()
		if status == "running" {
			ps.SetTitleAndData("Notebook", "view-in-browser")
			return controller.Render(rc, as, &vnotebook.Notebook{}, ps, "notebook")
		}
		ps.SetTitleAndData("Notebook Options", status)
		return controller.Render(rc, as, &vnotebook.Options{}, ps, "notebook", "Options")
	})
}

func NotebookAction(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		act, err := cutil.RCRequiredString(rc, "act", false)
		if err != nil {
			return "", err
		}
		switch act {
		case "start":
			err = notebookSvc.Start(as.Services.Exec)
			return controller.FlashAndRedir(true, "Notebook started", "/notebook", rc, ps)
		default:
			return "", errors.Errorf("invalid notebook action [%s]", act)
		}
	})
}
