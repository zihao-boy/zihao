package appVersion

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/appService/service"
)

type AppVersionController struct {
	appVersionService     service.AppVersionService
	appVersionJobService  service.AppVersionJobService
	appVersionAttrService service.AppVersionAttrService
}

func AppVersionControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/appVersion")
		aus      = AppVersionController{appVersionService: service.AppVersionService{}}
	)
	adinMenu.Get("/getAppVersion", hero.Handler(aus.getAppVersion))

	adinMenu.Post("/saveAppVersion", hero.Handler(aus.saveAppVersion))

	adinMenu.Post("/updateAppVersion", hero.Handler(aus.updateAppVersion))

	adinMenu.Post("/deleteAppVersion", hero.Handler(aus.deleteAppVersion))

	adinMenu.Get("/getAppVersionAttr", hero.Handler(aus.getAppVersionAttr))

	adinMenu.Get("/getAppVersionJob", hero.Handler(aus.getAppVersionJob))

	adinMenu.Post("/saveAppVersionJob", hero.Handler(aus.saveAppVersionJob))

	adinMenu.Post("/updateAppVersionJob", hero.Handler(aus.updateAppVersionJob))

	adinMenu.Post("/deleteAppVersionJob", hero.Handler(aus.deleteAppVersionJob))

	adinMenu.Post("/doJob", hero.Handler(aus.doJob))

	adinMenu.Post("/doJobHook", hero.Handler(aus.doJobHook))

	adinMenu.Get("/getAppVersionJobImages", hero.Handler(aus.getAppVersionJobImages))

	adinMenu.Post("/saveAppVersionJobImages", hero.Handler(aus.saveAppVersionJobImages))

	adinMenu.Post("/updateAppVersionJobImages", hero.Handler(aus.updateAppVersionJobImages))

	adinMenu.Post("/deleteAppVersionJobImages", hero.Handler(aus.deleteAppVersionJobImages))

	adinMenu.Get("/getJobLog", hero.Handler(aus.getJobLog))

	//web hooks
	adinMenu.Post("/payload", hero.Handler(aus.payload))

}

/**
查询 主机组
*/
func (aus *AppVersionController) getAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.GetAppVersions(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppVersionController) saveAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.SaveAppVersions(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppVersionController) updateAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.UpdateAppVersions(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppVersionController) deleteAppVersion(ctx iris.Context) {
	reslut := aus.appVersionService.DeleteAppVersions(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppVersionController) getAppVersionAttr(ctx iris.Context) {
	reslut := aus.appVersionAttrService.GetAppVersionAttrs(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppVersionController) getAppVersionJob(ctx iris.Context) {
	reslut := aus.appVersionJobService.GetAppVersionJobs(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppVersionController) saveAppVersionJob(ctx iris.Context) {
	reslut := aus.appVersionJobService.SaveAppVersionJobs(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppVersionController) updateAppVersionJob(ctx iris.Context) {
	reslut := aus.appVersionJobService.UpdateAppVersionJobs(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppVersionController) deleteAppVersionJob(ctx iris.Context) {
	reslut := aus.appVersionJobService.DeleteAppVersionJobs(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppVersionController) doJob(ctx iris.Context) {
	reslut := aus.appVersionJobService.DoJob(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppVersionController) doJobHook(ctx iris.Context) {
	reslut := aus.appVersionJobService.DoJobHook(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppVersionController) getAppVersionJobImages(ctx iris.Context) {
	reslut := aus.appVersionJobService.GetAppVersionJobImages(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppVersionController) saveAppVersionJobImages(ctx iris.Context) {
	reslut := aus.appVersionJobService.SaveAppVersionJobImages(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppVersionController) updateAppVersionJobImages(ctx iris.Context) {
	reslut := aus.appVersionJobService.UpdateAppVersionJobImages(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppVersionController) deleteAppVersionJobImages(ctx iris.Context) {
	reslut := aus.appVersionJobService.DeleteAppVersionJobImages(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppVersionController) getJobLog(ctx iris.Context) {
	reslut := aus.appVersionJobService.GetJobLog(ctx)

	ctx.JSON(reslut)
}

// webhooks
func (aus *AppVersionController) payload(ctx iris.Context) {
	reslut := aus.appVersionJobService.Payload(ctx)

	ctx.JSON(reslut)
}
