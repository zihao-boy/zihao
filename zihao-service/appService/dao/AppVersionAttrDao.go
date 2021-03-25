package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVersion"
)

const(
	query_appVersionAttr_count string = `
		select count(1) total
			from app_version_attr t
			where t.status_cd = '0'
			$if AvId != '' then
			and t.av_id = #AvId#
			$endif
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AttrId != '' then
			and t.attr_id = #AttrId#
			$endif
    	
	`
	query_appVersionAttr string = `
				select * from app_version_attr t
					where t.status_cd = '0'
					$if AvId != '' then
					and t.av_id = #AvId#
					$endif
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AttrId != '' then
					and t.attr_id = #AttrId#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVersionAttr string = `
	insert into app_version_attr(attr_id, av_id, version, tenant_id) 
values (#AttrId#,#AvId#,#Version#,#TenantId#)
`

	update_appVersionAttr string = `
	update app_version_attr t set 
		  t.version = #Version#
		where
		t.status_cd = '0'
		$if AvId != '' then
		and t.av_id = #AvId#
		$endif
		$if TenantId != '' then
		and t.tenant_id = #TenantId#
		$endif
		$if AttrId != '' then
		and t.attr_id = #AttrId#
		$endif
	`
	delete_appVersionAttr string = `
	update app_version_attr t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if AttrId != '' then
			and t.attr_id = #AttrId#
		$endif
	`
)

type AppVersionAttrDao struct {

}

/**
查询用户
*/
func (*AppVersionAttrDao) GetAppVersionAttrCount(appVersionAttrDto appVersion.AppVersionAttrDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_appVersionAttr_count,objectConvert.Struct2Map(appVersionAttrDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*AppVersionAttrDao) GetAppVersionAttrs(appVersionAttrDto appVersion.AppVersionAttrDto) ([]*appVersion.AppVersionAttrDto,error){
	var appVersionAttrDtos []*appVersion.AppVersionAttrDto
	sqlTemplate.SelectList(query_appVersionAttr,objectConvert.Struct2Map(appVersionAttrDto), func(db *gorm.DB) {
		db.Scan(&appVersionAttrDtos)
	},false)

	return appVersionAttrDtos,nil
}

/**
保存服务sql
*/
func (*AppVersionAttrDao) SaveAppVersionAttr(appVersionAttrDto appVersion.AppVersionAttrDto) error{
	return sqlTemplate.Insert(insert_appVersionAttr,objectConvert.Struct2Map(appVersionAttrDto),false)
}

/**
修改服务sql
*/
func (*AppVersionAttrDao) UpdateAppVersionAttr(appVersionAttrDto appVersion.AppVersionAttrDto) error{
	return sqlTemplate.Update(update_appVersionAttr,objectConvert.Struct2Map(appVersionAttrDto),false)
}

/**
删除服务sql
*/
func (*AppVersionAttrDao) DeleteAppVersionAttr(appVersionAttrDto appVersion.AppVersionAttrDto) error{
	return sqlTemplate.Delete(delete_appVersionAttr,objectConvert.Struct2Map(appVersionAttrDto),false)
}
