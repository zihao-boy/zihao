package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/tenant"
)

const(
	query_tenantSetting_count string = `
		select count(1) total 
					from tenant_setting t
					where
					t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if SettingId != '' then
					and t.setting_id = #SettingId#
					$endif
					$if SpecCd != '' then
					and t.spec_cd = #SpecCd#
					$endif
    	
	`
	query_tenantSetting string = `

					select * 
					from tenant_setting t
					where
					t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if SettingId != '' then
					and t.setting_id = #SettingId#
					$endif
					$if SpecCd != '' then
					and t.spec_cd = #SpecCd#
					$endif
					order by t.create_time desc
					$if Page != -1 then
						limit #Page#,#Row#
					$endif
	`

	insert_tenantSetting string = `
insert into tenant_setting(setting_id, tenant_id, spec_cd, value) 
VALUES (#SettingId#,#TenantId#,#SpecCd#,#Value#)
`

	update_tenantSetting string = `
	update tenant_setting t set t.value = #Value#
where t.setting_id = #SettingId#
	`
	delete_tenantSetting string = `
	update tenant_setting t set 
			t.status_cd = '1'
			where t.status_cd = '0'
			and t.setting_id = #SettingId#
	`
)

type TenantSettingDao struct {

}

/**
查询用户
*/
func (*TenantSettingDao) GetTenantSettingCount(tenantSettingDto tenant.TenantSettingDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_tenantSetting_count,objectConvert.Struct2Map(tenantSettingDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*TenantSettingDao) GetTenantSettings(tenantSettingDto tenant.TenantSettingDto) ([]*tenant.TenantSettingDto,error){
	var tenantSettingDtos []*tenant.TenantSettingDto
	sqlTemplate.SelectList(query_tenantSetting,objectConvert.Struct2Map(tenantSettingDto), func(db *gorm.DB) {
		db.Scan(&tenantSettingDtos)
	},false)

	return tenantSettingDtos,nil
}

/**
保存服务sql
*/
func (*TenantSettingDao) SaveTenantSetting(tenantSettingDto tenant.TenantSettingDto) error{
	return sqlTemplate.Insert(insert_tenantSetting,objectConvert.Struct2Map(tenantSettingDto),false)
}

/**
修改服务sql
*/
func (*TenantSettingDao) UpdateTenantSetting(tenantSettingDto tenant.TenantSettingDto) error{
	return sqlTemplate.Update(update_tenantSetting,objectConvert.Struct2Map(tenantSettingDto),false)
}

/**
删除服务sql
*/
func (*TenantSettingDao) DeleteTenantSetting(tenantSettingDto tenant.TenantSettingDto) error{
	return sqlTemplate.Delete(delete_tenantSetting,objectConvert.Struct2Map(tenantSettingDto),false)
}
