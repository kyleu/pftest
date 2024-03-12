// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/clib"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/pftest/app/util"
)

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Home)
	r.GET("/healthcheck", clib.Healthcheck)
	r.GET("/about", clib.About)

	r.GET(cutil.DefaultProfilePath, clib.Profile)
	r.POST(cutil.DefaultProfilePath, clib.ProfileSave)
	r.GET("/auth/{key}", clib.AuthDetail)
	r.GET("/auth/callback/{key}", clib.AuthCallback)
	r.GET("/auth/logout/{key}", clib.AuthLogout)
	r.GET(cutil.DefaultSearchPath, clib.Search)
	themeRoutes(r)
	generatedRoutes(r)

	// $PF_SECTION_START(routes)$
	notebookRoutes(r)
	harRoutes(r)
	// $PF_SECTION_END(routes)$

	r.GET("/docs", clib.Docs)
	r.GET("/docs/{path:*}", clib.Docs)

	r.GET("/graphql", controller.GraphQLIndex)
	r.GET("/graphql/{key}", controller.GraphQLDetail)
	r.POST("/graphql/{key}", controller.GraphQLRun)

	r.GET("/admin", clib.Admin)
	r.GET("/admin/audit", clib.AuditList)
	r.GET("/admin/audit/random", clib.AuditCreateFormRandom)
	r.GET("/admin/audit/new", clib.AuditCreateForm)
	r.POST("/admin/audit/new", clib.AuditCreate)
	r.GET("/admin/audit/record/{id}", clib.RecordDetail)
	r.GET("/admin/audit/{id}", clib.AuditDetail)
	r.GET("/admin/audit/{id}/edit", clib.AuditEditForm)
	r.POST("/admin/audit/{id}/edit", clib.AuditEdit)
	r.GET("/admin/audit/{id}/delete", clib.AuditDelete)
	r.GET("/admin/database", clib.DatabaseList)
	r.GET("/admin/database/{key}", clib.DatabaseDetail)
	r.GET("/admin/database/{key}/{act}", clib.DatabaseAction)
	r.GET("/admin/database/{key}/tables/{schema}/{table}", clib.DatabaseTableView)
	r.POST("/admin/database/{key}/sql", clib.DatabaseSQLRun)
	r.GET("/admin/sandbox", controller.SandboxList)
	r.GET("/admin/sandbox/{key}", controller.SandboxRun)
	execRoutes(r)
	scriptingRoutes(r)
	r.GET("/admin/{path:*}", clib.Admin)
	r.POST("/admin/{path:*}", clib.Admin)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/robots.txt", clib.RobotsTxt)
	r.GET("/assets/{_:*}", clib.Static)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFound

	clib.AppRoutesList = r.List()

	p := httpmetrics.NewMetrics(util.AppKey, logger)
	return fasthttp.CompressHandlerLevel(p.WrapHandler(r, true), fasthttp.CompressBestSpeed)
}
