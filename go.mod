module github.com/eehsiao/go-models

go 1.12

require (
	github.com/eehsiao/go-models/lib v0.0.1
	github.com/eehsiao/go-models/mysql v0.0.1
	github.com/go-redis/redis v6.15.2+incompatible // indirect
	github.com/go-sql-driver/mysql v1.4.1
)

replace github.com/eehsiao/go-models/mysql => ../go-models/mysql

replace github.com/eehsiao/go-models/lib => ../go-models/lib
