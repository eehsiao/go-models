// Author :		Eric<eehsiao@gmail.com>

package model

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/eehsiao/sqlbuilder"
)

// Dao : the data access object struct
type SqlDao struct {
	*sql.DB
	*sqlbuilder.SQLBuilder

	DaoStruct     string
	DaoStructType reflect.Type
}

// NewDao : create a new empty Dao
func NewSqlDao() *SqlDao {
	return &SqlDao{
		SQLBuilder: sqlbuilder.NewSQLBuilder(),
	}

}

func (dao *SqlDao) ScanType(rows *sql.Rows, tb interface{}) (t []interface{}, err error) {
	for rows.Next() {
		gTb := reflect.New(reflect.TypeOf(tb).Elem()).Interface()
		if err = rows.Scan(Struct4Scan(gTb)...); err == nil {
			t = append(t, gTb)
		}
	}

	return
}

func (dao *SqlDao) Scan(rows *sql.Rows) (t []interface{}, err error) {
	for rows.Next() {
		gTb := reflect.New(dao.DaoStructType).Interface()
		if err = rows.Scan(Struct4Scan(gTb)...); err == nil {
			t = append(t, gTb)
		}
	}

	return
}

func (dao *SqlDao) ScanRowType(row *sql.Row, tb interface{}) (t interface{}, err error) {
	t = reflect.New(reflect.TypeOf(tb).Elem()).Interface()
	err = row.Scan(Struct4Scan(t)...)

	return
}

func (dao *SqlDao) ScanRow(row *sql.Row) (t interface{}, err error) {
	t = reflect.New(dao.DaoStructType).Interface()
	err = row.Scan(Struct4Scan(t)...)

	return
}

func (dao *SqlDao) Get() (rows *sql.Rows, err error) {
	if !dao.IsHadBuildedSQL() {
		if !dao.IsHasSelects() {
			dao.Select(Struce4Query(dao.DaoStructType))
		}
		if !dao.CanBuildSelect() {
			return nil, errors.New("cannot select")
		}
		dao.BuildSelectSQL()
	}

	rows, err = dao.Query(dao.BuildedSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}

func (dao *SqlDao) GetRow() (row *sql.Row, err error) {
	if !dao.IsHadBuildedSQL() {
		if !dao.IsHasSelects() {
			dao.Select(Struce4Query(dao.DaoStructType))
		}
		if !dao.CanBuildSelect() {
			return nil, errors.New("cannot select")
		}
		dao.BuildSelectSQL()
	}

	row = dao.QueryRow(dao.BuildedSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}

func (dao *SqlDao) Update(s string) (r sql.Result, err error) {
	if s != "" {
		dao.FromOne(s)
	}
	if !dao.CanBuildUpdate() {
		return nil, errors.New("cannot update")
	}

	r, err = dao.Exec(dao.BuildedSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}

func (dao *SqlDao) Insert(s string) (r sql.Result, err error) {
	if s != "" {
		dao.Into(s)
	}
	if !dao.CanBuildInsert() {
		return nil, errors.New("cannot insert")
	}

	r, err = dao.Exec(dao.BuildedSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}

func (dao *SqlDao) Delete(s string) (r sql.Result, err error) {
	if s != "" {
		dao.FromOne(s)
	}
	if !dao.CanBuildDelete() {
		return nil, errors.New("cannot insert")
	}

	r, err = dao.Exec(dao.BuildedSQL())

	//reset sqlbuilder
	dao.ClearBuilder()

	return
}
