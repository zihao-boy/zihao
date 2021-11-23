package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

const (
	query_monitorEvent_count string = `
		select count(1) total
		 from monitor_event t
		where
		t.status_cd = '0'
		and t.tenant_id = #TenantId#
		$if EventId != '' then
		and t.event_id = #EventId#
		$endif
		$if NoticeType != '' then
		and t.notice_type = #NoticeType#
		$endif
		$if State != '' then
		and t.state= #State#
		$endif
		$if EventObjName != '' then
		and t.event_obj_name = #EventObjName#
		$endif
		$if EventObjId != '' then
		and t.event_obj_id = #EventObjId#
		$endif
    	
	`
	query_monitorEvent string = `
				select t.*,td.name notice_type_name from monitor_event t
				left join t_dict td on td.status_cd = t.notice_type and td.table_name = 'monitor_host_group' and td.table_columns = 'notice_type'

				where
				t.status_cd = '0'
				and t.tenant_id = #TenantId#
				$if EventId != '' then
				and t.event_id = #EventId#
				$endif
				$if NoticeType != '' then
				and t.notice_type = #NoticeType#
				$endif
				$if State != '' then
				and t.state= #State#
				$endif
				$if EventObjName != '' then
				and t.event_obj_name = #EventObjName#
				$endif
				$if EventObjId != '' then
				and t.event_obj_id = #EventObjId#
				$endif
				 order by t.create_time desc
				$if Row != 0 then
					limit #Page#,#Row#
				$endif
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
func (*MonitorEventDao) GetMonitorEventCount(monitorEventDto monitor.MonitorEventDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_monitorEvent_count, objectConvert.Struct2Map(monitorEventDto), &pageDto, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*MonitorEventDao) GetMonitorEvents(monitorEventDto monitor.MonitorEventDto) ([]*monitor.MonitorEventDto, error) {
	var monitorEventDtos []*monitor.MonitorEventDto
	sqlTemplate.SelectList(query_monitorEvent, objectConvert.Struct2Map(monitorEventDto), &monitorEventDtos, false)

	return monitorEventDtos, nil
}

/**
保存服务sql
*/
func (*MonitorEventDao) SaveMonitorEvent(monitorEventDto monitor.MonitorEventDto) error {
	return sqlTemplate.Insert(insert_monitorEvent, objectConvert.Struct2Map(monitorEventDto), false)
}

/**
修改服务sql
*/
func (*MonitorEventDao) UpdateMonitorEvent(monitorEventDto monitor.MonitorEventDto) error {
	return sqlTemplate.Update(update_monitorEvent, objectConvert.Struct2Map(monitorEventDto), false)
}

/**
删除服务sql
*/
func (*MonitorEventDao) DeleteMonitorEvent(monitorEventDto monitor.MonitorEventDto) error {
	return sqlTemplate.Delete(delete_monitorEvent, objectConvert.Struct2Map(monitorEventDto), false)
}
