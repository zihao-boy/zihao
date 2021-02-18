package menu

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/system/service"
)

type MenuController struct{
	systenInfoService service.SystemInfoService
}


func MenuControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/menu")
		aus      = MenuController{systenInfoService: service.SystemInfoService{}}
	)
	adinUser.Get("/getMenus", hero.Handler(aus.info))
}

func (aus *MenuController) info(ctx iris.Context) {
	relustDto := result.SuccessData(aus.systenInfoService.Info(ctx))
	ctx.JSON(relustDto)
}
