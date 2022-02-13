package resourcesOssDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"gorm.io/gorm"
)

const (
	query_resourcesOss_count string = `
	select count(1) total
	from resources_oss t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if OssId != '' then
	and t.oss_id = #OssId#
	$endif
    	
	`
	query_resourcesOss string = `
				select t.*
				from resources_oss t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if OssId != '' then
				and t.oss_id = #OssId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_resourcesOss string = `
	insert into resources_oss(oss_id, name, oss_type, bucket,access_key_secret,access_key_id,path,tenant_id,endpoint)
VALUES(#OssId#,#Name#,#OssType#,#Bucket#,#AccessKeySecret#,#AccessKeyId#,#Path#,#TenantId#,#Endpoint#)
`

	update_resourcesOss string = `
	update resources_oss set
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
	delete_resourcesOss string = `
	update resources_oss  set
                          status_cd = '1'
                          where status_cd = '0'
		$if OssId != '' then
		and oss_id = #OssId#
		$endif
	`
)

type ResourcesOssDao struct {
}

/**
查询用户
*/
func (*ResourcesOssDao) GetResourcesOssCount(resourcesOssDto resources.ResourcesOssDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_resourcesOss_count, objectConvert.Struct2Map(resourcesOssDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*ResourcesOssDao) GetResourcesOsss(resourcesOssDto resources.ResourcesOssDto) ([]*resources.ResourcesOssDto, error) {
	var resourcesOssDtos []*resources.ResourcesOssDto
	sqlTemplate.SelectList(query_resourcesOss, objectConvert.Struct2Map(resourcesOssDto), func(db *gorm.DB) {
		db.Scan(&resourcesOssDtos)
	}, false)

	return resourcesOssDtos, nil
}

/**
保存服务sql
*/
func (*ResourcesOssDao) SaveResourcesOss(resourcesOssDto resources.ResourcesOssDto) error {
	return sqlTemplate.Insert(insert_resourcesOss, objectConvert.Struct2Map(resourcesOssDto), false)
}

/**
修改服务sql
*/
func (*ResourcesOssDao) UpdateResourcesOss(resourcesOssDto resources.ResourcesOssDto) error {
	return sqlTemplate.Update(update_resourcesOss, objectConvert.Struct2Map(resourcesOssDto), false)
}

/**
删除服务sql
*/
func (*ResourcesOssDao) DeleteResourcesOss(resourcesOssDto resources.ResourcesOssDto) error {
	return sqlTemplate.Delete(delete_resourcesOss, objectConvert.Struct2Map(resourcesOssDto), false)
}
