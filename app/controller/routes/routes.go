// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/clib"
	"github.com/kyleu/pftest/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/pftest/app/util"
)

//nolint
func AppRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Home)
	r.GET("/healthcheck", clib.Healthcheck)
	r.GET("/about", clib.About)
	r.GET("/theme", clib.ThemeList)
	r.GET("/theme/{key}", clib.ThemeEdit)
	r.POST("/theme/{key}", clib.ThemeSave)
	r.GET(controller.DefaultSearchPath, clib.Search)

	r.GET(controller.DefaultProfilePath, clib.Profile)
	r.POST(controller.DefaultProfilePath, clib.ProfileSave)
	r.GET("/auth/{key}", clib.AuthDetail)
	r.GET("/auth/callback/{key}", clib.AuthCallback)
	r.GET("/auth/logout/{key}", clib.AuthLogout)

	// $PF_INJECT_START(codegen)$
	r.GET("/capital", controller.CapitalList)
	r.GET("/capital/random", controller.CapitalCreateFormRandom)
	r.GET("/capital/new", controller.CapitalCreateForm)
	r.POST("/capital/new", controller.CapitalCreate)
	r.GET("/capital/{id}", controller.CapitalDetail)
	r.GET("/capital/{id}/edit", controller.CapitalEditForm)
	r.POST("/capital/{id}/edit", controller.CapitalEdit)
	r.GET("/capital/{id}/delete", controller.CapitalDelete)
	r.GET("/capital/{id}/Version/{Version}", controller.CapitalVersion)
	r.GET("/audited", controller.AuditedList)
	r.GET("/audited/random", controller.AuditedCreateFormRandom)
	r.GET("/audited/new", controller.AuditedCreateForm)
	r.POST("/audited/new", controller.AuditedCreate)
	r.GET("/audited/{id}", controller.AuditedDetail)
	r.GET("/audited/{id}/edit", controller.AuditedEditForm)
	r.POST("/audited/{id}/edit", controller.AuditedEdit)
	r.GET("/audited/{id}/delete", controller.AuditedDelete)
	r.GET("/basic", controller.BasicList)
	r.GET("/basic/random", controller.BasicCreateFormRandom)
	r.GET("/basic/new", controller.BasicCreateForm)
	r.POST("/basic/new", controller.BasicCreate)
	r.GET("/basic/{id}", controller.BasicDetail)
	r.GET("/basic/{id}/edit", controller.BasicEditForm)
	r.POST("/basic/{id}/edit", controller.BasicEdit)
	r.GET("/basic/{id}/delete", controller.BasicDelete)
	r.GET("/group/group", controller.GroupGroupList)
	r.GET("/group/group/{group}", controller.GroupListByGroup)
	r.GET("/group/group/{group}/new", controller.GroupCreateFormByGroup)
	r.POST("/group/group/{group}/new", controller.GroupCreateByGroup)
	r.GET("/group/group/{group}/{id}", controller.GroupDetailByGroup)
	r.GET("/group/group/{group}/{id}/edit", controller.GroupEditFormByGroup)
	r.POST("/group/group/{group}/{id}/edit", controller.GroupEditByGroup)
	r.GET("/group", controller.GroupList)
	r.GET("/group/random", controller.GroupCreateFormRandom)
	r.GET("/group/new", controller.GroupCreateForm)
	r.POST("/group/new", controller.GroupCreate)
	r.GET("/group/{id}", controller.GroupDetail)
	r.GET("/group/{id}/edit", controller.GroupEditForm)
	r.POST("/group/{id}/edit", controller.GroupEdit)
	r.GET("/group/{id}/delete", controller.GroupDelete)
	r.GET("/history", controller.HistoryList)
	r.GET("/history/random", controller.HistoryCreateFormRandom)
	r.GET("/history/new", controller.HistoryCreateForm)
	r.POST("/history/new", controller.HistoryCreate)
	r.GET("/history/{id}", controller.HistoryDetail)
	r.GET("/history/{id}/edit", controller.HistoryEditForm)
	r.POST("/history/{id}/edit", controller.HistoryEdit)
	r.GET("/history/{id}/delete", controller.HistoryDelete)
	r.GET("/history/{id}/history/{historyID}", controller.HistoryHistory)
	r.GET("/mixedcase", controller.MixedCaseList)
	r.GET("/mixedcase/random", controller.MixedCaseCreateFormRandom)
	r.GET("/mixedcase/new", controller.MixedCaseCreateForm)
	r.POST("/mixedcase/new", controller.MixedCaseCreate)
	r.GET("/mixedcase/{id}", controller.MixedCaseDetail)
	r.GET("/mixedcase/{id}/edit", controller.MixedCaseEditForm)
	r.POST("/mixedcase/{id}/edit", controller.MixedCaseEdit)
	r.GET("/mixedcase/{id}/delete", controller.MixedCaseDelete)
	r.GET("/reference", controller.ReferenceList)
	r.GET("/reference/random", controller.ReferenceCreateFormRandom)
	r.GET("/reference/new", controller.ReferenceCreateForm)
	r.POST("/reference/new", controller.ReferenceCreate)
	r.GET("/reference/{id}", controller.ReferenceDetail)
	r.GET("/reference/{id}/edit", controller.ReferenceEditForm)
	r.POST("/reference/{id}/edit", controller.ReferenceEdit)
	r.GET("/reference/{id}/delete", controller.ReferenceDelete)
	r.GET("/relation", controller.RelationList)
	r.GET("/relation/random", controller.RelationCreateFormRandom)
	r.GET("/relation/new", controller.RelationCreateForm)
	r.POST("/relation/new", controller.RelationCreate)
	r.GET("/relation/{id}", controller.RelationDetail)
	r.GET("/relation/{id}/edit", controller.RelationEditForm)
	r.POST("/relation/{id}/edit", controller.RelationEdit)
	r.GET("/relation/{id}/delete", controller.RelationDelete)
	r.GET("/softdel", controller.SoftdelList)
	r.GET("/softdel/random", controller.SoftdelCreateFormRandom)
	r.GET("/softdel/new", controller.SoftdelCreateForm)
	r.POST("/softdel/new", controller.SoftdelCreate)
	r.GET("/softdel/{id}", controller.SoftdelDetail)
	r.GET("/softdel/{id}/edit", controller.SoftdelEditForm)
	r.POST("/softdel/{id}/edit", controller.SoftdelEdit)
	r.GET("/softdel/{id}/delete", controller.SoftdelDelete)
	r.GET("/timestamp", controller.TimestampList)
	r.GET("/timestamp/random", controller.TimestampCreateFormRandom)
	r.GET("/timestamp/new", controller.TimestampCreateForm)
	r.POST("/timestamp/new", controller.TimestampCreate)
	r.GET("/timestamp/{id}", controller.TimestampDetail)
	r.GET("/timestamp/{id}/edit", controller.TimestampEditForm)
	r.POST("/timestamp/{id}/edit", controller.TimestampEdit)
	r.GET("/timestamp/{id}/delete", controller.TimestampDelete)
	r.GET("/troub/le", controller.TroubleList)
	r.GET("/troub/le/random", controller.TroubleCreateFormRandom)
	r.GET("/troub/le/new", controller.TroubleCreateForm)
	r.POST("/troub/le/new", controller.TroubleCreate)
	r.GET("/troub/le/{from}/{where}", controller.TroubleDetail)
	r.GET("/troub/le/{from}/{where}/edit", controller.TroubleEditForm)
	r.POST("/troub/le/{from}/{where}/edit", controller.TroubleEdit)
	r.GET("/troub/le/{from}/{where}/delete", controller.TroubleDelete)
	r.GET("/troub/le/{from}/{where}/selectcol/{selectcol}", controller.TroubleSelectcol)
	r.GET("/version", controller.VersionList)
	r.GET("/version/random", controller.VersionCreateFormRandom)
	r.GET("/version/new", controller.VersionCreateForm)
	r.POST("/version/new", controller.VersionCreate)
	r.GET("/version/{id}", controller.VersionDetail)
	r.GET("/version/{id}/edit", controller.VersionEditForm)
	r.POST("/version/{id}/edit", controller.VersionEdit)
	r.GET("/version/{id}/delete", controller.VersionDelete)
	r.GET("/version/{id}/revision/{revision}", controller.VersionRevision)
	// $PF_INJECT_END(codegen)$

	// $PF_SECTION_START(routes)$
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
	r.GET("/admin/{path:*}", clib.Admin)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/robots.txt", clib.RobotsTxt)
	r.GET("/assets/{_:*}", clib.Static)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFound

	clib.AppRoutesList = r.List()

	p := httpmetrics.NewMetrics(util.AppKey)
	return fasthttp.CompressHandlerLevel(p.WrapHandler(r), fasthttp.CompressBestSpeed)
}