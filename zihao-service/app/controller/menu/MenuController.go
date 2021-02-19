package menu

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/system/service"
)

type MenuController struct{
	menuService service.MenuService
}


func MenuControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/menu")
		aus      = MenuController{menuService: service.MenuService{}}
	)
	adinMenu.Get("/getMenus", hero.Handler(aus.info))
}

func (aus *MenuController) info(ctx iris.Context) {
	tempMenus := aus.menuService.GetMenus(ctx)

	ctx.JSON(tempMenus)
}
