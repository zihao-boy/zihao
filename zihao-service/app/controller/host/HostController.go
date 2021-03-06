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

	adinMenu.Get("/getHosts", hero.Handler(aus.getHosts))

	adinMenu.Post("/saveHost", hero.Handler(aus.saveHost))

	adinMenu.Post("/updateHost", hero.Handler(aus.updateHost))

	adinMenu.Post("/deleteHost", hero.Handler(aus.deleteHost))

	adinMenu.Get("/getHostToken", hero.Handler(aus.getHostToken))



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



/**
查询 主机组
*/
func (aus *HostController) getHosts(ctx iris.Context) {
	reslut := aus.hostService.GetHosts(ctx)

	ctx.JSON(reslut)
}

/**
添加 主机组
*/
func (aus *HostController) saveHost(ctx iris.Context) {
	reslut := aus.hostService.SaveHost(ctx)

	ctx.JSON(reslut)
}


/**
修改 主机组
*/
func (aus *HostController) updateHost(ctx iris.Context) {
	reslut := aus.hostService.UpdateHost(ctx)

	ctx.JSON(reslut)
}


/**
删除 主机组
*/
func (aus *HostController) deleteHost(ctx iris.Context) {
	reslut := aus.hostService.DeleteHost(ctx)

	ctx.JSON(reslut)
}

/**
删除 主机组
*/
func (aus *HostController) getHostToken(ctx iris.Context) {
	reslut := aus.hostService.GetHostToken(ctx)

	ctx.JSON(reslut)
}