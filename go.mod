module github.com/eehsiao/go-models

go 1.12

require (
	github.com/eehsiao/go-models/lib v0.0.0-latest
	github.com/eehsiao/go-models/mysql v0.0.0-latest
	github.com/eehsiao/go-models/redis v0.0.0-latest
	github.com/eehsiao/go-models/sqlbuilder v0.0.0-latest
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	google.golang.org/appengine v1.6.2 // indirect
)

replace github.com/eehsiao/go-models/lib => ./lib

replace github.com/eehsiao/go-models/mysql => ./mysql

replace github.com/eehsiao/go-models/redis => ./redis

replace github.com/eehsiao/go-models/sqlbuilder => ./sqlbuilder
