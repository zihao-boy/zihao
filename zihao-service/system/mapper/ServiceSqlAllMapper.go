package mapper

import (
	"github.com/zihao-boy/zihao/zihao-service/common/db/mysql"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)



type ServiceSqlAllMapper struct {

}

/**
查询用户


*/
func (*ServiceSqlAllMapper) GetServiceSqls(serviceSqlDto serviceSql.ServiceSqlDto) ([]*serviceSql.ServiceSqlDto,error){
	var serviceSqlDtos []*serviceSql.ServiceSqlDto
	db := mysql.G_DB.Raw("select * from service_sql")
	if err:=db.Scan(&serviceSqlDtos).Error; err !=nil{
		return nil,err
	}

	return serviceSqlDtos,nil
}

