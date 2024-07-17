package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cg1/cg2"
)

const routeNew, routeRandom, routeEdit, routeDelete = "/_new", "/_random", "/edit", "/delete"

func generatedRoutes(r *mux.Router) {
	generatedRoutesTrouble(r, "/troub/le")
	generatedRoutesTimestamp(r, "/timestamp")
	generatedRoutesSoftdel(r, "/softdel")
	generatedRoutesSeed(r, "/seed")
	generatedRoutesRelation(r, "/relation")
	generatedRoutesReference(r, "/reference")
	generatedRoutesMixedCase(r, "/mixedcase")
	generatedRoutesPath(r, "/g1/g2/path")
	generatedRoutesCapital(r, "/capital")
	generatedRoutesBasic(r, "/basic")
	generatedRoutesAudited(r, "/audited")
}

func generatedRoutesTrouble(r *mux.Router, prefix string) {
	const pkn = "/{from}/{where}"
	makeRoute(r, http.MethodGet, prefix, controller.TroubleList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.TroubleCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.TroubleCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.TroubleRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.TroubleDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.TroubleEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.TroubleEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.TroubleDelete)
}

func generatedRoutesTimestamp(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.TimestampList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.TimestampCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.TimestampCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.TimestampRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.TimestampDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.TimestampEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.TimestampEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.TimestampDelete)
}

func generatedRoutesSoftdel(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.SoftdelList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.SoftdelCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.SoftdelCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.SoftdelRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.SoftdelDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.SoftdelEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.SoftdelEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.SoftdelDelete)
}

func generatedRoutesSeed(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.SeedList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.SeedCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.SeedCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.SeedRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.SeedDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.SeedEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.SeedEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.SeedDelete)
}

func generatedRoutesRelation(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.RelationList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.RelationCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.RelationCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.RelationRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.RelationDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.RelationEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.RelationEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.RelationDelete)
}

func generatedRoutesReference(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.ReferenceList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.ReferenceCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.ReferenceCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.ReferenceRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.ReferenceDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.ReferenceEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.ReferenceEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.ReferenceDelete)
}

func generatedRoutesMixedCase(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.MixedCaseList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.MixedCaseCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.MixedCaseCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.MixedCaseRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.MixedCaseDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.MixedCaseEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.MixedCaseEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.MixedCaseDelete)
}

func generatedRoutesPath(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, cg2.PathList)
	makeRoute(r, http.MethodGet, prefix+routeNew, cg2.PathCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, cg2.PathCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, cg2.PathRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, cg2.PathDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, cg2.PathEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, cg2.PathEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, cg2.PathDelete)
}

func generatedRoutesCapital(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.CapitalList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.CapitalCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.CapitalCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.CapitalRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.CapitalDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.CapitalEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.CapitalEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.CapitalDelete)
}

func generatedRoutesBasic(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.BasicList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.BasicCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.BasicCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.BasicRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.BasicDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.BasicEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.BasicEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.BasicDelete)
}

func generatedRoutesAudited(r *mux.Router, prefix string) {
	const pkn = "/{id}"
	makeRoute(r, http.MethodGet, prefix, controller.AuditedList)
	makeRoute(r, http.MethodGet, prefix+routeNew, controller.AuditedCreateForm)
	makeRoute(r, http.MethodPost, prefix+routeNew, controller.AuditedCreate)
	makeRoute(r, http.MethodGet, prefix+routeRandom, controller.AuditedRandom)
	makeRoute(r, http.MethodGet, prefix+pkn, controller.AuditedDetail)
	makeRoute(r, http.MethodGet, prefix+pkn+routeEdit, controller.AuditedEditForm)
	makeRoute(r, http.MethodPost, prefix+pkn+routeEdit, controller.AuditedEdit)
	makeRoute(r, http.MethodGet, prefix+pkn+routeDelete, controller.AuditedDelete)
}
