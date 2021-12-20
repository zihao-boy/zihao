package home

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/home/service"
)

type HomeController struct {
	homeService service.HomeService
}

func HomeControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/home")
		aus      = HomeController{homeService: service.HomeService{}}
	)
	//query platform data
	adinUser.Get("/platformData", hero.Handler(aus.platformData))
}

//query platform data
// 2021-12-20
func (aus *HomeController) platformData(ctx iris.Context) {
	resultDto := aus.homeService.PlatformData(ctx)
	ctx.JSON(resultDto)
}
