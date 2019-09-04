// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
	"github.com/eehsiao/go-models/redis"
)

var (
	stdLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	errLog = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	myDao     *mysql.Dao
	redDao    *redis.Dao
	users     []User
	user      User
	serialStr string
	rel       string
	status    string
	err       error
	redBool   bool
)

func main() {
	c := mysql.NewConfig("root", "mYaDmin", "127.0.0.1", "mysql")
	// You can modify config
	c.Net = "tcp"

	if myDao, err = mysql.NewMysql(c); err == nil {
		myUserDao := &MyUserDao{Dao: myDao}
		if users, err = myUserDao.GetUsers(); err == nil {
			if serialStr, err = lib.Serialize(users); err == nil {
				stdLog.Println(serialStr)
			}
		}
		myUserDao.Close()
	}

	if err != nil {
		panic(err.Error())
	}

	if len(users) > 0 {
		redUserModel := &RedUserModel{
			Dao: redis.NewRedis(
				redis.NewOptions("127.0.0.1:6379", "", 0),
			),
		}
		keyValues := make(map[string]interface{})
		for _, u := range users {
			if serialStr, err = lib.Serialize(u); err == nil {
				redKey := u.Host + u.User
				keyValues[redKey] = serialStr
				// HSet is github.com/go-redis/redis map to orgin redis command
				if redBool, err = redUserModel.HSet(userTable, redKey, serialStr).Result(); err != nil {
					panic(err.Error())
				}
			}
		}
		// UserHMSet is a data logical function , write by yourself
		if status, err = redUserModel.UserHMSet(userTable, keyValues); err != nil {
			panic(err.Error())
		}

		for k, _ := range keyValues {
			// HGet is github.com/go-redis/redis map to orgin redis command
			if rel, err = redUserModel.HGet(userTable, k).Result(); err == nil {
				if err = json.Unmarshal([]byte(rel), &user); err == nil {
					stdLog.Println(fmt.Sprintf("%s : %v", k, user))
				}
			}
		}
	}

}
