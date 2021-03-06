package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
)

type MonitorEventService struct {
	monitorEventDao dao.MonitorEventDao

}

/**
查询 系统信息
*/
func (monitorEventService *MonitorEventService) GetMonitorEventAll(monitorEventDto monitor.MonitorEventDto)  ([]*monitor.MonitorEventDto,error) {
	var (
		err       error
		monitorEventDtos []*monitor.MonitorEventDto
	)

	monitorEventDtos,err = monitorEventService.monitorEventDao.GetMonitorEvents(monitorEventDto)
	if(err != nil){
		return nil,err
	}

	return monitorEventDtos,nil

}

/**
查询 系统信息
*/
func (monitorEventService *MonitorEventService) GetMonitorEvents(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorEventDto = monitor.MonitorEventDto{}
		monitorEventDtos []*monitor.MonitorEventDto
	)

	monitorEventDtos,err = monitorEventService.monitorEventDao.GetMonitorEvents(monitorEventDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorEventDtos)

}


/**
保存 系统信息
*/
func (monitorEventService *MonitorEventService) SaveMonitorEvents(eventDto monitor.MonitorEventDto)  error {
	var (
		err       error
	)

	eventDto.EventId = seq.Generator()
	eventDto.State = "000" //待告警
	eventDto.StateRemark="待告警"
	err = monitorEventService.monitorEventDao.SaveMonitorEvent(eventDto)
	if(err != nil){
		return err
	}

	return nil

}


/**
修改 系统信息
*/
func (monitorEventService *MonitorEventService) UpdateMonitorEvents(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorEventDto monitor.MonitorEventDto
	)

	if err = ctx.ReadJSON(&monitorEventDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorEventService.monitorEventDao.UpdateMonitorEvent(monitorEventDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorEventDto)

}


/**
删除 系统信息
*/
func (monitorEventService *MonitorEventService) DeleteMonitorEvents(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorEventDto monitor.MonitorEventDto
	)

	if err = ctx.ReadJSON(&monitorEventDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorEventService.monitorEventDao.DeleteMonitorEvent(monitorEventDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorEventDto)

}