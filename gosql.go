package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/leftmike/sqltest/sqltestdb"
)

// Use the database drivers for Go to drive sqltest.

type sqlite3Dialect struct {
	sqltestdb.DefaultDialect
}

func (_ sqlite3Dialect) DriverName() string {
	return "sqlite3"
}

type postgresDialect struct{}

func (_ postgresDialect) DriverName() string {
	return "postgres"
}

func (_ postgresDialect) ColumnType(typ string) string {
	if typ == "BINARY" || typ == "VARBINARY" || typ == "BLOB" {
		return "BYTEA"
	}
	return typ
}

func (_ postgresDialect) ColumnTypeArg(typ string, arg int) string {
	if typ == "BINARY" || typ == "VARBINARY" || typ == "BLOB" {
		return "BYTEA"
	}
	if typ == "TEXT" {
		return "TEXT"
	}
	return fmt.Sprintf("%s(%d)", typ, arg)
}

type mysqlDialect struct {
	sqltestdb.DefaultDialect
}

func (_ mysqlDialect) DriverName() string {
	return "mysql"
}

var (
	update   = flag.Bool("update", false, "update expected to output")
	testData = flag.String("testdata", "testdata", "directory of testdata")
	psql     = flag.Bool("psql", false, "output in psql format")

	sqlite3Source  = flag.String("sqlite3", ":memory:", "data source to use for sqlite3")
	postgresSource = flag.String("postgres",
		"host=localhost port=5432 dbname=test sslmode=disable",
		"data source to use for postgres (host=localhost port=5432 dbname=test sslmode=disable)")
	mysqlSource = flag.String("mysql", "", "data source to use for mysql")
)

type driver struct {
	Driver  string
	dialect sqltestdb.Dialect
	Source  *string
}

var Drivers = map[string]driver{
	"sqlite3":  {"sqlite3", sqlite3Dialect{}, sqlite3Source},
	"postgres": {"postgres", postgresDialect{}, postgresSource},
	"mysql":    {"mysql", mysqlDialect{}, mysqlSource},
}

func (d driver) RunTests(r sqltestdb.Reporter) error {
	if *d.Source == "" {
		return fmt.Errorf("no source for driver %s", d.Driver)
	}

	var run sqltestdb.DBRunner
	err := run.Connect(d.Driver, *d.Source)
	if err != nil {
		return err
	}
	return sqltestdb.RunTests(*testData, &run, r, d.dialect, *update, *psql)
}
