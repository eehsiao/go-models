# go-models
`go-models` its lite and easy model.

---------------------------------------
  * [Features](#features)
  * [Requirements](#requirements)
  * [Go-Module](#go-module)
  * [Docker](#docker)
  * [Usage](#usage)
    * [Lib](#lib)
    * [SqlBuilder](#sqlbuilder)
    * [Example](#example)
    * [How-to](#how-to)
        * [MySQL](#mysql)
        * [Redis](#redis)


## Features
    That is querybuilder with data object models for SQLs.
    And easy way to build your data logical layer for access redis.
    This is a easy way to access data from database. That you focus on data processing logical.
    Now support MySQL, MariaDB, Redis

    TODO support: PostgreSQL, MSSQL, SQLite Mongodb, ...

## Requirements
    * Go 1.12 or higher.
    * [database/sql](https://golang.org/pkg/database/sql/) package
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) package
    * [go-redis/redis](https://github.com/go-redis/redis) package

## Go-Module
create `go.mod` file in your package folder, and fill below
```
module github.com/eehsiao/go-models-example

go 1.13

require (
	github.com/eehsiao/go-models/lib latest
	github.com/eehsiao/go-models/mysql latest
	github.com/eehsiao/go-models/redis latest
	github.com/eehsiao/go-models/sqlbuilder latest
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
)

```

## Docker
Easy to start the test evn. That you can run the example code.
```bash
$ docker-compose up -d
```

## Usage
```go
import (
    "database/sql"
	"fmt"

	"github.com/eehsiao/go-models/lib"
	"github.com/eehsiao/go-models/mysql"
	"github.com/eehsiao/go-models/redis"
)

// UserTb : sql table struct that to store into mysql
type UserTb struct {
	Host       sql.NullString `TbField:"Host"`
	User       sql.NullString `TbField:"User"`
	SelectPriv sql.NullString `TbField:"Select_priv"`
}

// User : json struct that to store into redis
type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	SelectPriv string `json:"select_priv"`
}

//new mysql dao
myUserDao := &MyUserDao{
    Dao: mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB(),
}

// example 1 : directly use the sqlbuilder
myUserDao.Select("Host", "User", "Select_priv").From("user").Where("User='root'").Limit(1)
fmt.Println("sqlbuilder", myUserDao.BuildSelectSQL())
if row, err = myUserDao.GetRow(); err == nil {
    if val, err = myUserDao.ScanRowType(row, (*UserTb)(nil)); err == nil {
        u, _ := val.(*UserTb)
        fmt.Println("UserTb", u)
    }
}
    
// set a struct for dao as default model (option)
// (*UserTb)(nil) : nil pointer of the UserTb struct
// "user" : is real table name in the db
myUserDao.SetDefaultModel((*UserTb)(nil), "user")

// call model's Get() , get all rows in user table
// return (rows *sql.Rows, err error)
rows, err = myDao.Get()

// call model's GetRow() , get first row in user rows
// return (row *sql.Row, err error)
row, err = myDao.GetRow()


//new redis dao
redUserModel := &RedUserModel{
    Dao: redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB(),
}

// set a struct for dao as default model (option)
// (*User)(nil) : nil pointer of the User struct
// "user" : is real table name in the db
SetDefaultModel((*User)(nil), "user")
```
### Lib
    lib.Iif : is a inline IF-ELSE logic
    lib.Struct4Scan : transfer a object struct to poiter slces, that easy to scan the sql results.
    lib.Struce4Query : transfer a struct to a string for sql select fields. ex "idx, name".
    Struce4QuerySlice : transfer a struct to a []string slice.
    lib.Serialize : serialize a object to a json string.

### SqlBuilder
sqlbuilder its recursive call function, that you can easy to build sql string

ex: dao.Select().From().Join().Where().Limit()
#### SqlBuilder functions
* build select :
    * Select(s ...string)
    * Distinct(b bool)
    * Top(i int)
    * From(s ...string)
    * Where(s string)
    * WhereAnd(s ...string)
    * WhereOr(s ...string)
    * Join(s string, c string)
    * InnerJoin(s string, c string)
    * LeftJoin(s string, c string)
    * RightJoin(s string, c string)
    * FullJoin(s string, c string)
    * GroupBy(s ...string)
    * OrderBy(s ...string)
    * OrderByAsc(s ...string)
    * OrderByDesc(s ...string)
    * Having(s string)
    * BuildSelectSQL()
* build update :
    * Set(s map[string]interface{})
    * FromOne(s string)
    * BuildUpdateSQL()
* build insert : 
    * Into(s string)
    * Fields(s ...string)
    * Values(s ...[]interface{})
    * BuildInsertSQL()
* build delete :
    * BuildDeleteSQL()
* common :
    * ClearBuilder()
    * BuildedSQL()
    * SetDbName(s string)
    * SetTbName(s string)
    * SwitchPanicToErrorLog(b bool)
    * PanicOrErrorLog(s string)

## Example
### 1 build-in
[example.go](https://github.com/eehsiao/go-models/blob/master/example/example.go)

The example will connect to local mysql and get user data.
Then connect to local redis and set user data, and get back.

### 2 example
`https://github.com/eehsiao/go-models-example/`


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
use Struce4QuerySlice to gen the sqlbuilder select fields
```go
m := mysql.NewDao().SetConfig("root", "mYaDmin", "127.0.0.1:3306", "mysql").OpenDB()
m.Select(lib.Struce4QuerySlice(m.DaoStructType)...).From(m.TbName).Limit(3)
```
#### 3.
scan the sql result to the struct of object
```go
row, err = m.GetRow()
if val, err = m.ScanRowType(row, (*UserTb)(nil)); err == nil {
    u, _ := val.(*UserTb)
    fmt.Println("UserTb", u)
}
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

#### 2.
create redis dao
```go
m := redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB()
```

#### 3.
directly use the go-redis command function
```go
redBool, err = m.HSet(userTable, redKey, serialStr).Result()
```
