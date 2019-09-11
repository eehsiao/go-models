// Author :		船長 <erichsiao@awoo.com.tw>

package redis

import (
	"errors"
	"reflect"

	"github.com/go-redis/redis"
)

type Dao struct {
	*redis.Client
	DaoStruct     string
	DaoStructType reflect.Type
	DbNnum        int
	DataKey       string
}

func NewDao() *Dao {
	return &Dao{}
}

// SetDefaultModel : register a table struct for this dao
func (dao *Dao) SetDefaultModel(tb interface{}, deftultDataKey string) (err error) {
	structType := reflect.TypeOf(tb).Elem()
	if db == nil || cfg == nil {
		err = errors.New("Do NewConfig() and NewDb() first !!")
	}

	dao.DaoStruct = structType.Name()
	dao.DaoStructType = structType
	dao.DataKey = deftultDataKey

	return
}

// GetConfig : return redis.Options
func (dao *Dao) GetConfig() *redis.Options {
	return getConfig()
}

// SetConfig : set config by user, pw, addr, db
func (dao *Dao) SetConfig(addr, pw string, db int) *Dao {
	setConfig(addr, pw, db)
	return dao
}

// SetOriginConfig : set config by redis.Options
func (dao *Dao) SetOriginConfig(c *redis.Options) *Dao {
	setOriginConfig(c)
	return dao
}

// OpenDB : connect to db
func (dao *Dao) OpenDB() *Dao {
	if _, err := openDB(); err != nil {
		panic("cannot connect to db")
	}
	dao.Client = db
	dao.DbNnum = getConfig().DB

	return dao
}

// OpenDBWithPoolConns : connect to db and set pool conns
func (dao *Dao) OpenDBWithPoolConns(active, idle int) *Dao {
	if _, err := openDBWithPoolConns(active, idle); err != nil {
		panic("cannot connect to db")
	}
	return dao

}
