package system

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/entity/result"
	"github.com/zihao-boy/zihao/zihao-service/system/service"
)



type SystemController struct{
	systenInfoService service.SystemInfoService
}


func SystemControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/system")
		aus      = SystemController{systenInfoService: service.SystemInfoService{}}
	)
	adinUser.Get("/info", hero.Handler(aus.info))
}

func (aus *SystemController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}