package tenant

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/user/service"
)

type TenantController struct {
	tenantService        service.TenantService
	tenantSettingService service.TenantSettingService
}

func TenantControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/tenant")
		aus      = TenantController{
			tenantService: service.TenantService{}}
	)

	//查询sql
	adinUser.Get("/getTenants", hero.Handler(aus.getTenants))

	//保存sql
	adinUser.Post("/saveTenant", hero.Handler(aus.saveTenant))

	//保存sql
	adinUser.Post("/updateTenant", hero.Handler(aus.updateTenant))

	//保存sql
	adinUser.Post("/deleteTenant", hero.Handler(aus.deleteTenant))

	//查询sql
	adinUser.Get("/getTenantSettings", hero.Handler(aus.getTenantSettings))

	//保存sql
	adinUser.Post("/saveTenantSetting", hero.Handler(aus.saveTenantSetting))

	//保存sql
	adinUser.Post("/updateTenantSetting", hero.Handler(aus.updateTenantSetting))

	//保存sql
	adinUser.Post("/deleteTenantSetting", hero.Handler(aus.deleteTenantSetting))
}

func (aus *TenantController) getTenants(ctx iris.Context) {
	relustDto := aus.tenantService.GetTenants(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) saveTenant(ctx iris.Context) {
	relustDto := aus.tenantService.SaveTenants(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) updateTenant(ctx iris.Context) {
	relustDto := aus.tenantService.UpdateTenants(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) deleteTenant(ctx iris.Context) {
	relustDto := aus.tenantService.DeleteTenants(ctx)
	ctx.JSON(relustDto)
}

func (aus *TenantController) getTenantSettings(ctx iris.Context) {
	relustDto := aus.tenantSettingService.GetTenantSettings(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) saveTenantSetting(ctx iris.Context) {
	relustDto := aus.tenantSettingService.SaveTenantSettings(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) updateTenantSetting(ctx iris.Context) {
	relustDto := aus.tenantSettingService.UpdateTenantSettings(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *TenantController) deleteTenantSetting(ctx iris.Context) {
	relustDto := aus.tenantSettingService.DeleteTenantSettings(ctx)
	ctx.JSON(relustDto)
}
