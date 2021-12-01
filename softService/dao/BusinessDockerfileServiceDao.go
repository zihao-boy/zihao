package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/businessDockerfile"
	"gorm.io/gorm"
)

const (
	query_businessDockerfile_count string = `
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
	query_businessDockerfile string = `
				select t.*,uu.username
				from business_package t
				left join u_user uu on t.create_user_id = uu.user_id and uu.status_cd = '0'
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

	insert_businessDockerfile string = `
	insert into business_package(id, name, varsion, path, create_user_id,tenant_id)
VALUES(#Id#,#Name#,#Varsion#,#Path#,#CreateUserId#,#TenantId#)
`

	update_businessDockerfile string = `
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
	delete_businessDockerfile string = `
	update business_package  set
                          status_cd = '1'
                          where status_cd = '0'

						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type BusinessDockerfileDao struct {
}

/**
查询用户
*/
func (*BusinessDockerfileDao) GetBusinessDockerfileCount(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_businessDockerfile_count, objectConvert.Struct2Map(businessDockerfileDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*BusinessDockerfileDao) GetBusinessDockerfiles(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) ([]*businessDockerfile.BusinessDockerfileDto, error) {
	var businessDockerfileDtos []*businessDockerfile.BusinessDockerfileDto
	sqlTemplate.SelectList(query_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), func(db *gorm.DB) {
		db.Scan(&businessDockerfileDtos)
	}, false)

	return businessDockerfileDtos, nil
}

/**
保存服务sql
*/
func (*BusinessDockerfileDao) SaveBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Insert(insert_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}

/**
修改服务sql
*/
func (*BusinessDockerfileDao) UpdateBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Update(update_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}

/**
删除服务sql
*/
func (*BusinessDockerfileDao) DeleteBusinessDockerfile(businessDockerfileDto businessDockerfile.BusinessDockerfileDto) error {
	return sqlTemplate.Delete(delete_businessDockerfile, objectConvert.Struct2Map(businessDockerfileDto), false)
}
