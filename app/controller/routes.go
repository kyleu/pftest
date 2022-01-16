package controller

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/pftest/app/util"
)

//nolint
func AppRoutes() fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", Home)
	r.GET("/healthcheck", Healthcheck)
	r.GET("/about", About)
	r.GET("/theme", ThemeList)
	r.GET("/theme/{key}", ThemeEdit)
	r.POST("/theme/{key}", ThemeSave)
	r.GET(defaultSearchPath, Search)

	r.GET(defaultProfilePath, Profile)
	r.POST(defaultProfilePath, ProfileSave)
	r.GET("/auth/{key}", AuthDetail)
	r.GET("/auth/callback/{key}", AuthCallback)
	r.GET("/auth/logout/{key}", AuthLogout)

	// $PF_INJECT_START(codegen)$
	r.GET("/basic", BasicList)
	r.GET("/basic/random", BasicCreateFormRandom)
	r.GET("/basic/new", BasicCreateForm)
	r.POST("/basic/new", BasicCreate)
	r.GET("/basic/{id}", BasicDetail)
	r.GET("/basic/{id}/edit", BasicEditForm)
	r.POST("/basic/{id}/edit", BasicEdit)
	r.GET("/basic/{id}/delete", BasicDelete)
	r.GET("/timestamp", TimestampList)
	r.GET("/timestamp/random", TimestampCreateFormRandom)
	r.GET("/timestamp/new", TimestampCreateForm)
	r.POST("/timestamp/new", TimestampCreate)
	r.GET("/timestamp/{id}", TimestampDetail)
	r.GET("/timestamp/{id}/edit", TimestampEditForm)
	r.POST("/timestamp/{id}/edit", TimestampEdit)
	r.GET("/timestamp/{id}/delete", TimestampDelete)
	r.GET("/version", VersionList)
	r.GET("/version/random", VersionCreateFormRandom)
	r.GET("/version/new", VersionCreateForm)
	r.POST("/version/new", VersionCreate)
	r.GET("/version/{id}", VersionDetail)
	r.GET("/version/{id}/edit", VersionEditForm)
	r.POST("/version/{id}/edit", VersionEdit)
	r.GET("/version/{id}/delete", VersionDelete)
	r.GET("/version/{id}/revision/{revision}", VersionRevision)
	r.GET("/softdel", SoftdelList)
	r.GET("/softdel/random", SoftdelCreateFormRandom)
	r.GET("/softdel/new", SoftdelCreateForm)
	r.POST("/softdel/new", SoftdelCreate)
	r.GET("/softdel/{id}", SoftdelDetail)
	r.GET("/softdel/{id}/edit", SoftdelEditForm)
	r.POST("/softdel/{id}/edit", SoftdelEdit)
	r.GET("/softdel/{id}/delete", SoftdelDelete)
	r.GET("/group/group", GroupGroupList)
	r.GET("/group/group/{group}", GroupListByGroup)
	r.GET("/group/group/{group}/new", GroupCreateFormByGroup)
	r.POST("/group/group/{group}/new", GroupCreateByGroup)
	r.GET("/group/group/{group}/{id}", GroupDetailByGroup)
	r.GET("/group/group/{group}/{id}/edit", GroupEditFormByGroup)
	r.POST("/group/group/{group}/{id}/edit", GroupEditByGroup)

	r.GET("/group", GroupList)
	r.GET("/group/random", GroupCreateFormRandom)
	r.GET("/group/new", GroupCreateForm)
	r.POST("/group/new", GroupCreate)
	r.GET("/group/{id}", GroupDetail)
	r.GET("/group/{id}/edit", GroupEditForm)
	r.POST("/group/{id}/edit", GroupEdit)
	r.GET("/group/{id}/delete", GroupDelete)
	r.GET("/mixed_case", MixedCaseList)
	r.GET("/mixed_case/random", MixedCaseCreateFormRandom)
	r.GET("/mixed_case/new", MixedCaseCreateForm)
	r.POST("/mixed_case/new", MixedCaseCreate)
	r.GET("/mixed_case/{id}", MixedCaseDetail)
	r.GET("/mixed_case/{id}/edit", MixedCaseEditForm)
	r.POST("/mixed_case/{id}/edit", MixedCaseEdit)
	r.GET("/mixed_case/{id}/delete", MixedCaseDelete)
	r.GET("/trouble", TroubleList)
	r.GET("/trouble/random", TroubleCreateFormRandom)
	r.GET("/trouble/new", TroubleCreateForm)
	r.POST("/trouble/new", TroubleCreate)
	r.GET("/trouble/{from}/{where}", TroubleDetail)
	r.GET("/trouble/{from}/{where}/edit", TroubleEditForm)
	r.POST("/trouble/{from}/{where}/edit", TroubleEdit)
	r.GET("/trouble/{from}/{where}/delete", TroubleDelete)
	r.GET("/trouble/{from}/{where}/selectcol/{selectcol}", TroubleSelectcol)
	// $PF_INJECT_END(codegen)$

	// $PF_SECTION_START(routes)$
	// $PF_SECTION_END(routes)$

	r.GET("/admin", Admin)
	r.GET("/admin/sandbox", SandboxList)
	r.GET("/admin/sandbox/{key}", SandboxRun)
	r.GET("/admin/{path:*}", Admin)

	r.GET("/favicon.ico", Favicon)
	r.GET("/robots.txt", RobotsTxt)
	r.GET("/assets/{_:*}", Static)

	r.OPTIONS("/", Options)
	r.OPTIONS("/{_:*}", Options)
	r.NotFound = NotFound

	p := httpmetrics.NewMetrics(util.AppKey)
	return fasthttp.CompressHandlerBrotliLevel(p.WrapHandler(r), fasthttp.CompressBrotliBestSpeed, fasthttp.CompressBestSpeed)
}
