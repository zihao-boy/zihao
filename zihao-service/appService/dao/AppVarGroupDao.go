package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/appVarGroup"
)

const(
	query_appVarGroup_count string = `
		select count(1) total
		from app_var_group t
		where t.status_cd = '0'
		$if TenantId != '' then
		and t.tenant_id = #TenantId#
		$endif
		$if AvgId != '' then
		and t.avg_id = #AvgId#
		$endif
		$if AvgType != '' then
		and t.avg_type = #AvgType#
		$endif
    	
	`
	query_appVarGroup string = `
		
				select t.* from app_var_group t
				where t.status_cd = '0'
				$if TenantId != '' then
				and t.tenant_id = #TenantId#
				$endif
				$if AvgId != '' then
				and t.avg_id = #AvgId#
				$endif
				$if AvgType != '' then
				and t.avg_type = #AvgType#
				$endif
					order by t.create_time desc
					$if Row != 0 then
						limit #Page#,#Row#
					$endif
	`

	insert_appVarGroup string = `
	insert into app_var_group(avg_id, avg_name, avg_type, tenant_id, avg_desc) 
values(#AvgId#, #AvgName#, #AvgType#, #TenantId#, #AvgDesc#)
`

	update_appVarGroup string = `
	update app_var_group t set
	$if AvgName != '' then
	t.avg_name = #AvgName#,
	$endif
	$if AvgDesc != '' then
	t.avg_desc = #AvgDesc#,
	$endif
	t.status_cd ='0'
	where t.status_cd = '0'
	$if TenantId != '' then
	and t.tenant_id = #TenantId#
	$endif
	$if AvgId != '' then
	and t.avg_id = #AvgId#
	$endif
	`
	delete_appVarGroup string = `
	update app_var_group t set
                          t.status_cd = '1'
                          where t.status_cd = '0'

		$if AvgId != '' then
		and t.avg_id = #AvgId#
		$endif
	`
)

type AppVarGroupDao struct {

}

/**
查询用户
*/
func (*AppVarGroupDao) GetAppVarGroupCount(appVarGroupDto appVarGroup.AppVarGroupDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_appVarGroup_count,objectConvert.Struct2Map(appVarGroupDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*AppVarGroupDao) GetAppVarGroups(appVarGroupDto appVarGroup.AppVarGroupDto) ([]*appVarGroup.AppVarGroupDto,error){
	var appVarGroupDtos []*appVarGroup.AppVarGroupDto
	sqlTemplate.SelectList(query_appVarGroup,objectConvert.Struct2Map(appVarGroupDto), func(db *gorm.DB) {
		db.Scan(&appVarGroupDtos)
	},false)

	return appVarGroupDtos,nil
}

/**
保存服务sql
*/
func (*AppVarGroupDao) SaveAppVarGroup(appVarGroupDto appVarGroup.AppVarGroupDto) error{
	return sqlTemplate.Insert(insert_appVarGroup,objectConvert.Struct2Map(appVarGroupDto),false)
}

/**
修改服务sql
*/
func (*AppVarGroupDao) UpdateAppVarGroup(appVarGroupDto appVarGroup.AppVarGroupDto) error{
	return sqlTemplate.Update(update_appVarGroup,objectConvert.Struct2Map(appVarGroupDto),false)
}

/**
删除服务sql
*/
func (*AppVarGroupDao) DeleteAppVarGroup(appVarGroupDto appVarGroup.AppVarGroupDto) error{
	return sqlTemplate.Delete(delete_appVarGroup,objectConvert.Struct2Map(appVarGroupDto),false)
}
