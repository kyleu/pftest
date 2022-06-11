// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views"
)

func About(rc *fasthttp.RequestCtx) {
	controller.Act("about", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "About " + util.AppName
		ps.Data = util.AppName + " v" + as.BuildInfo.Version
		return controller.Render(rc, as, &views.About{}, ps, "about")
	})
}
