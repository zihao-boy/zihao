package factory

import (
	"github.com/zihao-boy/zihao/common/cache/local"
	"github.com/zihao-boy/zihao/common/cache/redis"
	"github.com/zihao-boy/zihao/config"
	"github.com/zihao-boy/zihao/entity/dto/mapping"
	"github.com/zihao-boy/zihao/entity/dto/serviceSql"
)

type CacheFactory struct {
}

const (
	Cache_redis = "redis"
	Cache_local = "local"
)

func Init() {
	cacheSwatch := config.G_AppConfig.Cache

	if Cache_redis == cacheSwatch {
		redis.InitRedis()
	}

	if Cache_local == cacheSwatch {
		local.InitLocal()
	}

}

func SetToken(format, id string, token string) (err error) {
	cacheSwatch := config.G_AppConfig.Cache

	if Cache_redis == cacheSwatch {
		err = redis.G_Redis.SetToken(format, id, token)
	}

	if Cache_local == cacheSwatch {
		err = local.G_Local.SetToken(format, id, token)
	}
	return
}

func GetToken(format, id string) (token string, err error) {
	cacheSwatch := config.G_AppConfig.Cache

	if Cache_redis == cacheSwatch {
		token, err = redis.G_Redis.GetToken(format, id)
	}

	if Cache_local == cacheSwatch {
		token, err = local.G_Local.GetToken(format, id)
	}
	return
}

func DelToken(format, id string) (result int64, err error) {
	cacheSwatch := config.G_AppConfig.Cache

	if Cache_redis == cacheSwatch {
		result, err = redis.G_Redis.DelToken(format, id)
	}

	if Cache_local == cacheSwatch {
		result, err = local.G_Local.DelToken(format, id)
	}
	return
}

func SaveServiceSql(serviceSqlDto serviceSql.ServiceSqlDto) (err error) {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		err = redis.G_Redis.SaveServiceSql(serviceSqlDto)
	}
	if Cache_local == cacheSwatch {
		err = local.G_Local.SaveServiceSql(serviceSqlDto)
	}
	return err
}

func GetServiceSql(sqlCode string) (serviceSql serviceSql.ServiceSqlDto, err error) {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		serviceSql, err = redis.G_Redis.GetServiceSql(sqlCode)
	}
	if Cache_local == cacheSwatch {
		serviceSql, err = local.G_Local.GetServiceSql(sqlCode)
	}
	return serviceSql, err
}

func SetValue(key string, value string, timeout ...int64) (err error) {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		err = redis.G_Redis.SetValue(key, value, timeout...)
	}
	if Cache_local == cacheSwatch {
		err = local.G_Local.SetValue(key, value, timeout...)
	}
	return
}

func GetValue(key string) (string, error) {
	var value string
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		value, _ = redis.G_Redis.GetValue(key)
	}
	if Cache_local == cacheSwatch {
		value, _ = local.G_Local.GetValue(key)
	}

	return value, nil
}

func GetValueAndRemove(key string) (string, error) {
	var token string
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		token, _ = redis.G_Redis.GetValueAndRemove(key)
	}
	if Cache_local == cacheSwatch {
		token, _ = local.G_Local.GetValueAndRemove(key)
	}
	return token, nil
}

// Init
func InitServiceSql() {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		redis.InitServiceSql()
	}

	if Cache_local == cacheSwatch {
		local.InitServiceSql()
	}
}

func SaveMapping(mappingDto mapping.MappingDto) (err error) {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		err = redis.G_Redis.SaveMapping(mappingDto)
	}
	if Cache_local == cacheSwatch {
		err = local.G_Local.SaveMapping(mappingDto)
	}
	return err
}

func GetMapping(zKey string) (mapping mapping.MappingDto, err error) {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		mapping, err = redis.G_Redis.GetMapping(zKey)
	}
	if Cache_local == cacheSwatch {
		mapping, err = local.G_Local.GetMapping(zKey)
	}
	return mapping, err
}

// Init
func InitMapping() {
	cacheSwatch := config.G_AppConfig.Cache
	if Cache_redis == cacheSwatch {
		redis.InitMapping()
	}

	if Cache_local == cacheSwatch {
		local.InitMapping()
	}
}
