/*
input:
-- stmt: exec / query
-- identical: true
-- skip: postgres
-- only: maho
*/
package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"sqltest"
)

var (
	driver = flag.String("driver", "sqlite3", "database driver to use")
	source = flag.String("source", ":memory:", "data source to use")
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

	for _, arg := range args {
		var run sqltest.DBRunner
		err := run.Connect(*driver, *source)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("testing %s\n", arg)
		err = sqltest.RunTests(arg, &run, report{})
		if err != nil {
			log.Fatal(err)
		}
	}
}
