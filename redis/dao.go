// Author :		船長 <erichsiao@awoo.com.tw>

package redis

import (
	"github.com/go-redis/redis"
)

type Dao struct {
	*redis.Client
}

func newDao(opt *redis.Options) *Dao {
	return &Dao{
		Client: redis.NewClient(opt),
	}
}
