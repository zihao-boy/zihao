package resources

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/business/service/resourcesBackUpService"
	"github.com/zihao-boy/zihao/business/service/resourcesDbService"
	"github.com/zihao-boy/zihao/business/service/resourcesFtpService"
	"github.com/zihao-boy/zihao/business/service/resourcesOssService"
)

type ResourcesController struct {
	resourcesFtpService    resourcesFtpService.ResourcesFtpService
	resourcesOssService    resourcesOssService.ResourcesOssService
	resourcesDbService     resourcesDbService.ResourcesDbService
	resourcesBackUpService resourcesBackUpService.ResourcesBackUpService
}

func ResourcesControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/resources")
		aus      = ResourcesController{
			resourcesFtpService:    resourcesFtpService.ResourcesFtpService{},
			resourcesOssService:    resourcesOssService.ResourcesOssService{},
			resourcesDbService:     resourcesDbService.ResourcesDbService{},
			resourcesBackUpService: resourcesBackUpService.ResourcesBackUpService{},
		}
	)
	adinMenu.Get("/getFtp", hero.Handler(aus.getFtp))

	adinMenu.Post("/saveFtp", hero.Handler(aus.saveFtp))

	adinMenu.Post("/updateFtp", hero.Handler(aus.updateFtp))

	adinMenu.Post("/deleteFtp", hero.Handler(aus.deleteFtp))

	adinMenu.Get("/getOss", hero.Handler(aus.getOss))

	adinMenu.Post("/saveOss", hero.Handler(aus.saveOss))

	adinMenu.Post("/updateOss", hero.Handler(aus.updateOss))

	adinMenu.Post("/deleteOss", hero.Handler(aus.deleteOss))

	adinMenu.Get("/getDb", hero.Handler(aus.getDb))

	adinMenu.Post("/saveDb", hero.Handler(aus.saveDb))

	adinMenu.Post("/updateDb", hero.Handler(aus.updateDb))

	adinMenu.Post("/deleteDb", hero.Handler(aus.deleteDb))

	adinMenu.Get("/getBackUp", hero.Handler(aus.getBackUp))

	adinMenu.Post("/saveBackUp", hero.Handler(aus.saveBackUp))

	adinMenu.Post("/updateBackUp", hero.Handler(aus.updateBackUp))

	adinMenu.Post("/deleteBackUp", hero.Handler(aus.deleteBackUp))

}

/**
查询 主机组
*/
func (aus *ResourcesController) getFtp(ctx iris.Context) {
	reslut := aus.resourcesFtpService.GetResourcesFtps(ctx)
	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *ResourcesController) saveFtp(ctx iris.Context) {
	reslut := aus.resourcesFtpService.SaveResourcesFtps(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *ResourcesController) updateFtp(ctx iris.Context) {
	reslut := aus.resourcesFtpService.UpdateResourcesFtps(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *ResourcesController) deleteFtp(ctx iris.Context) {
	reslut := aus.resourcesFtpService.DeleteResourcesFtps(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *ResourcesController) getOss(ctx iris.Context) {
	reslut := aus.resourcesOssService.GetResourcesOsss(ctx)
	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *ResourcesController) saveOss(ctx iris.Context) {
	reslut := aus.resourcesOssService.SaveResourcesOsss(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *ResourcesController) updateOss(ctx iris.Context) {
	reslut := aus.resourcesOssService.UpdateResourcesOsss(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *ResourcesController) deleteOss(ctx iris.Context) {
	reslut := aus.resourcesOssService.DeleteResourcesOsss(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *ResourcesController) getDb(ctx iris.Context) {
	reslut := aus.resourcesDbService.GetResourcesDbs(ctx)
	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *ResourcesController) saveDb(ctx iris.Context) {
	reslut := aus.resourcesDbService.SaveResourcesDbs(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *ResourcesController) updateDb(ctx iris.Context) {
	reslut := aus.resourcesDbService.UpdateResourcesDbs(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *ResourcesController) deleteDb(ctx iris.Context) {
	reslut := aus.resourcesDbService.DeleteResourcesDbs(ctx)

	ctx.JSON(reslut)
}

/**
查询 主机组
*/
func (aus *ResourcesController) getBackUp(ctx iris.Context) {
	reslut := aus.resourcesBackUpService.GetResourcesBackUps(ctx)
	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *ResourcesController) saveBackUp(ctx iris.Context) {
	reslut := aus.resourcesBackUpService.SaveResourcesBackUps(ctx)

	ctx.JSON(reslut)
}

/**
修改 主机组
*/
func (aus *ResourcesController) updateBackUp(ctx iris.Context) {
	reslut := aus.resourcesBackUpService.UpdateResourcesBackUps(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *ResourcesController) deleteBackUp(ctx iris.Context) {
	reslut := aus.resourcesBackUpService.DeleteResourcesBackUps(ctx)

	ctx.JSON(reslut)
}


