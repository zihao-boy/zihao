package resourcesDbDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"gorm.io/gorm"
)

const (
	query_resourcesDb_count string = `
	select count(1) total
	from resources_db t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if DbId != '' then
	and t.db_id = #DbId#
	$endif
    	
	`
	query_resourcesDb string = `
				select t.*
				from resources_db t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if DbId != '' then
				and t.db_id = #DbId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_resourcesDb string = `
	insert into resources_db(db_id, name, ip, port,username,password,db_name,tenant_id)
VALUES(#FtpId#,#Name#,#Ip#,#Port#,#Username#,#Password#,#DbName#,#TenantId#)
`

	update_resourcesDb string = `
	update resources_db set
		$if Name != '' then
		name = #Name#,
		$endif
		$if Ip != '' then
		ip = #Ip#,
		$endif
		$if Port != '' then
		port = #Port#,
		$endif
		$if Username != '' then
		username = #Username#,
		$endif
		$if Password != '' then
		password = #Password#,
		$endif
		$if DbName != '' then
		db_name = #DbName#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if DbId != '' then
		and t.db_id = #DbId#
		$endif
	`
	delete_resourcesDb string = `
	update resources_db  set
                          status_cd = '1'
                          where status_cd = '0'
	$if DbId != '' then
	and t.db_id = #DbId#
	$endif
	`
)

type ResourcesDbDao struct {
}

/**
查询用户
*/
func (*ResourcesDbDao) GetResourcesDbCount(resourcesDbDto resources.ResourcesDbDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_resourcesDb_count, objectConvert.Struct2Map(resourcesDbDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*ResourcesDbDao) GetResourcesDbs(resourcesDbDto resources.ResourcesDbDto) ([]*resources.ResourcesDbDto, error) {
	var resourcesDbDtos []*resources.ResourcesDbDto
	sqlTemplate.SelectList(query_resourcesDb, objectConvert.Struct2Map(resourcesDbDto), func(db *gorm.DB) {
		db.Scan(&resourcesDbDtos)
	}, false)

	return resourcesDbDtos, nil
}

/**
保存服务sql
*/
func (*ResourcesDbDao) SaveResourcesDb(resourcesDbDto resources.ResourcesDbDto) error {
	return sqlTemplate.Insert(insert_resourcesDb, objectConvert.Struct2Map(resourcesDbDto), false)
}

/**
修改服务sql
*/
func (*ResourcesDbDao) UpdateResourcesDb(resourcesDbDto resources.ResourcesDbDto) error {
	return sqlTemplate.Update(update_resourcesDb, objectConvert.Struct2Map(resourcesDbDto), false)
}

/**
删除服务sql
*/
func (*ResourcesDbDao) DeleteResourcesDb(resourcesDbDto resources.ResourcesDbDto) error {
	return sqlTemplate.Delete(delete_resourcesDb, objectConvert.Struct2Map(resourcesDbDto), false)
}
