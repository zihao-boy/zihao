package dao

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVersion"
	"gorm.io/gorm"
)

const (
	query_appVersion_count string = `
		select count(1) total
			from app_version t
			where t.status_cd = '0'
			$if AvId != '' then 
			and t.av_id = #AvId#
			$endif
			$if TenantId != '' then
				and t.tenant_id = #TenantId#
			$endif
    	
	`
	query_appVersion string = `
				select * from app_version t
					where t.status_cd = '0'
					$if AvId != '' then 
					and t.av_id = #AvId#
					$endif
					$if TenantId != '' then
						and t.tenant_id = #TenantId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersion string = `
	insert into app_version(av_id, name, remark, tenant_id) 
VALUES (#AvId#,#Name#,#Remark#,#TenantId#)
`

	update_appVersion string = `
	update app_version t set
		$if Name != '' then 
			t.name = #Name#,
		$endif
		$if Remark != '' then
		 t.remark = #Remark#,
		$endif
		t.status_cd = '0'
		where t.status_cd = '0'
		$if TenantId != '' then
		and t.tenant_id = #TenantId#
		$endif
		$if AvId != '' then
		and t.av_id = #AvId#
		$endif
	`
	delete_appVersion string = `
	update app_version t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if AvId != '' then
			and t.av_id = #AvId#
		$endif
	`
)

type AppVersionDao struct {
}

/**
查询用户
*/
func (*AppVersionDao) GetAppVersionCount(appVersionDto appVersion.AppVersionDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appVersion_count, objectConvert.Struct2Map(appVersionDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppVersionDao) GetAppVersions(appVersionDto appVersion.AppVersionDto) ([]*appVersion.AppVersionDto, error) {
	var appVersionDtos []*appVersion.AppVersionDto
	sqlTemplate.SelectList(query_appVersion, objectConvert.Struct2Map(appVersionDto), func(db *gorm.DB) {
		db.Scan(&appVersionDtos)
	}, false)

	return appVersionDtos, nil
}

/**
保存服务sql
*/
func (*AppVersionDao) SaveAppVersion(appVersionDto appVersion.AppVersionDto) error {
	return sqlTemplate.Insert(insert_appVersion, objectConvert.Struct2Map(appVersionDto), false)
}

/**
修改服务sql
*/
func (*AppVersionDao) UpdateAppVersion(appVersionDto appVersion.AppVersionDto) error {
	return sqlTemplate.Update(update_appVersion, objectConvert.Struct2Map(appVersionDto), false)
}

/**
删除服务sql
*/
func (*AppVersionDao) DeleteAppVersion(appVersionDto appVersion.AppVersionDto) error {
	return sqlTemplate.Delete(delete_appVersion, objectConvert.Struct2Map(appVersionDto), false)
}
