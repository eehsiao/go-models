// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"github.com/eehsiao/go-models/redis"
)

type RedUserModel struct {
	*redis.Dao
}

type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	SelectPriv string `json:"select_priv"`
}

// UserHMSet : this is a data logical function, you can write more logical in there
// sample data logical function
func (r *RedUserModel) UserHMSet(hKey string, kv map[string]interface{}) (status string, err error) {
	if kv != nil && len(kv) > 0 {
		return r.HMSet(hKey, kv).Result()
	}

	return
}
