package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appService"
)

const(
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
	update app_service t set
		$if AsType != '' then
		 t.as_type = #AsType#,
		$endif
		$if AsName != '' then
		 t.as_name = #AsName#,
		$endif
		$if State != '' then
		 t.state = #State#,
		$endif
		$if AsCount != '' then
		 t.as_count = #AsCount#,
		$endif
		$if AsDesc != '' then
		 t.as_desc = #AsDesc#,
		$endif
		t.status_cd = '0'
		where t.status_cd = '0'
		$if TenantId != '' then
		and t.tenant_id = #TenantId#
		$endif
		$if AsId != '' then
		and t.as_id = #AsId#
		$endif
	`
	delete_appService string = `
	update app_service t set
                          t.status_cd = '1'
                          where t.status_cd = '0'

		$if AsId != '' then
		and t.as_id = #AsId#
		$endif
	`
)

type AppServiceDao struct {

}

/**
查询用户
*/
func (*AppServiceDao) GetAppServiceCount(appServiceDto appService.AppServiceDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_appService_count,objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*AppServiceDao) GetAppServices(appServiceDto appService.AppServiceDto) ([]*appService.AppServiceDto,error){
	var appServiceDtos []*appService.AppServiceDto
	sqlTemplate.SelectList(query_appService,objectConvert.Struct2Map(appServiceDto), func(db *gorm.DB) {
		db.Scan(&appServiceDtos)
	},false)

	return appServiceDtos,nil
}

/**
保存服务sql
*/
func (*AppServiceDao) SaveAppService(appServiceDto appService.AppServiceDto) error{
	return sqlTemplate.Insert(insert_appService,objectConvert.Struct2Map(appServiceDto),false)
}

/**
修改服务sql
*/
func (*AppServiceDao) UpdateAppService(appServiceDto appService.AppServiceDto) error{
	return sqlTemplate.Update(update_appService,objectConvert.Struct2Map(appServiceDto),false)
}

/**
删除服务sql
*/
func (*AppServiceDao) DeleteAppService(appServiceDto appService.AppServiceDto) error{
	return sqlTemplate.Delete(delete_appService,objectConvert.Struct2Map(appServiceDto),false)
}
