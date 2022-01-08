package dbLinkDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/dbLink"
	"gorm.io/gorm"
)

const (
	query_dbLink_count string = `
	select count(1) total
	from db_link t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if CreateUserId != '' then
	and t.create_user_id = #CreateUserId#
	$endif
	$if Id != '' then
	and t.id = #Id#
	$endif
    	
	`
	query_dbLink string = `
				select t.*
				from db_link t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
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

	insert_dbLink string = `
	insert into db_link(id, name, ip, port,username,password,db_name, create_user_id,tenant_id)
VALUES(#Id#,#Name#,#Ip#,#Port#,#Username#,#Password#,#DbName#,#CreateUserId#,#TenantId#)
`

	update_dbLink string = `
	update db_link set
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
		$if Id != '' then
		and id = #Id#
		$endif
	`
	delete_dbLink string = `
	update db_link  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type DbLinkDao struct {
}

/**
查询用户
*/
func (*DbLinkDao) GetDbLinkCount(dbLinkDto dbLink.DbLinkDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_dbLink_count, objectConvert.Struct2Map(dbLinkDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*DbLinkDao) GetDbLinks(dbLinkDto dbLink.DbLinkDto) ([]*dbLink.DbLinkDto, error) {
	var dbLinkDtos []*dbLink.DbLinkDto
	sqlTemplate.SelectList(query_dbLink, objectConvert.Struct2Map(dbLinkDto), func(db *gorm.DB) {
		db.Scan(&dbLinkDtos)
	}, false)

	return dbLinkDtos, nil
}

/**
保存服务sql
*/
func (*DbLinkDao) SaveDbLink(dbLinkDto dbLink.DbLinkDto) error {
	return sqlTemplate.Insert(insert_dbLink, objectConvert.Struct2Map(dbLinkDto), false)
}

/**
修改服务sql
*/
func (*DbLinkDao) UpdateDbLink(dbLinkDto dbLink.DbLinkDto) error {
	return sqlTemplate.Update(update_dbLink, objectConvert.Struct2Map(dbLinkDto), false)
}

/**
删除服务sql
*/
func (*DbLinkDao) DeleteDbLink(dbLinkDto dbLink.DbLinkDto) error {
	return sqlTemplate.Delete(delete_dbLink, objectConvert.Struct2Map(dbLinkDto), false)
}
