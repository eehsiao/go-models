// Author :		Eric<eehsiao@gmail.com>

package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Dao : the data access object struct
type Dao struct {
	*sql.DB
	Tb interface{}
}

// Tx : the transaction of Dao
type Tx struct {
	*sql.Tx
}

// newDao : create a new Dao and open mysql
func newDao(config *mysql.Config) (dao *Dao, err error) {
	var db *sql.DB
	if db, err = sql.Open("mysql", config.FormatDSN()); err != nil {
		return nil, err
	}
	dao = &Dao{DB: db}

	return
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
