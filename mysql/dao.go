// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Dao : the data access object struct
type Dao struct {
	Db *sql.DB
}

// Tx : the transaction of Dao
type Tx struct {
	Tx *sql.Tx
}

// newDao : create a new Dao and open mysql
func newDao(config *mysql.Config) (dao *Dao, err error) {
	var db *sql.DB
	if db, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		return nil, err
	}
	dao = &Dao{Db: db}

	return
}

// Ping : Ping the Dao
func (dao *Dao) Ping() (err error) {
	return dao.Db.Ping()
}

// Close : Close the Dao
func (dao *Dao) Close() {
	dao.Db.Close()
}

// Exec : execute a sql query via Dao
func (dao *Dao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dao.Db.Exec(query, args...)
}

// Query : do a sql query via Dao
// return : sql.Rows
func (dao *Dao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dao.Db.Query(query, args...)
}

// QueryRow : do a sql query via Dao and get one return row
// return : sql.Row
func (dao *Dao) QueryRow(query string, args ...interface{}) *sql.Row {
	return dao.Db.QueryRow(query, args...)
}

// Begin : begin a Dao transaction
func (dao *Dao) Begin() (t *Tx, err error) {
	if tx, err := dao.Db.Begin(); err == nil {
		t = &Tx{Tx: tx}
	}

	return
}

// Commit : commit a Dao transaction
func (t *Tx) Commit() error {
	return t.Tx.Commit()
}

// Rollback : rollback a Dao transaction
func (t *Tx) Rollback() error {
	return t.Tx.Rollback()
}

// Exec : execute a sql query via Dao transaction
func (t *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.Tx.Exec(query, args...)
}

// Query : do a sql query via Dao transaction
// return : sql.Rows
func (t *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return t.Tx.Query(query, args...)
}

// QueryRow : do a sql query via Dao transaction and get one return row
// return : sql.Row
func (t *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return t.Tx.QueryRow(query, args...)
}

// GetLock : get a session lock via Dao transaction
func (t *Tx) GetLock(key string, secs int) (cnt int, err error) {
	err = t.Tx.QueryRow("SELECT COALESCE(GET_LOCK(?, ?), 0)", key, secs).Scan(&cnt)

	return
}

// ReleaseLock : release a session lock via Dao transaction
func (t *Tx) ReleaseLock(key string) (cnt int, err error) {
	err = t.Tx.QueryRow("SELECT COALESCE(RELEASE_LOCK(?), 0)", key).Scan(&cnt)

	return
}
