package utils

import (
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)

func GetServiceSql(sqlCode string) serviceSql.ServiceSqlDto  {
	serviceSqlDto,_ := redis.G_Redis.GetServiceSql(sqlCode)
	return serviceSqlDto
}
