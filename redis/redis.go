// Author :		船長 <erichsiao@awoo.com.tw>

package redis

import (
	"errors"

	"github.com/go-redis/redis"
)

var (
	cfg *redis.Options
	db  *redis.Client
)

// setConfig : new redis options via go-models
// addr : must with port number. ex: '127.0.0.1:6379'
func setConfig(addr, pw string, db int) {

	if cfg != nil {
		panic("already had config !!")
	}

	cfg = &redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       db,
	}

	return
}

// setOriginConfig : set mysql config
func setOriginConfig(c *redis.Options) {
	if c != nil {
		cfg = c
	}
}

// getConfig : return the config
func getConfig() *redis.Options {
	return cfg
}

// openDB : open a new mysql connection
func openDB(c ...*redis.Options) (*redis.Client, error) {
	if len(c) > 0 && c[0] != nil {
		cfg = c[0]
	}

	if cfg == nil {
		return nil, errors.New("Do NewConfig() first")
	}

	if db != nil {
		return nil, errors.New("already connect to db")
	}

	db = redis.NewClient(cfg)

	return db, nil
}

// openDBWithPoolConns : open a new mysql connection and pool conns
func openDBWithPoolConns(active, idle int) (*redis.Client, error) {
	var err error

	// setting connections pool
	if cfg != nil {
		cfg.PoolSize = active
		cfg.MinIdleConns = idle
	}

	if db, err = openDB(); err != nil {
		return nil, err
	}

	return db, nil
}
