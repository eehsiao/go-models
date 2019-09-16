module github.com/eehsiao/go-models

go 1.12

require (
	github.com/eehsiao/go-models/lib v0.0.0-20190916032851-cbc471a88312
	github.com/eehsiao/go-models/mysql v0.0.0-20190916032851-cbc471a88312
	github.com/eehsiao/go-models/redis v0.0.0-20190916032851-cbc471a88312
	github.com/eehsiao/go-models/sqlbuilder v0.0.0-20190916032851-cbc471a88312
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
)

replace github.com/eehsiao/go-models/mysql => ../go-models/mysql

replace github.com/eehsiao/go-models/redis => ../go-models/redis

replace github.com/eehsiao/go-models/sqlbuilder => ../go-models/sqlbuilder

replace github.com/eehsiao/go-models/lib => ../go-models/lib
