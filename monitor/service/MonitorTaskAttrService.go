package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/monitor/dao"
)

type MonitorTaskAttrService struct {
	monitorTaskAttrDao dao.MonitorTaskAttrDao
}

/**
查询 系统信息
*/
func (monitorTaskAttrService *MonitorTaskAttrService) GetMonitorTaskAttrAll(monitorTaskAttrDto monitor.MonitorTaskAttrDto) ([]*monitor.MonitorTaskAttrDto, error) {
	var (
		err                 error
		monitorTaskAttrDtos []*monitor.MonitorTaskAttrDto
	)

	monitorTaskAttrDtos, err = monitorTaskAttrService.monitorTaskAttrDao.GetMonitorTaskAttrs(monitorTaskAttrDto)
	if err != nil {
		return nil, err
	}

	return monitorTaskAttrDtos, nil

}

/**
查询 系统信息
*/
func (monitorTaskAttrService *MonitorTaskAttrService) GetMonitorTaskAttrs(ctx iris.Context) result.ResultDto {

	var (
		err                 error
		monitorTaskAttrDto  = monitor.MonitorTaskAttrDto{}
		monitorTaskAttrDtos []*monitor.MonitorTaskAttrDto
	)

	monitorTaskAttrDto.TaskId = ctx.URLParam("taskId")

	monitorTaskAttrDtos, err = monitorTaskAttrService.monitorTaskAttrDao.GetMonitorTaskAttrs(monitorTaskAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskAttrDtos)

}

/**
保存 系统信息
*/
func (monitorTaskAttrService *MonitorTaskAttrService) SaveMonitorTaskAttrs(ctx iris.Context) result.ResultDto {
	var (
		err                error
		monitorTaskAttrDto monitor.MonitorTaskAttrDto
	)

	if err = ctx.ReadJSON(&monitorTaskAttrDto); err != nil {
		return result.Error("解析入参失败")
	}
	monitorTaskAttrDto.TaskId = seq.Generator()

	err = monitorTaskAttrService.monitorTaskAttrDao.SaveMonitorTaskAttr(monitorTaskAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskAttrDto)

}

/**
修改 系统信息
*/
func (monitorTaskAttrService *MonitorTaskAttrService) UpdateMonitorTaskAttrs(ctx iris.Context) result.ResultDto {
	var (
		err                error
		monitorTaskAttrDto monitor.MonitorTaskAttrDto
	)

	if err = ctx.ReadJSON(&monitorTaskAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorTaskAttrService.monitorTaskAttrDao.UpdateMonitorTaskAttr(monitorTaskAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskAttrDto)

}

/**
删除 系统信息
*/
func (monitorTaskAttrService *MonitorTaskAttrService) DeleteMonitorTaskAttrs(ctx iris.Context) result.ResultDto {
	var (
		err                error
		monitorTaskAttrDto monitor.MonitorTaskAttrDto
	)

	if err = ctx.ReadJSON(&monitorTaskAttrDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorTaskAttrService.monitorTaskAttrDao.DeleteMonitorTaskAttr(monitorTaskAttrDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskAttrDto)

}
