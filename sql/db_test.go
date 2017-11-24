package sql_test

import (
	"flag"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"sqltest"
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

type report struct {
	driver string
	t      *testing.T
}

func (r report) Report(test string, err error) error {
	if err == nil {
		r.t.Logf("%s: %s: passed", r.driver, test)
	} else if err == sqltest.Skipped {
		r.t.Logf("%s: %s: skipped", r.driver, test)
	} else {
		r.t.Errorf("%s: %s: failed: %s", r.driver, test, err)
	}
	return nil
}

func TestSQL(t *testing.T) {
	drivers := []struct {
		driver  string
		dialect sqltest.Dialect
		source  *string
	}{
		{"sqlite3", sqlite3Dialect{}, sqlite3Source},
		{"postgres", postgresDialect{}, postgresSource},
		{"mysql", mysqlDialect{}, mysqlSource},
	}

	for _, d := range drivers {
		t.Run(d.driver, func(t *testing.T) {
			if *d.source == "" {
				t.Errorf("no source for driver %s", d.driver)
				return
			}

			var run sqltest.DBRunner
			err := run.Connect(d.driver, *d.source)
			if err != nil {
				t.Error(err)
				return
			}
			err = sqltest.RunTests(*testData, &run, report{d.driver, t}, d.dialect, *update)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
