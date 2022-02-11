package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/tenant"
	"gorm.io/gorm"
)

const (
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

					select t.*,td.name spec_cd_name
from tenant_setting t
left join t_dict td on t.spec_cd = td.status_cd and td.table_name = 'tenant_setting' and td.table_columns = 'spec_cd'

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
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_tenantSetting string = `
insert into tenant_setting(setting_id, tenant_id, spec_cd, value) 
VALUES (#SettingId#,#TenantId#,#SpecCd#,#Value#)
`

	update_tenantSetting string = `
	update tenant_setting set value = #Value#
where setting_id = #SettingId#
	`
	delete_tenantSetting string = `
	update tenant_setting  set 
			status_cd = '1'
			where status_cd = '0'
			and setting_id = #SettingId#
	`
)

type TenantSettingDao struct {
}

/**
查询用户
*/
func (*TenantSettingDao) GetTenantSettingCount(tenantSettingDto tenant.TenantSettingDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_tenantSetting_count, objectConvert.Struct2Map(tenantSettingDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*TenantSettingDao) GetTenantSettings(tenantSettingDto tenant.TenantSettingDto) ([]*tenant.TenantSettingDto, error) {
	var tenantSettingDtos []*tenant.TenantSettingDto
	sqlTemplate.SelectList(query_tenantSetting, objectConvert.Struct2Map(tenantSettingDto), func(db *gorm.DB) {
		db.Scan(&tenantSettingDtos)
	}, false)

	return tenantSettingDtos, nil
}

/**
保存服务sql
*/
func (*TenantSettingDao) SaveTenantSetting(tenantSettingDto tenant.TenantSettingDto) error {
	return sqlTemplate.Insert(insert_tenantSetting, objectConvert.Struct2Map(tenantSettingDto), false)
}

/**
修改服务sql
*/
func (*TenantSettingDao) UpdateTenantSetting(tenantSettingDto tenant.TenantSettingDto) error {
	return sqlTemplate.Update(update_tenantSetting, objectConvert.Struct2Map(tenantSettingDto), false)
}

/**
删除服务sql
*/
func (*TenantSettingDao) DeleteTenantSetting(tenantSettingDto tenant.TenantSettingDto) error {
	return sqlTemplate.Delete(delete_tenantSetting, objectConvert.Struct2Map(tenantSettingDto), false)
}
