package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
	"gorm.io/gorm"
)

const (
	query_monitorHost_count string = `
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
	query_monitorHost string = `
		
				select t.*,h.name,h.ip,h.passwd,h.username from monitor_host t
				inner join host h on t.host_id = h.host_id and h.status_cd = '0'
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
					order by t.create_time desc
					$if Page != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_monitorHost string = `
insert into monitor_host(mh_id, mhg_id, host_id, tenant_id, cpu_rate, mem_rate, disk_rate, free_mem, free_disk, 
                         cpu_threshold, mem_threshold, disk_threshold, mon_disk, mon_date) 
                         values(#MhId#, #MhgId#, #HostId#, #TenantId#, #CpuRate#, #MemRate#, #DiskRate#, 
                         #FreeMem#, #FreeDisk#,#CpuThreshold#, #MemThreshold#, #DiskThreshold#, #MonDisk#, #MonDate#) 
`

	update_monitorHost string = `
	update monitor_host t set
    $if CpuRate != '' then
                          t.cpu_rate = #CpuRate#,
                          $endif
					$if CpuThreshold != '' then
                          t.cpu_threshold=#CpuThreshold#,
                          $endif
					$if DiskRate != '' then
                          t.disk_rate=#DiskRate#,
                          $endif
					$if DiskThreshold != '' then
                          t.disk_threshold=#DiskThreshold#,
                          $endif
					$if FreeDisk != '' then
                          t.free_disk = #FreeDisk#,
                          $endif
					$if FreeMem != '' then
                          t.free_mem=#FreeMem#,
                          $endif
					$if MemRate != '' then
                          t.mem_rate = #MemRate#,
                          $endif
					$if MemThreshold != '' then
                          t.mem_threshold=#MemThreshold#,
                          $endif
					$if MonDate != '' then
                          t.mon_date=#MonDate#,
                          $endif
					$if MonDisk != '' then
                          t.mon_disk=#MonDisk#,
                          $endif
                          t.status_cd = '0'
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
	delete_monitorHost string = `
	update monitor_host t set
                          t.status_cd = '1'
                          where t.status_cd = '0'

					$if MhId != '' then
					and t.mh_id = #MhId#
					$endif
	`
)

type MonitorHostDao struct {
}

/**
查询用户
*/
func (*MonitorHostDao) GetMonitorHostCount(monitorHostDto monitor.MonitorHostDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_monitorHost_count, objectConvert.Struct2Map(monitorHostDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*MonitorHostDao) GetMonitorHosts(monitorHostDto monitor.MonitorHostDto) ([]*monitor.MonitorHostDto, error) {
	var monitorHostDtos []*monitor.MonitorHostDto
	sqlTemplate.SelectList(query_monitorHost, objectConvert.Struct2Map(monitorHostDto), func(db *gorm.DB) {
		db.Scan(&monitorHostDtos)
	}, false)

	return monitorHostDtos, nil
}

/**
保存服务sql
*/
func (*MonitorHostDao) SaveMonitorHost(monitorHostDto monitor.MonitorHostDto) error {
	return sqlTemplate.Insert(insert_monitorHost, objectConvert.Struct2Map(monitorHostDto), false)
}

/**
修改服务sql
*/
func (*MonitorHostDao) UpdateMonitorHost(monitorHostDto monitor.MonitorHostDto) error {
	return sqlTemplate.Update(update_monitorHost, objectConvert.Struct2Map(monitorHostDto), false)
}

/**
删除服务sql
*/
func (*MonitorHostDao) DeleteMonitorHost(monitorHostDto monitor.MonitorHostDto) error {
	return sqlTemplate.Delete(delete_monitorHost, objectConvert.Struct2Map(monitorHostDto), false)
}
