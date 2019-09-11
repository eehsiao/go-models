// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/sqlbuilder"
	"github.com/go-sql-driver/mysql"
)

// Dao : the data access object struct
type Dao struct {
	*sql.DB
	*sqlbuilder.SQLBuilder
	DaoStruct     string
	DaoStructType reflect.Type
	DbName        string
	TbName        string
}

// NewDao : create a new empty Dao
func NewDao() *Dao {
	return &Dao{
		SQLBuilder: sqlbuilder.NewSQLBuilder(),
	}
}

// SetDefaultModel : set the struct for this dao default
func (dao *Dao) SetDefaultModel(tb interface{}, deftultTbName string) (err error) {
	structType := reflect.TypeOf(tb).Elem()
	if db == nil || cfg == nil {
		err = errors.New("Do NewConfig() and NewDb() first !!")
	}

	dao.DaoStruct = structType.Name()
	dao.DaoStructType = structType
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
	dao.DB = db
	dao.DbName = getConfig().DBName

	return dao
}

// OpenDBWithPoolConns : connect to db and set pool conns
func (dao *Dao) OpenDBWithPoolConns(active, idle int) *Dao {
	if _, err := openDBWithPoolConns(active, idle); err != nil {
		panic("cannot connect to db")
	}
	return dao

}

func (dao *Dao) ScanType(rows *sql.Rows, tb interface{}) (t []interface{}, err error) {
	for rows.Next() {
		gTb := reflect.New(reflect.TypeOf(tb).Elem()).Interface()
		if err = rows.Scan(lib.Struct4Scan(gTb)...); err == nil {
			t = append(t, gTb)
		}
	}

	return
}

func (dao *Dao) Scan(rows *sql.Rows) (t []interface{}, err error) {
	for rows.Next() {
		gTb := reflect.New(dao.DaoStructType).Interface()
		if err = rows.Scan(lib.Struct4Scan(gTb)...); err == nil {
			t = append(t, gTb)
		}
	}

	return
}

func (dao *Dao) ScanRowType(row *sql.Row, tb interface{}) (t interface{}, err error) {
	t = reflect.New(reflect.TypeOf(tb).Elem()).Interface()
	err = row.Scan(lib.Struct4Scan(t)...)

	return
}

func (dao *Dao) ScanRow(row *sql.Row) (t interface{}, err error) {
	t = reflect.New(dao.DaoStructType).Interface()
	err = row.Scan(lib.Struct4Scan(t)...)

	return
}

func (dao *Dao) Get() (rows *sql.Rows, err error) {
	if !dao.CanBuildSelect() {
		return nil, errors.New("cannot select")
	}
	rows, err = dao.Query(dao.BuildSelectSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}

func (dao *Dao) GetRow() (row *sql.Row, err error) {
	if !dao.CanBuildSelect() {
		return nil, errors.New("cannot select")
	}

	row = dao.QueryRow(dao.BuildSelectSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}
