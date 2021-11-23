package local

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zihao-boy/zihao/zihao-service/config"
)

var G_Local *Local

type Local struct {
	client *cache.Cache
}

func (r *Local) SetToken(format string, id string, token string) (err error) {
	r.client.Set(fmt.Sprintf(format, id), token,
		time.Duration(config.G_AppConfig.JWTTimeout)*time.Minute)
	return
}

func (r *Local) GetToken(format, id string) (token string, err error) {
	value, _ := r.client.Get(fmt.Sprintf(format, id))
	if value == nil {
		token = ""
		return
	}
	token = value.(string)
	return
}

func (r *Local) DelToken(format, id string) (result int64, err error) {
	r.client.Delete(fmt.Sprintf(format, id))
	return
}

// Init 初始化本地local
func InitLocal() {
	client := cache.New(-1, -1)
	G_Local = &Local{
		client: client,
	}
	return
}
