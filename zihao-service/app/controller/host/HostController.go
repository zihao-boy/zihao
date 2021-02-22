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

	adinMenu.Post("/saveHostGroup", hero.Handler(aus.saveHostGroup))

	adinMenu.Post("/updateHostGroup", hero.Handler(aus.updateHostGroup))

	adinMenu.Post("/deleteHostGroup", hero.Handler(aus.deleteHostGroup))

}

/**
查询 主机组
 */
func (aus *HostController) getHostGroup(ctx iris.Context) {
	reslut := aus.hostService.GetHostGroups(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *HostController) saveHostGroup(ctx iris.Context) {
	reslut := aus.hostService.SaveHostGroups(ctx)

	ctx.JSON(reslut)
}


/**
修改 主机组
*/
func (aus *HostController) updateHostGroup(ctx iris.Context) {
	reslut := aus.hostService.UpdateHostGroups(ctx)

	ctx.JSON(reslut)
}


/**
删除 主机组
*/
func (aus *HostController) deleteHostGroup(ctx iris.Context) {
	reslut := aus.hostService.DeleteHostGroups(ctx)

	ctx.JSON(reslut)
}