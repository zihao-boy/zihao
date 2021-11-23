package factory

import (
	"github.com/zihao-boy/zihao/zihao-service/common/cache/local"
	"github.com/zihao-boy/zihao/zihao-service/common/cache/redis"
	"github.com/zihao-boy/zihao/zihao-service/config"
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
