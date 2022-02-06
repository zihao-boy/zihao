package installAppDao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/installApp"
	"gorm.io/gorm"
)

const (
	query_installApp_count string = `
	select count(1) total
	from install_app t
	where t.status_cd = '0'
	$if AppId != '' then
	and t.app_id = #AppId#
	$endif
	$if AppName != '' then
	and t.app_name = #AppName#
	$endif
	$if ExtAppId != '' then
	and t.ext_app_id = #ExtAppId#
	$endif
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
    	
	`
	query_installApp string = `
				select t.*
				from install_app t
				where t.status_cd = '0'
				$if AppId != '' then
				and t.app_id = #AppId#
				$endif
				$if AppName != '' then
				and t.app_name = #AppName#
				$endif
				$if ExtAppId != '' then
				and t.ext_app_id = #ExtAppId#
				$endif
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_installApp string = `
	insert into install_app(app_id, app_name, version, ext_app_id,tenant_id,create_user_id)
VALUES(#AppId#,#AppName#,#Version#,#ExtAppId#,#TenantId#,#CreateUserId#)
`

	update_installApp string = `
	update install_app set
		$if AppName != '' then
		app_name = #AppName#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if AppId != '' then
		and app_id = #AppId#
		$endif
	`
	delete_installApp string = `
	update install_app  set
                          status_cd = '1'
                          where status_cd = '0'
						  $if Id != '' then
						  and id = #Id#
						  $endif
	`
)

type InstallAppDao struct {
}

/**
查询用户
*/
func (*InstallAppDao) GetInstallAppCount(installAppDto installApp.InstallAppDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_installApp_count, objectConvert.Struct2Map(installAppDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*InstallAppDao) GetInstallApps(installAppDto installApp.InstallAppDto) ([]*installApp.InstallAppDto, error) {
	var installAppDtos []*installApp.InstallAppDto
	sqlTemplate.SelectList(query_installApp, objectConvert.Struct2Map(installAppDto), func(db *gorm.DB) {
		db.Scan(&installAppDtos)
	}, false)

	return installAppDtos, nil
}

/**
保存服务sql
*/
func (*InstallAppDao) SaveInstallApp(installAppDto installApp.InstallAppDto) error {
	return sqlTemplate.Insert(insert_installApp, objectConvert.Struct2Map(installAppDto), false)
}

/**
修改服务sql
*/
func (*InstallAppDao) UpdateInstallApp(installAppDto installApp.InstallAppDto) error {
	return sqlTemplate.Update(update_installApp, objectConvert.Struct2Map(installAppDto), false)
}

/**
删除服务sql
*/
func (*InstallAppDao) DeleteInstallApp(installAppDto installApp.InstallAppDto) error {
	return sqlTemplate.Delete(delete_installApp, objectConvert.Struct2Map(installAppDto), false)
}
