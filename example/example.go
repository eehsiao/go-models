// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
)

var (
	stdLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	errLog = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
)

func main() {
	c := mysql.NewConfig("root", "mYaDmin", "127.0.0.1", "mysql")
	// You can modify config
	c.Net = "tcp"

	if m, err := mysql.NewMysql(c); err == nil {
		stdLog.Println(fmt.Sprint("m => ", m))
		user := &MyDao{
			Dao: m,
		}
		if users, err := user.GetUsers(); err == nil {
			if userSer, err := lib.Serialize(users); err == nil {
				stdLog.Println(userSer)
			}
		}

		user.Close()
	} else {
		errLog.Println("err", err)
	}

}
