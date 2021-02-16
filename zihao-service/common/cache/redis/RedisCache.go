package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/kataras/golog"
	"goiris/common"
	"time"
)

var G_Redis *Redis

type Redis struct {
	client *redis.Client
}

func (r *Redis) SetToken(format, id string, token string) (err error) {
	err = r.client.Set(fmt.Sprintf(format, id), token,
		time.Duration(common.G_AppConfig.JWTTimeout)*time.Minute).Err()
	return
}

func (r *Redis) GetToken(format, id string) (token string, err error) {
	token, err = r.client.Get(fmt.Sprintf(format, id)).Result()
	return
}

func (r *Redis) DelToken(format, id string) (result int64, err error) {
	result, err = r.client.Del(fmt.Sprintf(format, id)).Result()
	return
}

// Init
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     common.G_DBConfig.Redis.Addr,
		Password: common.G_DBConfig.Redis.Password,
		DB:       common.G_DBConfig.Redis.DB, // 连接的库位
		PoolSize: common.G_DBConfig.Redis.PoolSize,
	})
	if _, err := client.Ping().Result(); err != nil {
		golog.Fatalf("~~> Redis初始化错误,原因:%s", err.Error())
	}
	G_Redis = &Redis{
		client: client,
	}
	return
}

