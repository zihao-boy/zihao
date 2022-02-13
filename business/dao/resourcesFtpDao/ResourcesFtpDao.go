package resourcesFtpDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"gorm.io/gorm"
)

const (
	query_resourcesFtp_count string = `
	select count(1) total
	from resources_ftp t
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if Name != '' then
	and t.name = #Name#
	$endif
	$if FtpId != '' then
	and t.ftp_id = #FtpId#
	$endif
    	
	`
	query_resourcesFtp string = `
				select t.*
				from resources_ftp t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if Name != '' then
				and t.name = #Name#
				$endif
				$if FtpId != '' then
				and t.ftp_id = #FtpId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_resourcesFtp string = `
	insert into resources_ftp(ftp_id, name, ip, port,username,passwd,path,tenant_id)
VALUES(#FtpId#,#Name#,#Ip#,#Port#,#Username#,#Passwd#,#Path#,#TenantId#)
`

	update_resourcesFtp string = `
	update resources_ftp set
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
		$if Passwd != '' then
		passwd = #Passwd#,
		$endif
		$if Path != '' then
		path = #Path#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if FtpId != '' then
		and t.ftp_id = #FtpId#
		$endif
	`
	delete_resourcesFtp string = `
	update resources_ftp  set
                          status_cd = '1'
                          where status_cd = '0'
		$if FtpId != '' then
		and t.ftp_id = #FtpId#
		$endif
	`
)

type ResourcesFtpDao struct {
}

/**
查询用户
*/
func (*ResourcesFtpDao) GetResourcesFtpCount(resourcesFtpDto resources.ResourcesFtpDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_resourcesFtp_count, objectConvert.Struct2Map(resourcesFtpDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*ResourcesFtpDao) GetResourcesFtps(resourcesFtpDto resources.ResourcesFtpDto) ([]*resources.ResourcesFtpDto, error) {
	var resourcesFtpDtos []*resources.ResourcesFtpDto
	sqlTemplate.SelectList(query_resourcesFtp, objectConvert.Struct2Map(resourcesFtpDto), func(db *gorm.DB) {
		db.Scan(&resourcesFtpDtos)
	}, false)

	return resourcesFtpDtos, nil
}

/**
保存服务sql
*/
func (*ResourcesFtpDao) SaveResourcesFtp(resourcesFtpDto resources.ResourcesFtpDto) error {
	return sqlTemplate.Insert(insert_resourcesFtp, objectConvert.Struct2Map(resourcesFtpDto), false)
}

/**
修改服务sql
*/
func (*ResourcesFtpDao) UpdateResourcesFtp(resourcesFtpDto resources.ResourcesFtpDto) error {
	return sqlTemplate.Update(update_resourcesFtp, objectConvert.Struct2Map(resourcesFtpDto), false)
}

/**
删除服务sql
*/
func (*ResourcesFtpDao) DeleteResourcesFtp(resourcesFtpDto resources.ResourcesFtpDto) error {
	return sqlTemplate.Delete(delete_resourcesFtp, objectConvert.Struct2Map(resourcesFtpDto), false)
}
