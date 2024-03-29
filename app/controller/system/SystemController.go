package system

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/system/service"
)

type SystemController struct {
	systenInfoService service.SystemInfoService
	serviceSqlService service.ServiceSqlService
	mappingService    service.MappingService
}

func SystemControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/system")
		aus      = SystemController{systenInfoService: service.SystemInfoService{},
			serviceSqlService: service.ServiceSqlService{},
			mappingService:    service.MappingService{}}
	)
	adinUser.Get("/info", hero.Handler(aus.info))

	//查询sql
	adinUser.Get("/getServiceSqls", hero.Handler(aus.getServiceSqls))

	//保存sql
	adinUser.Post("/saveServiceSql", hero.Handler(aus.saveServiceSql))

	//保存sql
	adinUser.Post("/updateServiceSql", hero.Handler(aus.updateServiceSql))

	//保存sql
	adinUser.Post("/deleteServiceSql", hero.Handler(aus.deleteServiceSql))

	//查询sql
	adinUser.Get("/getMappings", hero.Handler(aus.getMappings))

	//保存sql
	adinUser.Post("/saveMapping", hero.Handler(aus.saveMapping))

	//保存sql
	adinUser.Post("/updateMapping", hero.Handler(aus.updateMapping))

	//保存sql
	adinUser.Post("/deleteMapping", hero.Handler(aus.deleteMapping))

	//查询sql
	adinUser.Get("/getDicts", hero.Handler(aus.getDicts))
}

func (aus *SystemController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}

func (aus *SystemController) getServiceSqls(ctx iris.Context) {
	relustDto := aus.serviceSqlService.GetServiceSqls(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) saveServiceSql(ctx iris.Context) {
	relustDto := aus.serviceSqlService.SaveServiceSqls(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) updateServiceSql(ctx iris.Context) {
	relustDto := aus.serviceSqlService.UpdateServiceSqls(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) deleteServiceSql(ctx iris.Context) {
	relustDto := aus.serviceSqlService.DeleteServiceSqls(ctx)
	ctx.JSON(relustDto)
}

func (aus *SystemController) getMappings(ctx iris.Context) {
	relustDto := aus.mappingService.GetMappings(ctx)
	ctx.JSON(relustDto)
}

func (aus *SystemController) getDicts(ctx iris.Context) {
	relustDto := aus.mappingService.GetDicts(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) saveMapping(ctx iris.Context) {
	relustDto := aus.mappingService.SaveMappings(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) updateMapping(ctx iris.Context) {
	relustDto := aus.mappingService.UpdateMappings(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *SystemController) deleteMapping(ctx iris.Context) {
	relustDto := aus.mappingService.DeleteMappings(ctx)
	ctx.JSON(relustDto)
}
