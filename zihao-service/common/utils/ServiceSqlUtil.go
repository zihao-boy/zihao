package utils

import (
	"github.com/zihao-boy/zihao/zihao-service/common/cache/local"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/common/encrypt"
	"github.com/zihao-boy/zihao/zihao-service/config"
	"github.com/zihao-boy/zihao/zihao-service/entity/dto/serviceSql"
)

const (
	Cache_redis = "redis"
	Cache_local = "local"
)

func GetServiceSql(sqlCode string) serviceSql.ServiceSqlDto {
	cacheSwatch := config.G_AppConfig.Cache
	var (
		serviceSqlDto serviceSql.ServiceSqlDto
	)
	if Cache_redis == cacheSwatch {
		serviceSqlDto, _ = redis.G_Redis.GetServiceSql(sqlCode)
	}

	if Cache_local == cacheSwatch {
		serviceSqlDto, _ = local.G_Local.GetServiceSql(sqlCode)
	}
	serviceSqlDto.SqlText = encrypt.Decode(serviceSqlDto.SqlText)
	return serviceSqlDto
}
