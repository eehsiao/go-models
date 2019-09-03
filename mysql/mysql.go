// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"github.com/go-sql-driver/mysql"
)

// NewConfig : new mysql config via go-models
// addr : can with port number. ex: '127.0.0.1:3306', if u want to use default port, just use ip addr ex: '127.0.0.1'
func NewConfig(user, pw, addr, db string) (c *mysql.Config) {
	c = mysql.NewConfig()
	c.User = user
	c.Passwd = pw
	c.Addr = addr
	c.DBName = db

	return
}

// NewMysql : create a new mysql Dao
func NewMysql(config *mysql.Config) (dao *Dao, err error) {
	if dao, err = newDao(config); err != nil {
		return nil, err
	}

	// setting connections pool by default
	dao.SetMaxOpenConns(defaultOpenConns)
	dao.SetMaxIdleConns(defaultIdleConns)
	return
}

// NewMysqlWithPoolConns : create a new mysql Dao with pool conns
func NewMysqlWithPoolConns(config *mysql.Config, active, idle int) (dao *Dao, err error) {

	if dao, err = newDao(config); err != nil {
		return nil, err
	}

	// setting connections pool
	dao.SetMaxOpenConns(active)
	dao.SetMaxIdleConns(idle)

	return
}
