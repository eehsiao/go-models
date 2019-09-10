# go-models
`go-models` its lite and easy model.

That is querybuilder with data object models for SQLs.
And easy way to build your data logical layer for access redis.


This is a easy way to access data from database. That you focus on data processing logical.
Now support MySQL, MariaDB, Redis
TODO: PostgreSQL, MSSQL, Mongodb, ...

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  * [Docker](#docker)
  * [Usage](#usage)
    * [Lib](#lib)
    * [Example](#example)
    * [How-to](#how-to)
        * [MySQL](#mysql)
        * [Redis](#redis)

## Features

## Requirements
    * Go 1.12 or higher.
    * [database/sql](https://golang.org/pkg/database/sql/) package
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) package
    * [go-redis/redis](https://github.com/go-redis/redis) package

## Installation
Easy install the package to your GOPATH from shell:
```bash
$ go get -u github.com/eehsiao/go-models
```

## Docker
Easy to start the test evn. That you can run the example code.
```bash
$ docker-compose up -d
```

## Usage
```go
import (
	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
	"github.com/eehsiao/go-models/redis"
)

//new mysql dao
myDao := mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB()
//register a struct for model
myDao.RegisterModel((*UserTb)(nil), "user")
// call model's GetAll() , get all rows in user table
users, err = myDao.GetAll()


//new redis dao
redDao := redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB()
//register a struct for model
RegisterModel((*User)(nil), "user")
```
### Lib
    lib.Iif : is a inline IF-ELSE logic
    lib.Struct4Scan : transfer a object struct to poiter slces, that easy to scan the sql results.
    lib.Struce4Query : transfer a struct to a string for sql select fields. ex "idx, name".
    lib.Serialize : serialize a object to a json string.

## Example
### 1 build-in
[example.go](https://github.com/eehsiao/go-models/blob/master/example/example.go)

The example will connect to local mysql and get user data.
Then connect to local redis and set user data, and get back.

### 2 real exam
`https://github.com/eehsiao/go-models-exam/`


## How-to 
How to design model data logical
### MySQL
#### 1.
create a table struct, and add the tag `TbField:"real table filed"`

`TbField` the tag is musted. `read table filed` also be same the table field.
```go
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}
```
#### 2.
use Struce4Query to gen the sql select fields
```go
	selSQL := "SELECT " + lib.Struce4Query(reflect.TypeOf(UserTb{}))
	selSQL += " FROM " + userTable
```
#### 3.
scan the sql result to the struct of object
```go
userTb := UserTb{}
err = rows.Scan(lib.Struct4Scan(&userTb)...)
```

### Redis
#### 1.
create a data struct, and add the tag `json:"name"`
```go
type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
    SelectPriv string `json:"select_priv"`
    IntVal     int    `json:"user,string"`
}
```

if you have integer value, you can add a transfer type desc.
such as json:"user,`string`"