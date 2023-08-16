// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/pftest/app/controller/clib"
)

func scriptingRoutes(r *router.Router) {
	r.GET("/admin/scripting", clib.ScriptingList)
	r.GET("/admin/scripting/{key}", clib.ScriptingDetail)
}
