package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/monitor"
)

const(
	query_monitorHostGroup_count string = `
		select count(1) total from monitorHostGroup t
					where t.status_cd = '0'
					$if MonitorHostGroupId != '' then
					and t.monitorHostGroup_id = #MonitorHostGroupId#
					$endif
					$if MonitorHostGroupName != '' then
					and t.monitorHostGroup_name = #MonitorHostGroupName#
					$endif
					$if MonitorHostGroupType != '' then
					and t.monitorHostGroup_type = #MonitorHostGroupType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
    	
	`
	query_monitorHostGroup string = `
		select t.*,uu.username from monitorHostGroup t
				left join u_user uu on t.monitorHostGroup_id = uu.monitorHostGroup_id and uu.user_role = '1001'
					where t.status_cd = '0'
					$if MonitorHostGroupId != '' then
					and t.monitorHostGroup_id = #MonitorHostGroupId#
					$endif
					$if MonitorHostGroupName != '' then
					and t.monitorHostGroup_name = #MonitorHostGroupName#
					$endif
					$if MonitorHostGroupType != '' then
					and t.monitorHostGroup_type = #MonitorHostGroupType#
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

	insert_monitorHostGroup string = `
insert into monitorHostGroup(monitorHostGroup_id, monitorHostGroup_name, address, person_name, phone, remark) 
VALUES(#MonitorHostGroupId#, #MonitorHostGroupName#, #Address#, #PersonName#, #Phone#, #Remark#) 
`

	update_monitorHostGroup string = `
	update monitorHostGroup t set
			$if MonitorHostGroupName != '' then
			 t.monitorHostGroup_name = #MonitorHostGroupName#,
			$endif
			$if MonitorHostGroupType != '' then
			 t.monitorHostGroup_type = #MonitorHostGroupType#,
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
			$if MonitorHostGroupId != '' then
			and t.monitorHostGroup_id = #MonitorHostGroupId#
			$endif
	`
	delete_monitorHostGroup string = `
	update monitorHostGroup t set
			t.status_cd = '1'
			where t.status_cd = '0'
			and t.monitorHostGroup_id = #MonitorHostGroupId#
	`
)

type MonitorHostGroupDao struct {

}

/**
查询用户
*/
func (*MonitorHostGroupDao) GetMonitorHostGroupCount(monitorHostGroupDto monitor.MonitorHostGroupDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_monitorHostGroup_count,objectConvert.Struct2Map(monitorHostGroupDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*MonitorHostGroupDao) GetMonitorHostGroups(monitorHostGroupDto monitor.MonitorHostGroupDto) ([]*monitor.MonitorHostGroupDto,error){
	var monitorHostGroupDtos []*monitor.MonitorHostGroupDto
	sqlTemplate.SelectList(query_monitorHostGroup,objectConvert.Struct2Map(monitorHostGroupDto), func(db *gorm.DB) {
		db.Scan(&monitorHostGroupDtos)
	},false)

	return monitorHostGroupDtos,nil
}

/**
保存服务sql
*/
func (*MonitorHostGroupDao) SaveMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error{
	return sqlTemplate.Insert(insert_monitorHostGroup,objectConvert.Struct2Map(monitorHostGroupDto),false)
}

/**
修改服务sql
*/
func (*MonitorHostGroupDao) UpdateMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error{
	return sqlTemplate.Update(update_monitorHostGroup,objectConvert.Struct2Map(monitorHostGroupDto),false)
}

/**
删除服务sql
*/
func (*MonitorHostGroupDao) DeleteMonitorHostGroup(monitorHostGroupDto monitor.MonitorHostGroupDto) error{
	return sqlTemplate.Delete(delete_monitorHostGroup,objectConvert.Struct2Map(monitorHostGroupDto),false)
}
