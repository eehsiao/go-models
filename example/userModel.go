// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"
	"reflect"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
)

const (
	thisTable = "user"
)

type MyDao struct {
	*mysql.Dao
}

type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

func (myDao *MyDao) GetUsers() (users []UserTb, err error) {
	selSQL := "SELECT " + lib.Struce4Query(reflect.TypeOf(UserTb{}))
	selSQL += " FROM " + thisTable

	var rows *sql.Rows
	if rows, err = myDao.Query(selSQL); err == nil {
		for rows.Next() {
			user := UserTb{}
			if err = rows.Scan(lib.Struct4Scan(&user)...); err == nil {
				users = append(users, user)
			}
		}
	}

	return
}
