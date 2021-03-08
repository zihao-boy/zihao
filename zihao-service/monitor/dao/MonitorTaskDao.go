package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

const(
	query_monitorTask_count string = `
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
    	
	`
	query_monitorTask string = `
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
				 order by t.create_time desc
				$if Row != 0 then
					limit #Page#,#Row#
				$endif
	`

	insert_monitorTask string = `
		insert into task(task_id, task_name, template_id, task_cron, state, tenant_id, host_id) 
		VALUES (#TaskId#, #TaskName#, #TemplateId#, #TaskCron#, #State#,  #TenantId#, #HostId#)
	`

	update_monitorTask string = `
	
	`
	delete_monitorTask string = `
	
	`
)

type MonitorTaskDao struct {

}

/**
查询用户
*/
func (*MonitorTaskDao) GetMonitorTaskCount(monitorTaskDto monitor.MonitorTaskDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_monitorTask_count,objectConvert.Struct2Map(monitorTaskDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*MonitorTaskDao) GetMonitorTasks(monitorTaskDto monitor.MonitorTaskDto) ([]*monitor.MonitorTaskDto,error){
	var monitorTaskDtos []*monitor.MonitorTaskDto
	sqlTemplate.SelectList(query_monitorTask,objectConvert.Struct2Map(monitorTaskDto), func(db *gorm.DB) {
		db.Scan(&monitorTaskDtos)
	},false)

	return monitorTaskDtos,nil
}

/**
保存服务sql
*/
func (*MonitorTaskDao) SaveMonitorTask(monitorTaskDto monitor.MonitorTaskDto) error{
	return sqlTemplate.Insert(insert_monitorTask,objectConvert.Struct2Map(monitorTaskDto),false)
}

/**
修改服务sql
*/
func (*MonitorTaskDao) UpdateMonitorTask(monitorTaskDto monitor.MonitorTaskDto) error{
	return sqlTemplate.Update(update_monitorTask,objectConvert.Struct2Map(monitorTaskDto),false)
}

/**
删除服务sql
*/
func (*MonitorTaskDao) DeleteMonitorTask(monitorTaskDto monitor.MonitorTaskDto) error{
	return sqlTemplate.Delete(delete_monitorTask,objectConvert.Struct2Map(monitorTaskDto),false)
}
