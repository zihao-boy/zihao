package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

const(
	query_monitorEvent_count string = `
		select count(1) total from monitor_host t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if MhgId != '' then
					and t.mhg_id = #MhgId#
					$endif
					$if HostId != '' then
					and t.host_id = #HostId#
					$endif
					$if MhId != '' then
					and t.mh_id = #MhId#
					$endif
    	
	`
	query_monitorEvent string = `
		
				select DATE_FORMAT(t.create_time,'%H') create_time,max(t.cpu_rate) cpu_rate,max(t.mem_rate) mem_rate,max(t.disk_rate) disk_rate
				from monitor_host_log t
				where t.host_id = #HostId#
				  and t.create_time >=(NOW() - interval 24 hour)
				 group by DATE_FORMAT(t.create_time,'%H')
				 order by t.create_time asc
	`

	insert_monitorEvent string = `
	insert into monitor_event(event_id, event_type, event_obj_id, event_obj_name, tenant_id, threshold_value, cur_value, remark,notice_type,state,state_remark)
VALUES (#EventId#, #EventType#, #EventObjId#, #EventObjName#,#TenantId#, #ThresholdValue#, #CurValue#, #Remark#,#NoticeType#,#State#,#StateRemark#)
	`

	update_monitorEvent string = `
	
	`
	delete_monitorEvent string = `
	
	`
)

type MonitorEventDao struct {

}

/**
查询用户
*/
func (*MonitorEventDao) GetMonitorEventCount(monitorEventDto monitor.MonitorEventDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_monitorEvent_count,objectConvert.Struct2Map(monitorEventDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*MonitorEventDao) GetMonitorEvents(monitorEventDto monitor.MonitorEventDto) ([]*monitor.MonitorEventDto,error){
	var monitorEventDtos []*monitor.MonitorEventDto
	sqlTemplate.SelectList(query_monitorEvent,objectConvert.Struct2Map(monitorEventDto), func(db *gorm.DB) {
		db.Scan(&monitorEventDtos)
	},false)

	return monitorEventDtos,nil
}

/**
保存服务sql
*/
func (*MonitorEventDao) SaveMonitorEvent(monitorEventDto monitor.MonitorEventDto) error{
	return sqlTemplate.Insert(insert_monitorEvent,objectConvert.Struct2Map(monitorEventDto),false)
}

/**
修改服务sql
*/
func (*MonitorEventDao) UpdateMonitorEvent(monitorEventDto monitor.MonitorEventDto) error{
	return sqlTemplate.Update(update_monitorEvent,objectConvert.Struct2Map(monitorEventDto),false)
}

/**
删除服务sql
*/
func (*MonitorEventDao) DeleteMonitorEvent(monitorEventDto monitor.MonitorEventDto) error{
	return sqlTemplate.Delete(delete_monitorEvent,objectConvert.Struct2Map(monitorEventDto),false)
}
