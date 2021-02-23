package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

const(
	query_monitorHost_count string = `
		select count(1) total from monitorHost t
					where t.status_cd = '0'
					$if MonitorHostId != '' then
					and t.monitorHost_id = #MonitorHostId#
					$endif
					$if MonitorHostName != '' then
					and t.monitorHost_name = #MonitorHostName#
					$endif
					$if MonitorHostType != '' then
					and t.monitorHost_type = #MonitorHostType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
    	
	`
	query_monitorHost string = `
		select t.*,uu.username from monitorHost t
				left join u_user uu on t.monitorHost_id = uu.monitorHost_id and uu.user_role = '1001'
					where t.status_cd = '0'
					$if MonitorHostId != '' then
					and t.monitorHost_id = #MonitorHostId#
					$endif
					$if MonitorHostName != '' then
					and t.monitorHost_name = #MonitorHostName#
					$endif
					$if MonitorHostType != '' then
					and t.monitorHost_type = #MonitorHostType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Page != -1 then
						limit #Page#,#Row#
					$endif
	`

	insert_monitorHost string = `
insert into monitorHost(monitorHost_id, monitorHost_name, address, person_name, phone, remark) 
VALUES(#MonitorHostId#, #MonitorHostName#, #Address#, #PersonName#, #Phone#, #Remark#) 
`

	update_monitorHost string = `
	update monitorHost t set
			$if MonitorHostName != '' then
			 t.monitorHost_name = #MonitorHostName#,
			$endif
			$if MonitorHostType != '' then
			 t.monitorHost_type = #MonitorHostType#,
			$endif
			$if Phone != '' then
			 t.phone = #Phone#,
			$endif
			$if State != '' then
			 t.state = #State#,
			$endif
			$if Address != '' then
			 t.address = #Address#,
			$endif
			$if PersonName != '' then
			 t.person_name = #PersonName#,
			$endif
			$if Remark != '' then
			 t.remark = #Remark#,
			$endif
			t.status_cd = '0'
			where t.status_cd = '0'
			$if MonitorHostId != '' then
			and t.monitorHost_id = #MonitorHostId#
			$endif
	`
	delete_monitorHost string = `
	update monitorHost t set
			t.status_cd = '1'
			where t.status_cd = '0'
			and t.monitorHost_id = #MonitorHostId#
	`
)

type MonitorHostDao struct {

}

/**
查询用户
*/
func (*MonitorHostDao) GetMonitorHostCount(monitorHostDto monitor.MonitorHostDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_monitorHost_count,objectConvert.Struct2Map(monitorHostDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*MonitorHostDao) GetMonitorHosts(monitorHostDto monitor.MonitorHostDto) ([]*monitor.MonitorHostDto,error){
	var monitorHostDtos []*monitor.MonitorHostDto
	sqlTemplate.SelectList(query_monitorHost,objectConvert.Struct2Map(monitorHostDto), func(db *gorm.DB) {
		db.Scan(&monitorHostDtos)
	},false)

	return monitorHostDtos,nil
}

/**
保存服务sql
*/
func (*MonitorHostDao) SaveMonitorHost(monitorHostDto monitor.MonitorHostDto) error{
	return sqlTemplate.Insert(insert_monitorHost,objectConvert.Struct2Map(monitorHostDto),false)
}

/**
修改服务sql
*/
func (*MonitorHostDao) UpdateMonitorHost(monitorHostDto monitor.MonitorHostDto) error{
	return sqlTemplate.Update(update_monitorHost,objectConvert.Struct2Map(monitorHostDto),false)
}

/**
删除服务sql
*/
func (*MonitorHostDao) DeleteMonitorHost(monitorHostDto monitor.MonitorHostDto) error{
	return sqlTemplate.Delete(delete_monitorHost,objectConvert.Struct2Map(monitorHostDto),false)
}
