// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/pftest/app/controller/clib"
)

func notebookRoutes(r *router.Router) {
	r.GET("/notebook", clib.Notebook)
	r.GET("/notebook/action/{act}", clib.NotebookAction)
}
