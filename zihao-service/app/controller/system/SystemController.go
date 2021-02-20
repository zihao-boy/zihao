package system

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/system/service"
)



type SystemController struct{
	systenInfoService service.SystemInfoService
	serviceSqlService service.ServiceSqlService
}


func SystemControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/system")
		aus      = SystemController{systenInfoService: service.SystemInfoService{},serviceSqlService: service.ServiceSqlService{}}
	)
	adinUser.Get("/info", hero.Handler(aus.info))

	//查询sql
	adinUser.Get("/getServiceSqls", hero.Handler(aus.getServiceSqls))
}

func (aus *SystemController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}


func (aus *SystemController) getServiceSqls(ctx iris.Context) {
	relustDto := aus.serviceSqlService.GetServiceSqls(ctx)
	ctx.JSON(relustDto)
}