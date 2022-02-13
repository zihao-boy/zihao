package resourcesBackUpDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"gorm.io/gorm"
)

const (
	query_resourcesBackUp_count string = `
	select count(1) total
	from resources_backup t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if TypeCd != '' then
	and t.type_cd = #TypeCd#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_resourcesBackUp string = `
select t.*,h.name src_host_name,rd.name src_db_name,rf.name target_ftp_name,ro.name target_oss_name,rd1.name target_db_name
from resources_backup t
left join host h on t.src_id = h.host_id and h.status_cd = '0'
left join resources_db rd on t.src_id = rd.db_id and rd.status_cd = '0'
left join resources_ftp rf on t.target_id = rf.ftp_id and rf.status_cd = '0'
left join resources_oss ro on t.target_id = ro.oss_id and ro.status_cd = '0'
left join resources_db rd1 on t.target_id = rd1.db_id and rd1.status_cd = '0'
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if TypeCd != '' then
				and t.type_cd = #TypeCd#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_resourcesBackUp string = `
	insert into resources_backup(id, name, exec_time, type_cd,src_id,src_object,target_id,tenant_id,state,back_time,target_type_cd)
VALUES(#Id#,#Name#,#ExecTime#,#TypeCd#,#SrcId#,#SrcObject#,#TargetId#,#TenantId#,#State#,#BackTime#,#TargetTypeCd#)
`

	update_resourcesBackUp string = `
	update resources_backup set
		$if Name != '' then
		name = #Name#,
		$endif
		$if TypeCd != '' then
		type_cd = #TypeCd#,
		$endif
		$if ExecTime != '' then
		exec_time = #ExecTime#,
		$endif
		$if SrcId != '' then
		src_id = #SrcId#,
		$endif
		$if SrcObject != '' then
		src_object = #SrcObject#,
		$endif
		$if TargetId != '' then
		target_id = #TargetId#,
		$endif
		$if State != '' then
		state = #State#,
		$endif
		$if BackTime != '' then
		back_time = #BackTime#,
		$endif
		$if TargetTypeCd != '' then
		target_type_cd = #TargetTypeCd#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_resourcesBackUp string = `
	update resources_backup  set
                          status_cd = '1'
                          where status_cd = '0'
		$if Id != '' then
		and id = #Id#
		$endif
	`
)

type ResourcesBackUpDao struct {
}

/**
查询用户
*/
func (*ResourcesBackUpDao) GetResourcesBackUpCount(resourcesBackUpDto resources.ResourcesBackUpDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_resourcesBackUp_count, objectConvert.Struct2Map(resourcesBackUpDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*ResourcesBackUpDao) GetResourcesBackUps(resourcesBackUpDto resources.ResourcesBackUpDto) ([]*resources.ResourcesBackUpDto, error) {
	var resourcesBackUpDtos []*resources.ResourcesBackUpDto
	sqlTemplate.SelectList(query_resourcesBackUp, objectConvert.Struct2Map(resourcesBackUpDto), func(db *gorm.DB) {
		db.Scan(&resourcesBackUpDtos)
	}, false)

	return resourcesBackUpDtos, nil
}

/**
保存服务sql
*/
func (*ResourcesBackUpDao) SaveResourcesBackUp(resourcesBackUpDto resources.ResourcesBackUpDto) error {
	return sqlTemplate.Insert(insert_resourcesBackUp, objectConvert.Struct2Map(resourcesBackUpDto), false)
}

/**
修改服务sql
*/
func (*ResourcesBackUpDao) UpdateResourcesBackUp(resourcesBackUpDto resources.ResourcesBackUpDto) error {
	return sqlTemplate.Update(update_resourcesBackUp, objectConvert.Struct2Map(resourcesBackUpDto), false)
}

/**
删除服务sql
*/
func (*ResourcesBackUpDao) DeleteResourcesBackUp(resourcesBackUpDto resources.ResourcesBackUpDto) error {
	return sqlTemplate.Delete(delete_resourcesBackUp, objectConvert.Struct2Map(resourcesBackUpDto), false)
}
