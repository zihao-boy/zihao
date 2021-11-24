package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/common/constants"
	"github.com/zihao-boy/zihao/common/seq"
	"github.com/zihao-boy/zihao/entity/dto/monitor"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"github.com/zihao-boy/zihao/entity/dto/user"
	"github.com/zihao-boy/zihao/monitor/dao"
)

type MonitorHostLogService struct {
	monitorHostDao dao.MonitorHostLogDao
}

/**
查询 系统信息
*/
func (monitorHostService *MonitorHostLogService) GetMonitorHostLogAll(monitorHostDto monitor.MonitorHostLogDto) ([]*monitor.MonitorHostLogDto, error) {
	var (
		err             error
		monitorHostDtos []*monitor.MonitorHostLogDto
	)

	monitorHostDtos, err = monitorHostService.monitorHostDao.GetMonitorHostLogs(monitorHostDto)
	if err != nil {
		return nil, err
	}

	return monitorHostDtos, nil

}

/**
查询 系统信息
*/
func (monitorHostService *MonitorHostLogService) GetMonitorHostLogs(ctx iris.Context) result.ResultDto {
	var (
		err             error
		monitorHostDto  = monitor.MonitorHostLogDto{}
		monitorHostDtos []*monitor.MonitorHostLogDto
	)

	hostId := ctx.URLParam("hostId")

	if hostId == "" {
		return result.Error("未包含主机ID")
	}

	monitorHostDto.HostId = hostId

	monitorHostDtos, err = monitorHostService.monitorHostDao.GetMonitorHostLogs(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDtos)

}

/**
保存 系统信息
*/
func (monitorHostService *MonitorHostLogService) SaveMonitorHostLogs(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostLogDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	monitorHostDto.TenantId = user.TenantId
	monitorHostDto.LogId = seq.Generator()
	monitorHostDto.CpuRate = "0"
	monitorHostDto.DiskRate = "0"
	monitorHostDto.MemRate = "0"

	err = monitorHostService.monitorHostDao.SaveMonitorHostLog(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}

/**
修改 系统信息
*/
func (monitorHostService *MonitorHostLogService) UpdateMonitorHostLogs(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostLogDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostService.monitorHostDao.UpdateMonitorHostLog(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}

/**
删除 系统信息
*/
func (monitorHostService *MonitorHostLogService) DeleteMonitorHostLogs(ctx iris.Context) result.ResultDto {
	var (
		err            error
		monitorHostDto monitor.MonitorHostLogDto
	)

	if err = ctx.ReadJSON(&monitorHostDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorHostService.monitorHostDao.DeleteMonitorHostLog(monitorHostDto)
	if err != nil {
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorHostDto)

}
