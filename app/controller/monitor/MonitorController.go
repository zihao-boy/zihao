package monitor

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	logTraceService "github.com/zihao-boy/zihao/business/service/logTrace"
	"github.com/zihao-boy/zihao/common/crontab"
	"github.com/zihao-boy/zihao/monitor/service"
)

type MonitorController struct {
	monitorHostService      service.MonitorHostService
	monitorHostGroupService service.MonitorHostGroupService
	monitorHostLogService   service.MonitorHostLogService
	monitorEventService     service.MonitorEventService
	monitorTaskService      service.MonitorTaskService
	monitorTaskAttrService  service.MonitorTaskAttrService
	logTraceService logTraceService.LogTraceService
}

func MonitorControllerRouter(party iris.Party) {
	var (
		adinUser = party.Party("/monitor")
		aus      = MonitorController{
			monitorHostService:      service.MonitorHostService{},
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

	//查询监控日志
	adinUser.Get("/getMonitorHostLog", hero.Handler(aus.getMonitorHostLog))

	//查询监控事件
	adinUser.Get("/getMonitorEvents", hero.Handler(aus.getMonitorEvents))

	//查询sql
	adinUser.Get("/getMonitorTasks", hero.Handler(aus.getMonitorTasks))

	//保存sql
	adinUser.Post("/saveMonitorTask", hero.Handler(aus.saveMonitorTask))

	//保存sql
	adinUser.Post("/updateMonitorTask", hero.Handler(aus.updateMonitorTask))

	//保存sql
	adinUser.Post("/deleteMonitorTask", hero.Handler(aus.deleteMonitorTask))

	//保存sql
	adinUser.Post("/startMonitorTask", hero.Handler(aus.startMonitorTask))

	//保存sql
	adinUser.Post("/stopMonitorTask", hero.Handler(aus.stopMonitorTask))

	//查询sql
	adinUser.Get("/listTaskTemplate", hero.Handler(aus.listTaskTemplate))
	//查询sql
	adinUser.Get("/listTaskAttrs", hero.Handler(aus.listTaskAttrs))

	//查询sql
	adinUser.Get("/getLogTrace", hero.Handler(aus.getLogTrace))
	//查询sql
	adinUser.Get("/getLogTraceDetail", hero.Handler(aus.getLogTraceDetail))

	//查询sql
	adinUser.Get("/getLogTraceParam", hero.Handler(aus.getLogTraceParam))
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

func (aus *MonitorController) getMonitorHostLog(ctx iris.Context) {
	relustDto := aus.monitorHostLogService.GetMonitorHostLogs(ctx)
	ctx.JSON(relustDto)
}

func (aus *MonitorController) getMonitorEvents(ctx iris.Context) {
	relustDto := aus.monitorEventService.GetMonitorEvents(ctx)
	ctx.JSON(relustDto)
}

func (aus *MonitorController) getMonitorTasks(ctx iris.Context) {
	relustDto := aus.monitorTaskService.GetMonitorTasks(ctx)
	ctx.JSON(relustDto)
}



/**
保存sql信息
*/
func (aus *MonitorController) saveMonitorTask(ctx iris.Context) {
	relustDto := aus.monitorTaskService.SaveMonitorTasks(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) updateMonitorTask(ctx iris.Context) {
	relustDto := aus.monitorTaskService.UpdateMonitorTasks(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) deleteMonitorTask(ctx iris.Context) {
	relustDto := aus.monitorTaskService.DeleteMonitorTasks(ctx)
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) startMonitorTask(ctx iris.Context) {
	relustDto := aus.monitorTaskService.StartMonitorTask(ctx)
	var (
		monitorJob = crontab.MonitorJob{}
	)
	monitorJob.Restart()
	ctx.JSON(relustDto)
}

/**
保存sql信息
*/
func (aus *MonitorController) stopMonitorTask(ctx iris.Context) {
	relustDto := aus.monitorTaskService.StopMonitorTasks(ctx)

	var (
		monitorJob = crontab.MonitorJob{}
	)
	monitorJob.Restart()
	ctx.JSON(relustDto)
}

func (aus *MonitorController) listTaskTemplate(ctx iris.Context) {
	relustDto := aus.monitorTaskService.ListTaskTemplate(ctx)
	ctx.JSON(relustDto)
}
func (aus *MonitorController) listTaskAttrs(ctx iris.Context) {
	relustDto := aus.monitorTaskAttrService.GetMonitorTaskAttrs(ctx)
	ctx.JSON(relustDto)
}

func (aus *MonitorController) getLogTrace(ctx iris.Context) {
	relustDto := aus.logTraceService.GetLogTraces(ctx)
	ctx.JSON(relustDto)
}
func (aus *MonitorController) getLogTraceDetail(ctx iris.Context) {
	relustDto := aus.logTraceService.GetLogTraceDetail(ctx)
	ctx.JSON(relustDto)
}

func (aus *MonitorController) getLogTraceParam(ctx iris.Context) {
	relustDto := aus.logTraceService.GetLogTraceParam(ctx)
	ctx.JSON(relustDto)
}



