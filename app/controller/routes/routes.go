// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/clib"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
)

func makeRoute(x *mux.Router, method string, path string, f http.HandlerFunc) {
	cutil.AddRoute(method, path)
	x.HandleFunc(path, f).Methods(method)
}

//nolint:revive
func AppRoutes(as *app.State, logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, "/", controller.Home)
	makeRoute(r, http.MethodGet, "/healthcheck", clib.Healthcheck)
	makeRoute(r, http.MethodGet, "/about", clib.About)

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.Profile)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)
	makeRoute(r, http.MethodGet, "/auth/{key}", clib.AuthDetail)
	makeRoute(r, http.MethodGet, "/auth/callback/{key}", clib.AuthCallback)
	makeRoute(r, http.MethodGet, "/auth/logout/{key}", clib.AuthLogout)
	makeRoute(r, http.MethodGet, cutil.DefaultSearchPath, clib.Search)

	themeRoutes(r)
	generatedRoutes(r)

	// $PF_SECTION_START(routes)$
	notebookRoutes(r)
	harRoutes(r)
	// $PF_SECTION_END(routes)$

	makeRoute(r, http.MethodGet, "/docs", clib.Docs)
	makeRoute(r, http.MethodGet, "/docs/{path:.*}", clib.Docs)

	makeRoute(r, http.MethodGet, "/graphql", controller.GraphQLIndex)
	makeRoute(r, http.MethodGet, "/graphql/{key}", controller.GraphQLDetail)
	makeRoute(r, http.MethodPost, "/graphql/{key}", controller.GraphQLRun)

	makeRoute(r, http.MethodGet, "/admin", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/", clib.Admin)
	makeRoute(r, http.MethodPost, "/admin/", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/audit", clib.AuditList)
	makeRoute(r, http.MethodGet, "/admin/audit/random", clib.AuditCreateFormRandom)
	makeRoute(r, http.MethodGet, "/admin/audit/new", clib.AuditCreateForm)
	makeRoute(r, http.MethodPost, "/admin/audit/new", clib.AuditCreate)
	makeRoute(r, http.MethodGet, "/admin/audit/record/{id}/view", clib.RecordDetail)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}", clib.AuditDetail)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}/edit", clib.AuditEditForm)
	makeRoute(r, http.MethodPost, "/admin/audit/{id}/edit", clib.AuditEdit)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}/delete", clib.AuditDelete)
	makeRoute(r, http.MethodGet, "/admin/database", clib.DatabaseList)
	makeRoute(r, http.MethodGet, "/admin/database/{key}", clib.DatabaseDetail)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/{act}", clib.DatabaseAction)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/tables/{schema}/{table}", clib.DatabaseTableView)
	makeRoute(r, http.MethodPost, "/admin/database/{key}/sql", clib.DatabaseSQLRun)
	makeRoute(r, http.MethodGet, "/admin/sandbox", controller.SandboxList)
	makeRoute(r, http.MethodGet, "/admin/sandbox/{key}", controller.SandboxRun)
	execRoutes(r)
	scriptingRoutes(r)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodOptions, "/", controller.Options)
	r.HandleFunc("/", controller.NotFoundAction)

	return cutil.WireRouter(r, logger)
}
