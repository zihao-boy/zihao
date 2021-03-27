package redis

import (
	"sync"
	"time"
)

const (
	Default_timeOut int64 = 0
)
var mutex sync.Mutex

func (r *Redis) SetValue(key string, value string,timeout ...int64) (err error) {
	var (
		tt int64 = Default_timeOut
	)
	if len(timeout) >0{
		tt = timeout[0]
	}
	err = r.client.Set(key, value,
		time.Duration(tt)*time.Minute).Err()
	return
}

func (r *Redis) GetValue(key string) ( string,  error) {
	token, err := r.client.Get(key).Result()
	return token, err
}

func (r *Redis) GetValueAndRemove(key string) (string,error)  {
	mutex.Lock()
	token, err := r.client.Get(key).Result()
	r.client.Del(key)
	mutex.Unlock()
	return token, err
}



