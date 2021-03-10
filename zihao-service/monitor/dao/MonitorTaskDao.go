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
		 from task t 
where 1=1
    $if TenantId != '' then
  and t.tenant_id = #TenantId#
  $endif
    $if HostId != '' then
  and t.host_id = #HostId#
   $endif
  and t.status_cd = '0'
 $if TaskName != '' then
    and t.task_name = #TaskName#
$endif
	`
	query_monitorTask string = `
				select t.*,tt.template_name, tt.class_bean,h.name host_name,h.ip from task t
left join task_template tt on t.template_id = tt.template_id and tt.status_cd='0'
left join host h on t.host_id = h.host_id and h.status_cd = '0'
					where 1=1
						$if TenantId != '' then
					  and t.tenant_id = #TenantId#
					  $endif
						$if HostId != '' then
					  and t.host_id = #HostId#
					   $endif
					  and t.status_cd = '0'
					 $if TaskName != '' then
						and t.task_name = #TaskName#
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
			update task t set
			$if TaskName != '' then
			t.task_name = #TaskName#,
			$endif
			$if TaskCron != '' then
			t.task_cron = #TaskCron#,
			$endif
			$if HostId != '' then
			t.host_id = #HostId#,
			$endif
			$if TemplateId != '' then
			t.template_id = #TemplateId#,
			$endif
			$if State != '' then
			t.state = #State#,
			$endif
			t.status_cd = '0'
			where
			 1=1 
			 and t.task_id = #TaskId#
			and t.status_cd = '0'
	`
	delete_monitorTask string = `
	update task t set
			t.status_cd = '1'
			where
			 1=1 
			 and t.task_id = #TaskId#
			and t.status_cd = '0'
	`

	query_monitorTaskTemplate string =`
	select t.*,tt.spec_id, tt.template_id, tt.spec_cd, tt.spec_name, tt.spec_desc, tt.is_show
from task_template t
left join task_template_spec tt on t.template_id = tt.template_id and tt.status_cd = '0'
where t.status_cd = '0'
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
查询用户
*/
func (*MonitorTaskDao) GetTaskTemplate(monitorTaskTemplateDto monitor.MonitorTaskTemplateDto) ([]*monitor.MonitorTaskTemplateDto,error){
	var monitorTaskDtos []*monitor.MonitorTaskTemplateDto
	sqlTemplate.SelectList(query_monitorTaskTemplate,objectConvert.Struct2Map(monitorTaskTemplateDto), func(db *gorm.DB) {
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
