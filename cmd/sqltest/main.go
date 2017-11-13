package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"sqltest"
)

type sqlite3Dialect struct{}

func (_ sqlite3Dialect) DriverName() string {
	return "sqlite3"
}

func (_ sqlite3Dialect) ColumnType(typ string, arg []int) string {
	if len(arg) > 0 {
		return fmt.Sprintf("%s(%d)", typ, arg[0])
	}
	return typ
}

type postgresDialect struct{}

func (_ postgresDialect) DriverName() string {
	return "postgres"
}

func (_ postgresDialect) ColumnType(typ string, arg []int) string {
	typ = strings.ToUpper(typ)
	if typ == "BINARY" || typ == "VARBINARY" || typ == "BLOB" {
		return "BYTEA"
	}
	if typ == "TEXT" {
		return "TEXT"
	}
	if len(arg) > 0 {
		return fmt.Sprintf("%s(%d)", typ, arg[0])
	}
	return typ
}

type mysqlDialect struct{}

func (_ mysqlDialect) DriverName() string {
	return "mysql"
}

func (_ mysqlDialect) ColumnType(typ string, arg []int) string {
	if len(arg) > 0 {
		return fmt.Sprintf("%s(%d)", typ, arg[0])
	}
	return typ
}

var (
	driver = flag.String("driver", "sqlite3", "database driver to use")
	source = flag.String("source", ":memory:", "data source to use")

	dialects = map[string]sqltest.Dialect{
		"sqlite3":  sqlite3Dialect{},
		"postgres": postgresDialect{},
		"mysql":    mysqlDialect{},
	}
)

type report struct{}

func (r report) Report(test string, err error) error {
	if err == nil {
		log.Printf("%s: passed\n", test)
	} else {
		log.Printf("%s: failed: %s\n", test, err)
	}
	return nil
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"testdata"}
	}

	dialect, ok := dialects[*driver]
	if !ok {
		log.Fatal(fmt.Errorf("missing dialect for driver %s", *driver))
	}

	for _, arg := range args {
		var run sqltest.DBRunner
		err := run.Connect(*driver, *source)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("testing %s\n", arg)
		err = sqltest.RunTests(arg, &run, report{}, dialect)
		if err != nil {
			log.Fatal(err)
		}
	}
}
