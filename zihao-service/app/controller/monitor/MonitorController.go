package monitor

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/zihao-boy/zihao/zihao-service/common/crontab"
	"github.com/zihao-boy/zihao/zihao-service/monitor/service"
)



type MonitorController struct{
	monitorHostService service.MonitorHostService
	monitorHostGroupService service.MonitorHostGroupService
}


func MonitorControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/monitor")
		aus      = MonitorController{
		monitorHostService: service.MonitorHostService{},
			monitorHostGroupService: service.MonitorHostGroupService{},
		}
	)

	//查询sql
	adinUser.Get("/getMonitorHosts", hero.Handler(aus.getMonitorHosts))

	//保存sql
	adinUser.Post("/saveMonitorHost", hero.Handler(aus.saveMonitorHost))

	//保存sql
	adinUser.Post("/updateMonitorHost", hero.Handler(aus.updateMonitorHost))

	//保存sql
	adinUser.Post("/deleteMonitorHost", hero.Handler(aus.deleteMonitorHost))

	//查询sql
	adinUser.Get("/getMonitorHostGroups", hero.Handler(aus.getMonitorHostGroups))

	//保存sql
	adinUser.Post("/saveMonitorHostGroup", hero.Handler(aus.saveMonitorHostGroup))

	//保存sql
	adinUser.Post("/updateMonitorHostGroup", hero.Handler(aus.updateMonitorHostGroup))

	//保存sql
	adinUser.Post("/deleteMonitorHostGroup", hero.Handler(aus.deleteMonitorHostGroup))

	//保存sql
	adinUser.Post("/startMonitorHostGroup", hero.Handler(aus.startMonitorHostGroup))


	//保存sql
	adinUser.Post("/stopMonitorHostGroup", hero.Handler(aus.stopMonitorHostGroup))
}



func (aus *MonitorController) getMonitorHosts(ctx iris.Context) {
	relustDto := aus.monitorHostService.GetMonitorHosts(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) saveMonitorHost(ctx iris.Context) {
	relustDto := aus.monitorHostService.SaveMonitorHosts(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) updateMonitorHost(ctx iris.Context) {
	relustDto := aus.monitorHostService.UpdateMonitorHosts(ctx)
	ctx.JSON(relustDto)
}


/**
保存sql信息
*/
func (aus *MonitorController) deleteMonitorHost(ctx iris.Context) {
	relustDto := aus.monitorHostService.DeleteMonitorHosts(ctx)
	ctx.JSON(relustDto)
}


func (aus *MonitorController) getMonitorHostGroups(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.GetMonitorHostGroups(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) saveMonitorHostGroup(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.SaveMonitorHostGroups(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) updateMonitorHostGroup(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.UpdateMonitorHostGroups(ctx)
	ctx.JSON(relustDto)
}


/**
保存sql信息
*/
func (aus *MonitorController) deleteMonitorHostGroup(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.DeleteMonitorHostGroups(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) startMonitorHostGroup(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.StartMonitorHostGroups(ctx)
	var (
		monitorJob = crontab.MonitorJob{}
	)
	monitorJob.Restart()
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) stopMonitorHostGroup(ctx iris.Context) {
	relustDto := aus.monitorHostGroupService.StopMonitorHostGroups(ctx)

	var (
		monitorJob = crontab.MonitorJob{}
	)
	monitorJob.Restart()
	ctx.JSON(relustDto)
}

