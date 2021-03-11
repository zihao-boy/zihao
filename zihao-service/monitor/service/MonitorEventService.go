package service

import (
	"github.com/kataras/iris/v12"
	"github.com/zihao-boy/zihao/zihao-service/common/constants"
	"github.com/zihao-boy/zihao/zihao-service/common/httpReq"
	"github.com/zihao-boy/zihao/zihao-service/common/seq"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/result"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/tenant"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/user"
	"github.com/zihao-boy/zihao/zihao-service/monitor/dao"
	"github.com/zihao-boy/zihao/zihao-service/user/service"
	"strconv"
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
	var user *user.UserDto = ctx.Values().Get(constants.UINFO).(*user.UserDto)

	var (
		err       error
		page int64
		row int64
		total int64
		monitorEventDto = monitor.MonitorEventDto{TenantId: user.TenantId}
		monitorEventDtos []*monitor.MonitorEventDto
	)


	page,err =  strconv.ParseInt(ctx.URLParam("page"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	row,err =  strconv.ParseInt(ctx.URLParam("row"),10,64)

	if err != nil{
		return result.Error(err.Error())
	}

	monitorEventDto.Row = row * page

	monitorEventDto.Page = (page -1) * row

	monitorEventDto.EventObjName = ctx.URLParam("eventObjName")

	total,err = monitorEventService.monitorEventDao.GetMonitorEventCount(monitorEventDto)

	if err != nil{
		return result.Error(err.Error())
	}

	if total < 1{
		return result.Success()
	}

	monitorEventDtos,err = monitorEventService.monitorEventDao.GetMonitorEvents(monitorEventDto)
	if(err != nil){
		return result.Error(err.Error())
	}

	return result.SuccessData(monitorEventDtos,total,row)

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

	//钉钉告警
	if eventDto.NoticeType == "2002"{
		sendToDingDing(eventDto)
	}else if eventDto.NoticeType == "3003"{
		sendToCompanyWechat(eventDto)
	}

	return nil

}


func sendToCompanyWechat(eventDto monitor.MonitorEventDto) (string,error){

	//eventDto.TenantId

	var tenantSettingService service.TenantSettingService
	var tenantSettingDto = tenant.TenantSettingDto{
		TenantId: eventDto.TenantId,
		SpecCd: "300302",
	}
	tenantSettingDtos,err := tenantSettingService.GetTenantSettingAll(tenantSettingDto)

	if err != nil || len(tenantSettingDtos)<1{
		return "",err
	}
	//根据告警类型告警相应平台
	var url string = tenantSettingDtos[0].Value
	// 1、构建需要的参数
	context := map[string]string{
		"content": "[梓豪平台告警]"+eventDto.Remark,
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"text": context,
	}
	resp, err := httpReq.SendRequest(url,data,nil,"POST")
	return string(resp),err
}

func sendToDingDing(eventDto monitor.MonitorEventDto) (string,error){

	//eventDto.TenantId

	var tenantSettingService service.TenantSettingService
	var tenantSettingDto = tenant.TenantSettingDto{
		TenantId: eventDto.TenantId,
		SpecCd: "300301",
	}
	tenantSettingDtos,err := tenantSettingService.GetTenantSettingAll(tenantSettingDto)

	if err != nil || len(tenantSettingDtos)<1{
		return "",err
	}
	//根据告警类型告警相应平台
	var url string = tenantSettingDtos[0].Value
	// 1、构建需要的参数
	context := map[string]string{
		"content": "[梓豪平台告警]"+eventDto.Remark,
	}
	data := map[string]interface{}{
		"msgtype": "text",
		"text": context,
	}
	resp, err := httpReq.SendRequest(url,data,nil,"POST")
	return string(resp),err
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