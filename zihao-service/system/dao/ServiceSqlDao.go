package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/sqlTemplate"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)

const(
	query_service_sql_count string = `
		select count(1) total from service_sql t
		where 1=1
		  $if SqlId != '' then
			and t.sql_id = #SqlId#
		  $endif
		  $if SqlCode != '' then
			and t.sql_code = #SqlCode#
		  $endif
          and t.status_cd = '0'
		  order by t.create_time desc
    	
	`
	query_service_sql string = `
		select * from service_sql t
		where 1=1
		  $if SqlId != '' then
			and t.sql_id = #SqlId#
		  $endif
		  $if SqlCode != '' then
			and t.sql_code = #SqlCode#
		  $endif
          and t.status_cd = '0'
		  order by t.create_time desc
    	$if Page != -1 then
			limit #Page#,#Row#
		$endif
	`

	insert_service_sql string = `
	insert into service_sql( sql_code, sql_text, remark) 
VALUES (#SqlCode#,#SqlText#,#Remark#)
`

	update_service_sql string = `
	update service_sql t set t.sql_code = #SqlCode#, t.sql_text=#SqlText#,t.remark=#Remark#
	where t.sql_id = #SqlId#
	`
	delete_service_sql string = `
	update service_sql t set t.status_cd = '1'
	where t.sql_id = #SqlId#
	`
)

type ServiceSqlDao struct {

}

/**
查询用户
*/
func (*ServiceSqlDao) GetServiceSqlCount(serviceSqlDto serviceSql.ServiceSqlDto) (int64,error){
	var (
		pageDto dto.PageDto
		err error
	)

	sqlTemplate.SelectOne(query_service_sql_count,objectConvert.Struct2Map(serviceSqlDto), func(db *gorm.DB) {
		err  = db.Scan(&pageDto).Error
	},false)


	return pageDto.Total,err
}
/**
查询用户
*/
func (*ServiceSqlDao) GetServiceSqls(serviceSqlDto serviceSql.ServiceSqlDto) ([]*serviceSql.ServiceSqlDto,error){
	var serviceSqlDtos []*serviceSql.ServiceSqlDto
	sqlTemplate.SelectList(query_service_sql,objectConvert.Struct2Map(serviceSqlDto), func(db *gorm.DB) {
		db.Scan(&serviceSqlDtos)
	},false)

	return serviceSqlDtos,nil
}

/**
保存服务sql
*/
func (*ServiceSqlDao) SaveServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) error{
	return sqlTemplate.Insert(insert_service_sql,objectConvert.Struct2Map(serviceSqlDto),false)
}

/**
修改服务sql
*/
func (*ServiceSqlDao) UpdateServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) error{
	return sqlTemplate.Update(update_service_sql,objectConvert.Struct2Map(serviceSqlDto),false)
}

/**
删除服务sql
*/
func (*ServiceSqlDao) DeleteServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) error{
	return sqlTemplate.Delete(delete_service_sql,objectConvert.Struct2Map(serviceSqlDto),false)
}
