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
	proxyRoutes("/proxy", r)
	// $PF_SECTION_END(routes)$

	makeRoute(r, http.MethodGet, "/docs", clib.Docs)
	makeRoute(r, http.MethodGet, "/docs/{path:.*}", clib.Docs)

	makeRoute(r, http.MethodGet, "/graphql", clib.GraphQLIndex)
	makeRoute(r, http.MethodGet, "/graphql/{key}", clib.GraphQLDetail)
	makeRoute(r, http.MethodPost, "/graphql/{key}", clib.GraphQLRun)

	execRoutes(r)
	scriptingRoutes(r)
	adminRoutes(r)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/robots.txt", clib.RobotsTxt)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodOptions, "/", controller.Options)

	return cutil.WireRouter(r, defaultHandler, logger)
}
