module github.com/eehsiao/go-models/mysql

go 1.12

require (
	github.com/eehsiao/go-models/lib v0.0.0-latest
	github.com/eehsiao/go-models/sqlbuilder v0.0.0-latest
	github.com/go-sql-driver/mysql v1.4.1
)
replace github.com/eehsiao/go-models/lib => ../lib
replace github.com/eehsiao/go-models/sqlbuilder => ../sqlbuilder
