package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVar"
)

const(
	query_appVar_count string = `
		select count(1) total
		  from app_var t
		where t.status_cd ='0'
		$if AvgId != '' then
		and t.avg_id = #AvgId#
		$endif
		$if TenantId != '' then
		and t.tenant_id = #TenantId#
		$endif
		$if AvId != '' then
		and t.av_id = #AvId#
		$endif
		$if VarType != '' then
		and t.var_type = #VarType#
		$endif
		$if VarSpec != '' then
		 and t.var_spec = #VarSpec#
		$endif
    	
	`
	query_appVar string = `
				select t.*,avg.avg_name from app_var t
				left join app_var_group avg on t.avg_id = avg.avg_id and avg.status_cd = '0'
				where t.status_cd ='0'
					$if AvgId != '' then
					and t.avg_id = #AvgId#
					$endif
					$if TenantId != '' then
					and t.tenant_id = #TenantId#
					$endif
					$if AvId != '' then
					and t.av_id = #AvId#
					$endif
					$if VarType != '' then
					and t.var_type = #VarType#
					$endif
					$if VarSpec != '' then
					 and t.var_spec = #VarSpec#
					$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVar string = `

insert into app_var(av_id, avg_id, tenant_id, var_name, var_type,var_spec) 
VALUES (#AvId#, #AvgId#, #TenantId#, #VarName#, #VarType#,#VarSpec#) 
`

	update_appVar string = `
	update app_var t set
	$if VarName != '' then
	  t.var_name = #VarName#,
	$endif
	$if VarType != '' then
	 t.var_type = #VarType#,
	$endif
	$if VarSpec != '' then
	 t.var_spec = #VarSpec#,
	$endif
	t.status_cd ='0'
	where t.status_cd ='0'
	$if AvgId != '' then
	and t.avg_id = #AvgId#
	$endif
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if AvId != '' then
	and t.av_id = #AvId#
	$endif
	`
	delete_appVar string = `
	update app_var t set
                          t.status_cd = '1'
                          where t.status_cd = '0'
		$if AvId != '' then
	and t.av_id = #AvId#
	$endif
	`
)

type AppVarDao struct {

}

/**
查询用户
*/
func (*AppVarDao) GetAppVarCount(appVarDto appVar.AppVarDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_appVar_count,objectConvert.Struct2Map(appVarDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*AppVarDao) GetAppVars(appVarDto appVar.AppVarDto) ([]*appVar.AppVarDto,error){
	var appVarDtos []*appVar.AppVarDto
	sqlTemplate.SelectList(query_appVar,objectConvert.Struct2Map(appVarDto), func(db *gorm.DB) {
		db.Scan(&appVarDtos)
	},false)

	return appVarDtos,nil
}

/**
保存服务sql
*/
func (*AppVarDao) SaveAppVar(appVarDto appVar.AppVarDto) error{
	return sqlTemplate.Insert(insert_appVar,objectConvert.Struct2Map(appVarDto),false)
}

/**
修改服务sql
*/
func (*AppVarDao) UpdateAppVar(appVarDto appVar.AppVarDto) error{
	return sqlTemplate.Update(update_appVar,objectConvert.Struct2Map(appVarDto),false)
}

/**
删除服务sql
*/
func (*AppVarDao) DeleteAppVar(appVarDto appVar.AppVarDto) error{
	return sqlTemplate.Delete(delete_appVar,objectConvert.Struct2Map(appVarDto),false)
}
