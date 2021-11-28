package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/businessPackage"
	"gorm.io/gorm"
)

const (
	query_businessPackage_count string = `
	select count(1) total
	from business_package t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if Varsion != '' then
	and t.varsion = #Varsion#
	$endif
	$if CreateUserId != '' then
	and t.create_user_id = #CreateUserId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_businessPackage string = `
				select t.*
				from business_package t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if Varsion != '' then
				and t.varsion = #Varsion#
				$endif
				$if CreateUserId != '' then
				and t.create_user_id = #CreateUserId#
				$endif
				$if Id != '' then
				and t.id = #Id#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_businessPackage string = `
	insert into business_package(id, name, varsion, path, create_user_id,tenant_id)
VALUES(#Id#,#Name#,#Version#,#Path#,#CreateUserId#,#TenantId#)
`

	update_businessPackage string = `
	update business_package set
		$if Name != '' then
		name = #Name#,
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
	delete_businessPackage string = `
	update business_package  set
                          status_cd = '1'
                          where status_cd = '0'

						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type BusinessPackageDao struct {
}

/**
查询用户
*/
func (*BusinessPackageDao) GetBusinessPackageCount(businessPackageDto businessPackage.BusinessPackageDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessPackage_count, objectConvert.Struct2Map(businessPackageDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*BusinessPackageDao) GetBusinessPackages(businessPackageDto businessPackage.BusinessPackageDto) ([]*businessPackage.BusinessPackageDto, error) {
	var businessPackageDtos []*businessPackage.BusinessPackageDto
	sqlTemplate.SelectList(query_businessPackage, objectConvert.Struct2Map(businessPackageDto), func(db *gorm.DB) {
		db.Scan(&businessPackageDtos)
	}, false)

	return businessPackageDtos, nil
}

/**
保存服务sql
*/
func (*BusinessPackageDao) SaveBusinessPackage(businessPackageDto businessPackage.BusinessPackageDto) error {
	return sqlTemplate.Insert(insert_businessPackage, objectConvert.Struct2Map(businessPackageDto), false)
}

/**
修改服务sql
*/
func (*BusinessPackageDao) UpdateBusinessPackage(businessPackageDto businessPackage.BusinessPackageDto) error {
	return sqlTemplate.Update(update_businessPackage, objectConvert.Struct2Map(businessPackageDto), false)
}

/**
删除服务sql
*/
func (*BusinessPackageDao) DeleteBusinessPackage(businessPackageDto businessPackage.BusinessPackageDto) error {
	return sqlTemplate.Delete(delete_businessPackage, objectConvert.Struct2Map(businessPackageDto), false)
}
