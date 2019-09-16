module github.com/eehsiao/go-models

go 1.12

require (
	github.com/eehsiao/go-models/mysql v0.0.0-20190916083412-99e09061e334
	github.com/eehsiao/go-models/redis v0.0.0-20190916083412-99e09061e334
	github.com/eehsiao/go-models/sqlbuilder v0.0.0-20190916083412-99e09061e334
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	google.golang.org/appengine v1.6.2 // indirect
)

replace github.com/eehsiao/go-models/mysql => ../go-models/mysql

replace github.com/eehsiao/go-models/redis => ../go-models/redis

replace github.com/eehsiao/go-models/sqlbuilder => ../go-models/sqlbuilder
