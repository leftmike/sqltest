package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"sqltest"
)

var (
	driver = flag.String("driver", "sqlite3", "database driver to use")
	source = flag.String("source", ":memory:", "data source to use")

	dialects = map[string]sqltest.Dialect{
		"sqlite3": sqltest.Dialect{
			Name: "sqlite3",
		},
		"postgres": sqltest.Dialect{
			Name: "postgres",
		},
		"mysql": sqltest.Dialect{
			Name: "mysql",
		},
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
