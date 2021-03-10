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
	monitorTaskAttrDao dao.MonitorTaskAttrDao

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

}/**
查询 系统信息
*/
func (monitorTaskService *MonitorTaskService) ListTaskTemplate(ctx iris.Context)  (result.ResultDto) {

	var (
		err       error

		monitorTaskTemplateDto = monitor.MonitorTaskTemplateDto{}
		monitorTaskTemplateDtos []*monitor.MonitorTaskTemplateDto

		tmpTemp = make([]*monitor.MonitorTaskTemplateDto,0,1)

	)

	monitorTaskTemplateDtos,err = monitorTaskService.monitorTaskDao.GetTaskTemplate(monitorTaskTemplateDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	//先封装 template
	for _, item := range monitorTaskTemplateDtos{
		if !checkHasTemplate(tmpTemp,*item){
			tmpTemp = append(tmpTemp,item)
		}
	}

	for _,item := range tmpTemp{
		var tmpTempSpec = make([]monitor.MonitorTaskTemplateSpecDto,0,1)
		for _,allItme := range monitorTaskTemplateDtos{
			if item.TemplateId == allItme.TemplateId{
				var tempSpecDto = monitor.MonitorTaskTemplateSpecDto{
					SpecCd: allItme.SpecCd,
					SpecName: allItme.SpecName,
					SpecDesc: allItme.SpecDesc,
					IsShow: allItme.IsShow,
				}

				tmpTempSpec = append(tmpTempSpec,tempSpecDto)
			}

		}

		item.MonitorTaskTemplateSpecDto = tmpTempSpec
	}


	return result.SuccessData(tmpTemp)

}

func checkHasTemplate(i []*monitor.MonitorTaskTemplateDto, dto monitor.MonitorTaskTemplateDto) bool {
	for _,item := range i{
		if dto.TemplateId == item.TemplateId{
			return true
		}
	}

	return false
}


/**
保存 系统信息
*/
func (monitorTaskService *MonitorTaskService) SaveMonitorTasks(ctx iris.Context)  (result.ResultDto){
	var (
		err       error
		monitorTaskDto  monitor.MonitorTaskDto
		attr monitor.MonitorTaskAttrDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)
	monitorTaskDto.TenantId = user.TenantId
	monitorTaskDto.TaskId = seq.Generator()
	monitorTaskDto.State = "001"


	err = monitorTaskService.monitorTaskDao.SaveMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}


	attrs :=monitorTaskDto.Attr


	if len(attrs)<1{
		return result.SuccessData(monitorTaskDto)
	}

	for _,item := range attrs{
		attr = monitor.MonitorTaskAttrDto{
			AttrId: seq.Generator(),
			TaskId: monitorTaskDto.TaskId,
			SpecCd: item.SpecCd,
			Value: item.Value,
		}

		monitorTaskService.monitorTaskAttrDao.SaveMonitorTaskAttr(attr)
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
		attr monitor.MonitorTaskAttrDto
	)

	if err = ctx.ReadJSON(&monitorTaskDto); err != nil {
		return result.Error("解析入参失败")
	}

	err = monitorTaskService.monitorTaskDao.UpdateMonitorTask(monitorTaskDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	attrs :=monitorTaskDto.Attr


	if len(attrs)<1{
		return result.SuccessData(monitorTaskDto)
	}

	for _,item := range attrs{
		attr = monitor.MonitorTaskAttrDto{
			AttrId: item.AttrId,
			TaskId: monitorTaskDto.TaskId,
			SpecCd: item.SpecCd,
			Value: item.Value,
		}

		if attr.AttrId != "" && attr.AttrId != "-1"{
			monitorTaskService.monitorTaskAttrDao.UpdateMonitorTaskAttr(attr)
		}else{
			attr.AttrId=seq.Generator()
			monitorTaskService.monitorTaskAttrDao.SaveMonitorTaskAttr(attr)
		}

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