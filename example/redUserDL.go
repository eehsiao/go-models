// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"github.com/eehsiao/go-models/redis"
)

// RedUserModel : extend from redis.Dao
type RedUserModel struct {
	*redis.Dao
}

// User : json struct that to store into redis
type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	SelectPriv string `json:"select_priv"`
}

// UserHMSet : this is a data logical function, you can write more logical in there
// sample function of the data logical
func (r *RedUserModel) UserHMSet(hKey string, kv map[string]interface{}) (status string, err error) {
	if kv != nil && len(kv) > 0 {
		return r.HMSet(hKey, kv).Result()
	}

	return
}
