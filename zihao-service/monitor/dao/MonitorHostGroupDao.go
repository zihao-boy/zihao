package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"gorm.io/gorm"
)

const (
	query_monitorHostGroup_count string = `
		select count(1) total  from monitor_host_group t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if State != '' then
				and t.state = #State#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
					$if MhgId != '' then
				and t.mhg_id = #MhgId#
					$endif
					$if MonCron != '' then
				and t.mon_cron = #MonCron#
					$endif
					$if NoticeType != '' then
				and t.notice_type = #NoticeType#
					$endif
    	
	`
	query_monitorHostGroup string = `
		select t.*,td.name state_name, td1.name notice_type_name from monitor_host_group t
left join t_dict td on t.state = td.status_cd and td.table_name = 'monitor_host_group' and td.table_columns = 'state'
left join t_dict td1 on t.notice_type = td1.status_cd and td1.table_name = 'monitor_host_group' and td1.table_columns = 'notice_type'
where  t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if State != '' then
				and t.state = #State#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
					$if MhgId != '' then
				and t.mhg_id = #MhgId#
					$endif
					$if MonCron != '' then
				and t.mon_cron = #MonCron#
					$endif
					$if NoticeType != '' then
				and t.notice_type = #NoticeType#
					$endif
				order by t.create_time desc
				$if Page != 0 then
					limit #Page#,#Row#
				$endif
	`

	insert_monitorHostGroup string = `
insert into monitor_host_group(mhg_id, name, mon_cron, mon_date, notice_type, remark, tenant_id) 
VALUES (#MhgId#, #Name#, #MonCron#, #MonDate#, #NoticeType#, #Remark#, #TenantId#)
`

	update_monitorHostGroup string = `
	update monitor_host_group t set
					$if State != '' then
					 t.state = #State#,
					$endif
					$if Name != '' then
					 t.name = #Name#,
					$endif
					$if MonCron != '' then
					 t.mon_cron = #MonCron#,
					$endif
					$if NoticeType != '' then
					t.notice_type = #NoticeType#,
					$endif
					t.status_cd = '0'
					where  t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if MhgId != '' then
					and t.mhg_id = #MhgId#
					$endif
	`
	delete_monitorHostGroup string = `
	update monitor_host_group t set
					t.status_cd = '1'
					where  t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if MhgId != '' then
					and t.mhg_id = #MhgId#
					$endif
	`
)

type MonitorHostGroupDao struct {
}

/**
查询用户
*/
func (*MonitorHostGroupDao) GetMonitorHostGroupCount(monitorHostGroupDto monitor.MonitorHostGroupDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_monitorHostGroup_count, objectConvert.Struct2Map(monitorHostGroupDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*MonitorHostGroupDao) GetMonitorHostGroups(monitorHostGroupDto monitor.MonitorHostGroupDto) ([]*monitor.MonitorHostGroupDto, error) {
	var monitorHostGroupDtos []*monitor.MonitorHostGroupDto
	sqlTemplate.SelectList(query_monitorHostGroup, objectConvert.Struct2Map(monitorHostGroupDto), func(db *gorm.DB) {
		db.Scan(&monitorHostGroupDtos)
	}, false)

	return monitorHostGroupDtos, nil
}

/**
保存服务sql
*/
func (*MonitorHostGroupDao) SaveMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error {
	return sqlTemplate.Insert(insert_monitorHostGroup, objectConvert.Struct2Map(monitorHostGroupDto), false)
}

/**
修改服务sql
*/
func (*MonitorHostGroupDao) UpdateMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error {
	return sqlTemplate.Update(update_monitorHostGroup, objectConvert.Struct2Map(monitorHostGroupDto), false)
}

/**
删除服务sql
*/
func (*MonitorHostGroupDao) DeleteMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error {
	return sqlTemplate.Delete(delete_monitorHostGroup, objectConvert.Struct2Map(monitorHostGroupDto), false)
}
