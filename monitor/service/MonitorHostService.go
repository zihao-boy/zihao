package service

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/date"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/monitor/dao"
)

type MonitorHostService struct {
	monitorHostDao dao.MonitorHostDao
}

/**
查询 系统信息
*/
func (monitorHostService *MonitorHostService) GetMonitorHostAll(monitorHostDto monitor.MonitorHostDto) ([]*monitor.MonitorHostDto, error) {
	var (
		err             error
		monitorHostDtos []*monitor.MonitorHostDto
	)

	monitorHostDtos, err = monitorHostService.monitorHostDao.GetMonitorHosts(monitorHostDto)
	if err != nil {
		return nil, err
	}

	return monitorHostDtos, nil

}

/**
查询 系统信息
*/
func (monitorHostService *MonitorHostService) GetMonitorHosts(ctx iris.Context) result.ResultDto {
	var (
		err             error
		page            int64
		row             int64
		total           int64
		monitorHostDto  = monitor.MonitorHostDto{}
		monitorHostDtos []*monitor.MonitorHostDto
	)

	page, err = strconv.ParseInt(ctx.URLParam("page"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	row, err = strconv.ParseInt(ctx.URLParam("row"), 10, 64)

	if err != nil {
		return result.Error(err.Error())
	}

	monitorHostDto.Row = row * page

	monitorHostDto.Page = (page - 1) * row

	total, err = monitorHostService.monitorHostDao.GetMonitorHostCount(monitorHostDto)

	if err != nil {
		return result.Error(err.Error())
	}

	if total < 1 {
		return result.Success()
	}

	monitorHostDtos, err = monitorHostService.monitorHostDao.GetMonitorHosts(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDtos, total, row)

}

/**
保存 系统信息
*/
func (monitorHostService *MonitorHostService) SaveMonitorHosts(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	monitorHostDto.TenantId = user.TenantId
	monitorHostDto.MhId = seq.Generator()
	monitorHostDto.MonDate = date.GetNowTimeString()
	monitorHostDto.CpuRate = "0"
	monitorHostDto.DiskRate = "0"
	monitorHostDto.FreeDisk = "0"
	monitorHostDto.FreeMem = "0"
	monitorHostDto.MemRate = "0"

	err = monitorHostService.monitorHostDao.SaveMonitorHost(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}

/**
修改 系统信息
*/
func (monitorHostService *MonitorHostService) UpdateMonitorHosts(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostService.monitorHostDao.UpdateMonitorHost(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}

/**
删除 系统信息
*/
func (monitorHostService *MonitorHostService) DeleteMonitorHosts(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostService.monitorHostDao.DeleteMonitorHost(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}
