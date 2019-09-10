// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
)

const (
	userTable = "user"
)

// MyUserDao : extend from mysql.Dao
type MyUserDao struct {
	*mysql.Dao
}

// UserTb : sql table struct that to store into mysql
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

// GetUsers : this is a data logical function, you can write more logical in there
// sample data logical function
func (m *MyUserDao) GetUsers() (users []User, err error) {
	var vals []interface{}
	if vals, err = m.GetAll(); err == nil {
		for _, v := range vals {
			// var u *UserTb
			u, _ := v.(*UserTb)

			user := User{
				Host:       lib.Iif(u.Host.Valid, u.Host.String, "").(string),
				User:       lib.Iif(u.User.Valid, u.User.String, "").(string),
				SelectPriv: lib.Iif(u.SelectPriv.Valid, u.SelectPriv.String, "").(string),
			}
			users = append(users, user)
		}
	}

	return
}
