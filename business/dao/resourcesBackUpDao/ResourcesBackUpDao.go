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
				select t.*
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
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_resourcesBackUp string = `
	insert into resources_backup(id, name, exec_time, type_cd,src_id,src_object,target_id,tenant_id)
VALUES(#Id#,#Name#,#ExecTime#,#TypeCd#,#SrcId#,#SrcObject#,#TargetId#,#TenantId#)
`

	update_resourcesBackUp string = `
	update resources_backup set
		$if Name != '' then
		name = #Name#,
		$endif
		$if OssType != '' then
		oss_type = #OssType#,
		$endif
		$if Bucket != '' then
		bucket = #Bucket#,
		$endif
		$if AccessKeySecret != '' then
		access_key_secret = #AccessKeySecret#,
		$endif
		$if AccessKeyId != '' then
		access_key_id = #AccessKeyId#,
		$endif
		$if Path != '' then
		path = #Path#,
		$endif
		$if Endpoint != '' then
		endpoint = #Endpoint#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if OssId != '' then
		and oss_id = #OssId#
		$endif
	`
	delete_resourcesBackUp string = `
	update resources_oss  set
                          status_cd = '1'
                          where status_cd = '0'
		$if OssId != '' then
		and oss_id = #OssId#
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
