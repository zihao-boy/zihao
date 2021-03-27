package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/tenant"
)

const(
	query_tenant_count string = `
		select count(1) total from tenant t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if TenantName != '' then
					and t.tenant_name = #TenantName#
					$endif
					$if TenantType != '' then
					and t.tenant_type = #TenantType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
    	
	`
	query_tenant string = `
		select t.*,uu.username from tenant t
				left join u_user uu on t.tenant_id = uu.tenant_id and uu.user_role = '1001'
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if TenantName != '' then
					and t.tenant_name = #TenantName#
					$endif
					$if TenantType != '' then
					and t.tenant_type = #TenantType#
					$endif
					$if Phone != '' then
					and t.phone = #Phone#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					order by t.create_time desc
					$if Page != -1 then
						limit #Page#,#Row#
					$endif
	`

	insert_tenant string = `
insert into tenant(tenant_id, tenant_name, address, person_name, phone, remark) 
VALUES(#TenantId#, #TenantName#, #Address#, #PersonName#, #Phone#, #Remark#) 
`

	update_tenant string = `
	update tenant t set
			$if TenantName != '' then
			 t.tenant_name = #TenantName#,
			$endif
			$if TenantType != '' then
			 t.tenant_type = #TenantType#,
			$endif
			$if Phone != '' then
			 t.phone = #Phone#,
			$endif
			$if State != '' then
			 t.state = #State#,
			$endif
			$if Address != '' then
			 t.address = #Address#,
			$endif
			$if PersonName != '' then
			 t.person_name = #PersonName#,
			$endif
			$if Remark != '' then
			 t.remark = #Remark#,
			$endif
			t.status_cd = '0'
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
	`
	delete_tenant string = `
	update tenant t set
			t.status_cd = '1'
			where t.status_cd = '0'
			and t.tenant_id = #TenantId#
	`
)

type TenantDao struct {

}

/**
查询用户
*/
func (*TenantDao) GetTenantCount(tenantDto tenant.TenantDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_tenant_count,objectConvert.Struct2Map(tenantDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*TenantDao) GetTenants(tenantDto tenant.TenantDto) ([]*tenant.TenantDto,error){
	var tenantDtos []*tenant.TenantDto
	sqlTemplate.SelectList(query_tenant,objectConvert.Struct2Map(tenantDto), func(db *gorm.DB) {
		db.Scan(&tenantDtos)
	},false)

	return tenantDtos,nil
}

/**
保存服务sql
*/
func (*TenantDao) SaveTenant(tenantDto tenant.TenantDto) error{
	return sqlTemplate.Insert(insert_tenant,objectConvert.Struct2Map(tenantDto),false)
}

/**
修改服务sql
*/
func (*TenantDao) UpdateTenant(tenantDto tenant.TenantDto) error{
	return sqlTemplate.Update(update_tenant,objectConvert.Struct2Map(tenantDto),false)
}

/**
删除服务sql
*/
func (*TenantDao) DeleteTenant(tenantDto tenant.TenantDto) error{
	return sqlTemplate.Delete(delete_tenant,objectConvert.Struct2Map(tenantDto),false)
}
