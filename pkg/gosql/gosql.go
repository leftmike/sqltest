/*
Package godoc use the database drivers for Go to drive sqltest.
*/
package gosql

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/leftmike/sqltest/pkg/sqltest"
)

type sqlite3Dialect struct {
	sqltest.DefaultDialect
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
	sqltest.DefaultDialect
}

func (_ mysqlDialect) DriverName() string {
	return "mysql"
}

var (
	update   = flag.Bool("update", false, "update expected to output")
	testData = flag.String("testdata", "testdata", "directory of testdata")

	sqlite3Source  = flag.String("sqlite3", ":memory:", "data source to use for sqlite3")
	postgresSource = flag.String("postgres", "", "data source to use for postgres")
	mysqlSource    = flag.String("mysql", "", "data source to use for mysql")
)

type driver struct {
	Driver  string
	dialect sqltest.Dialect
	source  *string
}

var Drivers = map[string]driver{
	"sqlite3":  {"sqlite3", sqlite3Dialect{}, sqlite3Source},
	"postgres": {"postgres", postgresDialect{}, postgresSource},
	"mysql":    {"mysql", mysqlDialect{}, mysqlSource},
}

func (d driver) RunTests(r sqltest.Reporter) error {
	if *d.source == "" {
		return fmt.Errorf("no source for driver %s", d.Driver)
	}

	var run sqltest.DBRunner
	err := run.Connect(d.Driver, *d.source)
	if err != nil {
		return err
	}
	return sqltest.RunTests(*testData, &run, r, d.dialect, *update)
}
