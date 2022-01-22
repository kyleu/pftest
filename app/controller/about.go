// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views"
)

func About(rc *fasthttp.RequestCtx) {
	act("about", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = util.AppName + " v" + as.BuildInfo.Version
		return render(rc, as, &views.About{}, ps, "about")
	})
}
