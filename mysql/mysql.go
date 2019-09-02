// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"github.com/go-sql-driver/mysql"
)

const (
	defaultOpenConns = 5
	defaultIdleConns = 1
)

// NewConfig : new mysql config via go-models
func NewConfig(user, pw, addr, db string) (c *mysql.Config) {
	c = mysql.NewConfig()
	c.User = user
	c.Passwd = pw
	c.Addr = addr
	c.DBName = db

	return
}

// NewMysql : new mysql Dao
func NewMysql(config *mysql.Config) (dao *Dao, err error) {
	if dao, err = newDao(config); err != nil {
		return nil, err
	}

	dao.Db.SetMaxOpenConns(defaultOpenConns)
	dao.Db.SetMaxIdleConns(defaultIdleConns)
	return
}

// NewMysqlWithPoolConns : new mysql Dao with pool conns
func NewMysqlWithPoolConns(config *mysql.Config, active, idle int) (dao *Dao, err error) {

	if dao, err = newDao(config); err != nil {
		return nil, err
	}

	dao.Db.SetMaxOpenConns(active)
	dao.Db.SetMaxIdleConns(idle)

	return
}
