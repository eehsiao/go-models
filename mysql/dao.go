// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/eehsiao/go-models/lib"
	"github.com/go-sql-driver/mysql"
)

// Dao : the data access object struct
type Dao struct {
	*sql.DB
	DaoStruct     string
	DaoStructType reflect.Type
	DbName        string
	TbName        string
}

// NewDao : create a new empty Dao
func NewDao() *Dao {
	return &Dao{}
}

// RegisterModel : register a table struct for this dao
func (dao *Dao) RegisterModel(tb interface{}, deftultTbName string) (err error) {
	structType := reflect.TypeOf(tb).Elem()
	if db == nil || cfg == nil {
		err = errors.New("Do NewConfig() and NewDb() first !!")
	}

	dao.DB = db
	dao.DaoStruct = structType.Name()
	dao.DaoStructType = structType
	dao.DbName = getConfig().DBName
	dao.TbName = deftultTbName

	return
}

// GetConfig : return mysql.Config
func (dao *Dao) GetConfig() *mysql.Config {
	return getConfig()
}

// SetConfig : set config by user, pw, addr, db
func (dao *Dao) SetConfig(user, pw, addr, db string) *Dao {
	setConfig(user, pw, addr, db)
	return dao
}

// SetOriginConfig : set config by mysql.Config
func (dao *Dao) SetOriginConfig(c *mysql.Config) *Dao {
	setOriginConfig(c)
	return dao
}

// OpenDB : connect to db
func (dao *Dao) OpenDB() *Dao {
	if _, err := openDB(); err != nil {
		panic("cannot connect to db")
	}
	return dao
}

// OpenDBWithPoolConns : connect to db and set pool conns
func (dao *Dao) OpenDBWithPoolConns(active, idle int) *Dao {
	if _, err := openDBWithPoolConns(active, idle); err != nil {
		panic("cannot connect to db")
	}
	return dao

}

func (dao *Dao) GetAll() (t []interface{}, err error) {
	selSQL := "SELECT " + lib.Struce4Query(dao.DaoStructType)
	selSQL += " FROM " + dao.DbName + "." + dao.TbName

	var rows *sql.Rows
	if rows, err = dao.Query(selSQL); err == nil {
		for rows.Next() {
			gTb := reflect.New(dao.DaoStructType).Interface()
			if err = rows.Scan(lib.Struct4Scan(gTb)...); err == nil {
				t = append(t, gTb)
			}
		}
	}
	rows.Close()

	return
}

func (dao *Dao) GetRow() (t interface{}, err error) {
	selSQL := "SELECT " + lib.Struce4Query(dao.DaoStructType)
	selSQL += " FROM " + dao.DbName + "." + dao.TbName

	var rows *sql.Rows
	if rows, err = dao.Query(selSQL); err == nil {
		gTb := reflect.New(dao.DaoStructType).Interface()
		err = rows.Scan(lib.Struct4Scan(gTb)...)
		t = gTb
	}
	rows.Close()

	return
}
