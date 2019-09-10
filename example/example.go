// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"fmt"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
	"github.com/eehsiao/go-models/redis"
)

var (
	myDao     *mysql.Dao
	redDao    *redis.Dao
	users     []User
	user      User
	serialStr string
	keyValues = make(map[string]interface{})
	status    string
	err       error
	redBool   bool
)

func main() {
	myUserDao := &MyUserDao{
		Dao: mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB(),
	}

	if err = myUserDao.RegisterModel((*UserTb)(nil), "user"); err != nil {
		panic(err.Error())
	}

	if users, err = myUserDao.GetUsers(); len(users) > 0 {
		redUserModel := &RedUserModel{
			Dao: redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB(),
		}

		if err = redUserModel.RegisterModel((*User)(nil), "user"); err != nil {
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
				fmt.Println(fmt.Sprintf("%s : %v", k, user))
			}
		}
	}

}
