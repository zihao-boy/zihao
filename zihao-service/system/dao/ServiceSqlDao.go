package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/common/objectConvert"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)

const(
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
	`
)

type ServiceSqlDao struct {

}

/**
查询用户
*/
func (*ServiceSqlDao) GetServiceSqls(serviceSqlDto serviceSql.ServiceSqlDto) ([]*serviceSql.ServiceSqlDto,error){
	var serviceSqlDtos []*serviceSql.ServiceSqlDto
	mysql.SelectList(query_service_sql,objectConvert.Struct2Map(serviceSqlDto), func(db *gorm.DB) {
		db.Scan(&serviceSqlDtos)
	})

	return serviceSqlDtos,nil
}
