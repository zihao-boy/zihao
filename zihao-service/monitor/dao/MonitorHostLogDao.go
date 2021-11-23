package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"gorm.io/gorm"
)

const (
	query_monitorHostLog_count string = `
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
	query_monitorHostLog string = `
		
				select DATE_FORMAT(t.create_time,'%H') create_time,max(t.cpu_rate) cpu_rate,max(t.mem_rate) mem_rate,max(t.disk_rate) disk_rate
				from monitor_host_log t
				where t.host_id = #HostId#
				  and t.create_time >=(NOW() - interval 24 hour)
				 group by DATE_FORMAT(t.create_time,'%H')
				 order by t.create_time asc
	`

	insert_monitorHostLog string = `
		insert into monitor_host_log(log_id, host_id, tenant_id, cpu_rate, mem_rate, disk_rate)
		values (#LogId#, #HostId#, #TenantId#, #CpuRate#, #MemRate#, #DiskRate#)
	`

	update_monitorHostLog string = `
	
	`
	delete_monitorHostLog string = `
	
	`
)

type MonitorHostLogDao struct {
}

/**
查询用户
*/
func (*MonitorHostLogDao) GetMonitorHostLogCount(monitorHostLogDto monitor.MonitorHostLogDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_monitorHostLog_count, objectConvert.Struct2Map(monitorHostLogDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*MonitorHostLogDao) GetMonitorHostLogs(monitorHostLogDto monitor.MonitorHostLogDto) ([]*monitor.MonitorHostLogDto, error) {
	var monitorHostLogDtos []*monitor.MonitorHostLogDto
	sqlTemplate.SelectList(query_monitorHostLog, objectConvert.Struct2Map(monitorHostLogDto), func(db *gorm.DB) {
		db.Scan(&monitorHostLogDtos)
	}, false)

	return monitorHostLogDtos, nil
}

/**
保存服务sql
*/
func (*MonitorHostLogDao) SaveMonitorHostLog(monitorHostLogDto monitor.MonitorHostLogDto) error {
	return sqlTemplate.Insert(insert_monitorHostLog, objectConvert.Struct2Map(monitorHostLogDto), false)
}

/**
修改服务sql
*/
func (*MonitorHostLogDao) UpdateMonitorHostLog(monitorHostLogDto monitor.MonitorHostLogDto) error {
	return sqlTemplate.Update(update_monitorHostLog, objectConvert.Struct2Map(monitorHostLogDto), false)
}

/**
删除服务sql
*/
func (*MonitorHostLogDao) DeleteMonitorHostLog(monitorHostLogDto monitor.MonitorHostLogDto) error {
	return sqlTemplate.Delete(delete_monitorHostLog, objectConvert.Struct2Map(monitorHostLogDto), false)
}
