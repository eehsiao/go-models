// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"
	"reflect"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
)

const (
	userTable = "user"
)

type MyUserDao struct {
	*mysql.Dao
}

type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

// GetUsers : this is a data logical function, you can write more logical in there
// sample data logical function
func (m *MyUserDao) GetUsers() (users []User, err error) {
	selSQL := "SELECT " + lib.Struce4Query(reflect.TypeOf(UserTb{}))
	selSQL += " FROM " + userTable

	var rows *sql.Rows
	if rows, err = m.Query(selSQL); err == nil {
		for rows.Next() {
			userTb := UserTb{}
			if err = rows.Scan(lib.Struct4Scan(&userTb)...); err == nil {
				user := User{
					Host:       lib.Iif(userTb.Host.Valid, userTb.Host.String, "").(string),
					User:       lib.Iif(userTb.User.Valid, userTb.User.String, "").(string),
					SelectPriv: lib.Iif(userTb.SelectPriv.Valid, userTb.SelectPriv.String, "").(string),
				}

				users = append(users, user)
			}
		}
	}

	return
}
