package cacheFactory

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

func (cache *CacheFactory) Init() {
	cacheSwatch := config.G_AppConfig.Cache

	if Cache_redis == cacheSwatch {
		redis.InitRedis()
	}

	if Cache_local == cacheSwatch {
		local.InitLocal()
	}

}
