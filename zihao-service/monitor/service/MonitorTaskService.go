package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
	"strconv"
)

type MonitorTaskService struct {
	monitorTaskDao dao.MonitorTaskDao

}

/**
查询 系统信息
*/
func (monitorTaskService *MonitorTaskService) GetMonitorTaskAll(monitorTaskDto monitor.MonitorTaskDto)  ([]*monitor.MonitorTaskDto,error) {
	var (
		err       error
		monitorTaskDtos []*monitor.MonitorTaskDto
	)

	monitorTaskDtos,err = monitorTaskService.monitorTaskDao.GetMonitorTasks(monitorTaskDto)
	if(err != nil){
		return nil,err
	}

	return monitorTaskDtos,nil

}

/**
查询 系统信息
*/
func (monitorTaskService *MonitorTaskService) GetMonitorTasks(ctx iris.Context)  (result.ResultDto) {
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		err       error
		page int64
		row int64
		total int64
		monitorTaskDto = monitor.MonitorTaskDto{TenantId: user.TenantId}
		monitorTaskDtos []*monitor.MonitorTaskDto
	)


	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	monitorTaskDto.Row = row * page

	monitorTaskDto.Page = (page -1) * row
	total,err = monitorTaskService.monitorTaskDao.GetMonitorTaskCount(monitorTaskDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}

	monitorTaskDtos,err = monitorTaskService.monitorTaskDao.GetMonitorTasks(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskDtos,total,row)

}


/**
保存 系统信息
*/
func (monitorTaskService *MonitorTaskService) SaveMonitorTasks(ctx iris.Context)  (result.ResultDto){
	var (
		err       error
		monitorTaskDto  monitor.MonitorTaskDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	monitorTaskDto.TenantId = user.TenantId
	monitorTaskDto.TaskId = seq.Generator()


	err = monitorTaskService.monitorTaskDao.SaveMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskDto)

}


/**
修改 系统信息
*/
func (monitorTaskService *MonitorTaskService) UpdateMonitorTasks(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorTaskDto monitor.MonitorTaskDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorTaskService.monitorTaskDao.UpdateMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskDto)

}


/**
删除 系统信息
*/
func (monitorTaskService *MonitorTaskService) DeleteMonitorTasks(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorTaskDto monitor.MonitorTaskDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorTaskService.monitorTaskDao.DeleteMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorTaskDto)

}


/**
停止组
*/
func (monitorTaskService *MonitorTaskService) StartMonitorTask(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorTaskDto monitor.MonitorTaskDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}
	monitorTaskDto.State= "002"

	err = monitorTaskService.monitorTaskDao.UpdateMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}


	return result.SuccessData(monitorTaskDto)

}

/**
停止组
*/
func (monitorTaskService *MonitorTaskService) StopMonitorTasks(ctx iris.Context)  (result.ResultDto) {
	var (
		err       error
		monitorTaskDto monitor.MonitorTaskDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}
	monitorTaskDto.State= "001"

	err = monitorTaskService.monitorTaskDao.UpdateMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}


	return result.SuccessData(monitorTaskDto)

}