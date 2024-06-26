package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cg1/cg2"
)

func generatedRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/troub/le", controller.TroubleList)
	makeRoute(r, http.MethodGet, "/troub/le/_new", controller.TroubleCreateForm)
	makeRoute(r, http.MethodPost, "/troub/le/_new", controller.TroubleCreate)
	makeRoute(r, http.MethodGet, "/troub/le/_random", controller.TroubleRandom)
	makeRoute(r, http.MethodGet, "/troub/le/{from}/{where}", controller.TroubleDetail)
	makeRoute(r, http.MethodGet, "/troub/le/{from}/{where}/edit", controller.TroubleEditForm)
	makeRoute(r, http.MethodPost, "/troub/le/{from}/{where}/edit", controller.TroubleEdit)
	makeRoute(r, http.MethodGet, "/troub/le/{from}/{where}/delete", controller.TroubleDelete)
	makeRoute(r, http.MethodGet, "/timestamp", controller.TimestampList)
	makeRoute(r, http.MethodGet, "/timestamp/_new", controller.TimestampCreateForm)
	makeRoute(r, http.MethodPost, "/timestamp/_new", controller.TimestampCreate)
	makeRoute(r, http.MethodGet, "/timestamp/_random", controller.TimestampRandom)
	makeRoute(r, http.MethodGet, "/timestamp/{id}", controller.TimestampDetail)
	makeRoute(r, http.MethodGet, "/timestamp/{id}/edit", controller.TimestampEditForm)
	makeRoute(r, http.MethodPost, "/timestamp/{id}/edit", controller.TimestampEdit)
	makeRoute(r, http.MethodGet, "/timestamp/{id}/delete", controller.TimestampDelete)
	makeRoute(r, http.MethodGet, "/softdel", controller.SoftdelList)
	makeRoute(r, http.MethodGet, "/softdel/_new", controller.SoftdelCreateForm)
	makeRoute(r, http.MethodPost, "/softdel/_new", controller.SoftdelCreate)
	makeRoute(r, http.MethodGet, "/softdel/_random", controller.SoftdelRandom)
	makeRoute(r, http.MethodGet, "/softdel/{id}", controller.SoftdelDetail)
	makeRoute(r, http.MethodGet, "/softdel/{id}/edit", controller.SoftdelEditForm)
	makeRoute(r, http.MethodPost, "/softdel/{id}/edit", controller.SoftdelEdit)
	makeRoute(r, http.MethodGet, "/softdel/{id}/delete", controller.SoftdelDelete)
	makeRoute(r, http.MethodGet, "/seed", controller.SeedList)
	makeRoute(r, http.MethodGet, "/seed/_new", controller.SeedCreateForm)
	makeRoute(r, http.MethodPost, "/seed/_new", controller.SeedCreate)
	makeRoute(r, http.MethodGet, "/seed/_random", controller.SeedRandom)
	makeRoute(r, http.MethodGet, "/seed/{id}", controller.SeedDetail)
	makeRoute(r, http.MethodGet, "/seed/{id}/edit", controller.SeedEditForm)
	makeRoute(r, http.MethodPost, "/seed/{id}/edit", controller.SeedEdit)
	makeRoute(r, http.MethodGet, "/seed/{id}/delete", controller.SeedDelete)
	makeRoute(r, http.MethodGet, "/relation", controller.RelationList)
	makeRoute(r, http.MethodGet, "/relation/_new", controller.RelationCreateForm)
	makeRoute(r, http.MethodPost, "/relation/_new", controller.RelationCreate)
	makeRoute(r, http.MethodGet, "/relation/_random", controller.RelationRandom)
	makeRoute(r, http.MethodGet, "/relation/{id}", controller.RelationDetail)
	makeRoute(r, http.MethodGet, "/relation/{id}/edit", controller.RelationEditForm)
	makeRoute(r, http.MethodPost, "/relation/{id}/edit", controller.RelationEdit)
	makeRoute(r, http.MethodGet, "/relation/{id}/delete", controller.RelationDelete)
	makeRoute(r, http.MethodGet, "/reference", controller.ReferenceList)
	makeRoute(r, http.MethodGet, "/reference/_new", controller.ReferenceCreateForm)
	makeRoute(r, http.MethodPost, "/reference/_new", controller.ReferenceCreate)
	makeRoute(r, http.MethodGet, "/reference/_random", controller.ReferenceRandom)
	makeRoute(r, http.MethodGet, "/reference/{id}", controller.ReferenceDetail)
	makeRoute(r, http.MethodGet, "/reference/{id}/edit", controller.ReferenceEditForm)
	makeRoute(r, http.MethodPost, "/reference/{id}/edit", controller.ReferenceEdit)
	makeRoute(r, http.MethodGet, "/reference/{id}/delete", controller.ReferenceDelete)
	makeRoute(r, http.MethodGet, "/mixedcase", controller.MixedCaseList)
	makeRoute(r, http.MethodGet, "/mixedcase/_new", controller.MixedCaseCreateForm)
	makeRoute(r, http.MethodPost, "/mixedcase/_new", controller.MixedCaseCreate)
	makeRoute(r, http.MethodGet, "/mixedcase/_random", controller.MixedCaseRandom)
	makeRoute(r, http.MethodGet, "/mixedcase/{id}", controller.MixedCaseDetail)
	makeRoute(r, http.MethodGet, "/mixedcase/{id}/edit", controller.MixedCaseEditForm)
	makeRoute(r, http.MethodPost, "/mixedcase/{id}/edit", controller.MixedCaseEdit)
	makeRoute(r, http.MethodGet, "/mixedcase/{id}/delete", controller.MixedCaseDelete)
	makeRoute(r, http.MethodGet, "/g1/g2/path", cg2.PathList)
	makeRoute(r, http.MethodGet, "/g1/g2/path/_new", cg2.PathCreateForm)
	makeRoute(r, http.MethodPost, "/g1/g2/path/_new", cg2.PathCreate)
	makeRoute(r, http.MethodGet, "/g1/g2/path/_random", cg2.PathRandom)
	makeRoute(r, http.MethodGet, "/g1/g2/path/{id}", cg2.PathDetail)
	makeRoute(r, http.MethodGet, "/g1/g2/path/{id}/edit", cg2.PathEditForm)
	makeRoute(r, http.MethodPost, "/g1/g2/path/{id}/edit", cg2.PathEdit)
	makeRoute(r, http.MethodGet, "/g1/g2/path/{id}/delete", cg2.PathDelete)
	makeRoute(r, http.MethodGet, "/capital", controller.CapitalList)
	makeRoute(r, http.MethodGet, "/capital/_new", controller.CapitalCreateForm)
	makeRoute(r, http.MethodPost, "/capital/_new", controller.CapitalCreate)
	makeRoute(r, http.MethodGet, "/capital/_random", controller.CapitalRandom)
	makeRoute(r, http.MethodGet, "/capital/{id}", controller.CapitalDetail)
	makeRoute(r, http.MethodGet, "/capital/{id}/edit", controller.CapitalEditForm)
	makeRoute(r, http.MethodPost, "/capital/{id}/edit", controller.CapitalEdit)
	makeRoute(r, http.MethodGet, "/capital/{id}/delete", controller.CapitalDelete)
	makeRoute(r, http.MethodGet, "/basic", controller.BasicList)
	makeRoute(r, http.MethodGet, "/basic/_new", controller.BasicCreateForm)
	makeRoute(r, http.MethodPost, "/basic/_new", controller.BasicCreate)
	makeRoute(r, http.MethodGet, "/basic/_random", controller.BasicRandom)
	makeRoute(r, http.MethodGet, "/basic/{id}", controller.BasicDetail)
	makeRoute(r, http.MethodGet, "/basic/{id}/edit", controller.BasicEditForm)
	makeRoute(r, http.MethodPost, "/basic/{id}/edit", controller.BasicEdit)
	makeRoute(r, http.MethodGet, "/basic/{id}/delete", controller.BasicDelete)
	makeRoute(r, http.MethodGet, "/audited", controller.AuditedList)
	makeRoute(r, http.MethodGet, "/audited/_new", controller.AuditedCreateForm)
	makeRoute(r, http.MethodPost, "/audited/_new", controller.AuditedCreate)
	makeRoute(r, http.MethodGet, "/audited/_random", controller.AuditedRandom)
	makeRoute(r, http.MethodGet, "/audited/{id}", controller.AuditedDetail)
	makeRoute(r, http.MethodGet, "/audited/{id}/edit", controller.AuditedEditForm)
	makeRoute(r, http.MethodPost, "/audited/{id}/edit", controller.AuditedEdit)
	makeRoute(r, http.MethodGet, "/audited/{id}/delete", controller.AuditedDelete)
}
