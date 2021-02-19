package host

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/assets/service"
)

type HostController struct{
	hostService service.HostService
}


func HostControllerRouter(party iris.Party) {
	var (
		adinMenu = party.Party("/host")
		aus      = HostController{hostService: service.HostService{}}
	)
	adinMenu.Get("/getHostGroup", hero.Handler(aus.getHostGroup))
}

/**
查询 主机组
 */
func (aus *HostController) getHostGroup(ctx iris.Context) {
	reslut := aus.hostService.GetHostGroups(ctx)

	ctx.JSON(reslut)
}
