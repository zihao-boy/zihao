package dao

import (
	"github.com/zihao-boy/zihao/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/common/objectConvert"
	"github.com/zihao-boy/zihao/entity/dto"
	"github.com/zihao-boy/zihao/entity/dto/appService"
	"gorm.io/gorm"
)

const (
	query_appService_count string = `
		select count(1) total
			from app_service t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if AsType != '' then
			and t.as_type = #AsType#
			$endif
			$if State != '' then
			and t.state = #State#
			$endif
			$if AsName != '' then
			and t.as_name = #AsName#
			$endif
    	
	`
	query_appService string = `
		
				select t.*,td.name as_type_name,td1.name state_name from app_service t
left join t_dict td on t.as_type = td.status_cd and td.table_name = 'app_service' and td.table_columns = 'as_type'
left join t_dict td1 on t.state = td1.status_cd and td1.table_name = 'app_service' and td1.table_columns = 'state'
where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if AsType != '' then
					and t.as_type = #AsType#
					$endif
					$if State != '' then
					and t.state = #State#
					$endif
					$if AsName != '' then
					and t.as_name = #AsName#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appService string = `
	insert into app_service(as_id, as_name, as_type, tenant_id, as_desc,state,as_count)
VALUES(#AsId#,#AsName#,#AsType#,#TenantId#,#AsDesc#,#State#,#AsCount#)
`

	update_appService string = `
	update app_service set
		$if AsType != '' then
		 as_type = #AsType#,
		$endif
		$if AsName != '' then
		 as_name = #AsName#,
		$endif
		$if State != '' then
		 state = #State#,
		$endif
		$if AsCount != '' then
		 as_count = #AsCount#,
		$endif
		$if AsDesc != '' then
		 as_desc = #AsDesc#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if AsId != '' then
		and as_id = #AsId#
		$endif
	`
	delete_appService string = `
	update app_service  set
                          status_cd = '1'
                          where status_cd = '0'

		$if AsId != '' then
		and as_id = #AsId#
		$endif
	`



	query_appServiceVar_count string = `
		select count(1) total
			from app_service_var t
			where t.status_cd = '0'
			$if TenantId != '' then
			and t.tenant_id = #TenantId#
			$endif
			$if AsId != '' then
			and t.as_id = #AsId#
			$endif
			$if AvId != '' then
			and t.av_id = #AvId#
			$endif
			$if VarName != '' then
			and t.var_name = #VarName#
			$endif
			$if VarSpec != '' then
			and t.var_spec = #VarSpec#
			$endif
    	
	`
	query_appServiceVar string = `
				select t.*
				from app_service_var t
					where t.status_cd = '0'
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AsId != '' then
					and t.as_id = #AsId#
					$endif
					$if AvId != '' then
					and t.av_id = #AvId#
					$endif
					$if VarName != '' then
					and t.var_name = #VarName#
					$endif
					$if VarSpec != '' then
					and t.var_spec = #VarSpec#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appServiceVar string = `
	insert into app_service_var(av_id, as_id, tenant_id, var_spec, var_name,var_value)
VALUES(#AvId#,#AsId#,#TenantId#,#VarSpec#,#VarName#,#VarValue#)
`

	update_appServiceVar string = `
	update app_service_var set
		$if VarSpec != '' then
		 var_spec = #VarSpec#,
		$endif
		$if VarName != '' then
		 var_name = #VarName#,
		$endif
		$if VarValue != '' then
		 var_value = #VarValue#,
		$endif
		status_cd = '0'
		where status_cd = '0'
		$if TenantId != '' then
		and tenant_id = #TenantId#
		$endif
		$if AvId != '' then
		and av_id = #AvId#
		$endif
	`
	delete_appServiceVar string = `
	update app_service_var  set
                          status_cd = '1'
                          where status_cd = '0'

		$if AvId != '' then
		and av_id = #AvId#
		$endif
	`
)

type AppServiceDao struct {
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceCount(appServiceDto appService.AppServiceDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appService_count, objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServices(appServiceDto appService.AppServiceDto) ([]*appService.AppServiceDto, error) {
	var appServiceDtos []*appService.AppServiceDto
	sqlTemplate.SelectList(query_appService, objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		db.Scan(&appServiceDtos)
	}, false)

	return appServiceDtos, nil
}

/**
保存服务sql
*/
func (*AppServiceDao) SaveAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Insert(insert_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
修改服务sql
*/
func (*AppServiceDao) UpdateAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Update(update_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
删除服务sql
*/
func (*AppServiceDao) DeleteAppService(appServiceDto appService.AppServiceDto) error {
	return sqlTemplate.Delete(delete_appService, objectConvert.Struct2Map(appServiceDto), false)
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceVarCount(appServiceVarDto appService.AppServiceVarDto) (int64, error) {
	var (
		pageDto dto.PageDto
		err     error
	)

	sqlTemplate.SelectOne(query_appServiceVar_count, objectConvert.Struct2Map(appServiceVarDto), func(db *gorm.DB) {
		err = db.Scan(&pageDto).Error
	}, false)

	return pageDto.Total, err
}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceVars(appServiceVarDto appService.AppServiceVarDto) ([]*appService.AppServiceVarDto, error) {
	var appServiceVarDtos []*appService.AppServiceVarDto
	sqlTemplate.SelectList(query_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), func(db *gorm.DB) {
		db.Scan(&appServiceVarDtos)
	}, false)

	return appServiceVarDtos, nil
}

/**
保存服务sql
*/
func (*AppServiceDao) SaveAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Insert(insert_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}

/**
修改服务sql
*/
func (*AppServiceDao) UpdateAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Update(update_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}

/**
删除服务sql
*/
func (*AppServiceDao) DeleteAppServiceVar(appServiceVarDto appService.AppServiceVarDto) error {
	return sqlTemplate.Delete(delete_appServiceVar, objectConvert.Struct2Map(appServiceVarDto), false)
}