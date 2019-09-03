// Author :		船長 <erichsiao@awoo.com.tw>

package redis

import (
	"github.com/go-redis/redis"
)

// NewOptions : new redis options via go-models
// addr : must with port number. ex: '127.0.0.1:6379'
func NewOptions(addr, pw string, db int) (opt *redis.Options) {
	opt = &redis.Options{
		Addr:         addr,
		Password:     pw,
		DB:           db,
		PoolSize:     defaultOpenConns,
		MinIdleConns: defaultIdleConns,
	}

	return
}

// NewRedis : create a new redis dao
func NewRedis(opt *redis.Options) *Dao {
	return newDao(opt)
}
