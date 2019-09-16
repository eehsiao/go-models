// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"database/sql"
	"fmt"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
	"github.com/eehsiao/go-models/redis"
)

var (
	myDao     *mysql.Dao
	redDao    *redis.Dao
	users     []*User
	user      *User
	serialStr string
	keyValues = make(map[string]interface{})
	status    string
	err       error
	redBool   bool
	val       interface{}
	row       *sql.Row
)

func main() {
	myUserDao := &MyUserDao{
		Dao: mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB(),
	}

	// example 1 : use sql builder
	sets := map[string]interface{}{"foo": 1, "bar": "2", "test": true}
	myUserDao.Set(sets).From("user").Where("abc=1")
	fmt.Println("sqlbuilder", myUserDao.BuildUpdateSQL().BuildedSQL())

	// example 2 : directly use the sqlbuilder
	myUserDao.Select("Host", "User", "Select_priv").From("user").Where("User='root'").Limit(1)
	if row, err = myUserDao.GetRow(); err == nil {
		if val, err = myUserDao.ScanRowType(row, (*UserTb)(nil)); err == nil {
			u, _ := val.(*UserTb)
			fmt.Println("UserTb", u)
		}
	}

	// example 3 : use the data logical
	// set a struct for dao as default model (option)
	// (*UserTb)(nil) : nil pointer of the UserTb struct
	// "user" : is real table name in the db
	if err = myUserDao.SetDefaultModel((*UserTb)(nil), "user"); err != nil {
		panic(err.Error())
	}

	if user, err = myUserDao.GetFirstUser(); err == nil {
		fmt.Println("GetFirstUser", user)
	}

	if users, err = myUserDao.GetUsers(); len(users) > 0 {
		fmt.Println("GetUsers", users)

		redUserModel := &RedUserModel{
			Dao: redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB(),
		}

		if err = redUserModel.SetDefaultModel((*User)(nil), "user"); err != nil {
			panic(err.Error())
		}

		for _, u := range users {
			if serialStr, err = lib.Serialize(u); err == nil {
				redKey := u.Host + u.User
				keyValues[redKey] = serialStr
				// HSet is github.com/go-redis/redis original command
				if redBool, err = redUserModel.HSet(userTable, redKey, serialStr).Result(); err != nil {
					panic(err.Error())
				}
			}
		}
		// UserHMSet is a data logical function
		// its a multiple Set to call HMSet, write in redUserDL data logical
		if status, err = redUserModel.UserHMSet(keyValues); err != nil {
			panic(err.Error())
		}

		for k, _ := range keyValues {
			// UserHGet is a data logical function
			// its a multiple HGet to call HMSet, write in redUserDL data logical
			if user, err = redUserModel.UserHGet(k); err == nil {
				fmt.Println(fmt.Sprintf("UserHGet : %s = %v", k, user))
			}
		}
	}

}
