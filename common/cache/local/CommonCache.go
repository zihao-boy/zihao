package local

import (
	"sync"
	"time"
)

const (
	Default_timeOut int64 = 0
)

var mutex sync.Mutex

func (r *Local) SetValue(key string, value string, timeout ...int64) (err error) {
	var (
		tt int64 = Default_timeOut
	)
	if len(timeout) > 0 {
		tt = timeout[0]
	}
	r.client.Set(key, value,
		time.Duration(tt)*time.Minute)
	return
}

func (r *Local) GetValue(key string) (string, error) {
	value, _ := r.client.Get(key)

	return value.(string), nil
}

func (r *Local) GetValueAndRemove(key string) (string, error) {
	mutex.Lock()
	token, _ := r.client.Get(key)
	if token == "" {
		return "", nil
	}
	r.client.Delete(key)
	mutex.Unlock()
	return token.(string), nil
}
