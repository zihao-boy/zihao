package appService

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/appService/service"
)

type AppServiceController struct {
	appServiceService  service.AppServiceService
	appVarGroupService service.AppVarGroupService
	appVarService      service.AppVarService
}

func AppServiceControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/appService")
		aus      = AppServiceController{appServiceService: service.AppServiceService{}}
	)
	adinMenu.Get("/getAppService", hero.Handler(aus.getAppService))

	adinMenu.Post("/saveAppService", hero.Handler(aus.saveAppService))

	adinMenu.Post("/updateAppService", hero.Handler(aus.updateAppService))
	adinMenu.Post("/upgradeAppService", hero.Handler(aus.upgradeAppService))



	adinMenu.Post("/deleteAppService", hero.Handler(aus.deleteAppService))

	adinMenu.Get("/getAppVarGroup", hero.Handler(aus.getAppVarGroup))

	adinMenu.Post("/saveAppVarGroup", hero.Handler(aus.saveAppVarGroup))

	adinMenu.Post("/updateAppVarGroup", hero.Handler(aus.updateAppVarGroup))

	adinMenu.Post("/deleteAppVarGroup", hero.Handler(aus.deleteAppVarGroup))

	adinMenu.Get("/getAppVar", hero.Handler(aus.getAppVar))

	adinMenu.Post("/saveAppVar", hero.Handler(aus.saveAppVar))

	adinMenu.Post("/updateAppVar", hero.Handler(aus.updateAppVar))

	adinMenu.Post("/deleteAppVar", hero.Handler(aus.deleteAppVar))

	adinMenu.Get("/getAppServiceVar", hero.Handler(aus.getAppServiceVar))

	adinMenu.Post("/saveAppServiceVar", hero.Handler(aus.saveAppServiceVar))

	adinMenu.Post("/updateAppServiceVar", hero.Handler(aus.updateAppServiceVar))

	adinMenu.Post("/deleteAppServiceVar", hero.Handler(aus.deleteAppServiceVar))

	adinMenu.Post("/deleteAppVar", hero.Handler(aus.deleteAppVar))

	adinMenu.Get("/getAppServiceHosts", hero.Handler(aus.getAppServiceHosts))

	adinMenu.Post("/saveAppServiceHosts", hero.Handler(aus.saveAppServiceHosts))

	adinMenu.Post("/updateAppServiceHosts", hero.Handler(aus.updateAppServiceHosts))

	adinMenu.Post("/deleteAppServiceHosts", hero.Handler(aus.deleteAppServiceHosts))

	adinMenu.Get("/getAppServiceDir", hero.Handler(aus.getAppServiceDir))

	adinMenu.Post("/saveAppServiceDir", hero.Handler(aus.saveAppServiceDir))

	adinMenu.Post("/updateAppServiceDir", hero.Handler(aus.updateAppServiceDir))

	adinMenu.Post("/deleteAppServiceDir", hero.Handler(aus.deleteAppServiceDir))

	adinMenu.Get("/getAppServicePort", hero.Handler(aus.getAppServicePort))

	adinMenu.Post("/saveAppServicePort", hero.Handler(aus.saveAppServicePort))

	adinMenu.Post("/updateAppServicePort", hero.Handler(aus.updateAppServicePort))

	adinMenu.Post("/deleteAppServicePort", hero.Handler(aus.deleteAppServicePort))

	adinMenu.Get("/getAppServiceContainer", hero.Handler(aus.getAppServiceContainer))

	adinMenu.Post("/saveAppServiceContainer", hero.Handler(aus.saveAppServiceContainer))

	adinMenu.Post("/updateAppServiceContainer", hero.Handler(aus.updateAppServiceContainer))

	adinMenu.Post("/deleteAppServiceContainer", hero.Handler(aus.deleteAppServiceContainer))

	adinMenu.Post("/startAppService", hero.Handler(aus.startAppService))

	adinMenu.Post("/stopAppService", hero.Handler(aus.stopAppService))
	adinMenu.Post("/restartAppServices", hero.Handler(aus.restartAppServices))



	adinMenu.Post("/copyAppService", hero.Handler(aus.copyAppService))

	//get faster deploy app service
	adinMenu.Get("/getFasterDeploy", hero.Handler(aus.getFasterDeploy))

	adinMenu.Post("/saveFasterDeploy", hero.Handler(aus.saveFasterDeploy))

	adinMenu.Post("/updateFasterDeploy", hero.Handler(aus.updateFasterDeploy))

	adinMenu.Post("/deleteFasterDeploy", hero.Handler(aus.deleteFasterDeploy))

	// export app config
	adinMenu.Get("/exportAppService", hero.Handler(aus.exportAppService))

	// import app config
	adinMenu.Post("/importAppService", hero.Handler(aus.importAppService))
}

/**
get app service
*/
func (aus *AppServiceController) getAppService(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServices(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppService(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServices(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppService(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServices(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) upgradeAppService(ctx iris.Context) {
	reslut := aus.appServiceService.UpgradeAppService(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) copyAppService(ctx iris.Context) {
	reslut := aus.appServiceService.CopyAppServices(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppService(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServices(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.GetAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.SaveAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.UpdateAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppVarGroup(ctx iris.Context) {
	reslut := aus.appVarGroupService.DeleteAppVarGroups(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppVar(ctx iris.Context) {
	reslut := aus.appVarService.GetAppVars(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppVar(ctx iris.Context) {
	reslut := aus.appVarService.SaveAppVars(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppVar(ctx iris.Context) {
	reslut := aus.appVarService.UpdateAppVars(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppVar(ctx iris.Context) {
	reslut := aus.appVarService.DeleteAppVars(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceVar(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceVars(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceHosts(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceHosts(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceDir(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceDir(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceDir(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceDir(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceDir(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceDir(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceDir(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceDir(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppServicePort(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServicePort(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServicePort(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServicePort(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServicePort(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServicePort(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServicePort(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServicePort(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *AppServiceController) getAppServiceContainer(ctx iris.Context) {
	reslut := aus.appServiceService.GetAppServiceContainer(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveAppServiceContainer(ctx iris.Context) {
	reslut := aus.appServiceService.SaveAppServiceContainer(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateAppServiceContainer(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateAppServiceContainer(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteAppServiceContainer(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteAppServiceContainer(ctx)

	ctx.JSON(reslut)
}

/**
开启容器
*/
func (aus *AppServiceController) startAppService(ctx iris.Context) {
	reslut := aus.appServiceService.StartAppService(ctx)

	ctx.JSON(reslut)
}
/**
 restart apps
*/
func (aus *AppServiceController) restartAppServices(ctx iris.Context) {
	reslut := aus.appServiceService.RestartAppServices(ctx)

	ctx.JSON(reslut)
}


/**
停止容器
*/
func (aus *AppServiceController) stopAppService(ctx iris.Context) {
	reslut := aus.appServiceService.StopAppService(ctx)

	ctx.JSON(reslut)
}

// get faster deploy app service log
func (aus *AppServiceController) getFasterDeploy(ctx iris.Context) {
	reslut := aus.appServiceService.GetFasterDeploys(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *AppServiceController) saveFasterDeploy(ctx iris.Context) {
	reslut := aus.appServiceService.SaveFasterDeploys(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *AppServiceController) updateFasterDeploy(ctx iris.Context) {
	reslut := aus.appServiceService.UpdateFasterDeploys(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *AppServiceController) deleteFasterDeploy(ctx iris.Context) {
	reslut := aus.appServiceService.DeleteFasterDeploys(ctx)

	ctx.JSON(reslut)
}

/**
export app service
*/
func (aus *AppServiceController) exportAppService(ctx iris.Context) {
	aus.appServiceService.ExportAppService(ctx)
}

/**
export app service
*/
func (aus *AppServiceController) importAppService(ctx iris.Context) {
	reslut := aus.appServiceService.ImportAppService(ctx)
	ctx.JSON(reslut)
}
