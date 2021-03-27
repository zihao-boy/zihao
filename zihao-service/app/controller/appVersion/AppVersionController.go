package appVersion

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/appService/service"
)

type AppVersionController struct{
	appVersionService service.AppVersionService
	appVersionJobService service.AppVersionJobService
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