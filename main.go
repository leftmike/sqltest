package main

import (
	"flag"
	"log"

	"github.com/leftmike/sqltest/sqltestdb"
)

var (
	useRDS    = flag.Bool("rds", false, "use an AWS RDS database")
	useAurora = flag.Bool("aurora", false, "use an AWS Aurora database")
)

type report struct {
	driver string
}

func (r report) Report(test string, err error) error {
	if err == nil {
		log.Printf("%s: %s: passed\n", r.driver, test)
	} else if err == sqltestdb.Skipped {
		log.Printf("%s: %s: skipped\n", r.driver, test)
	} else {
		log.Printf("%s: %s: failed: %s\n", r.driver, test, err)
	}
	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		log.Printf("testing %s\n", arg)
		d, ok := Drivers[arg]
		if !ok {
			log.Printf("invalid driver: %s\n", arg)
			continue
		}
		if d.Driver == "postgres" && (*useRDS || *useAurora) {
			s, err := EnsurePostgresAWS("sqltest-postgresql", *useAurora)
			if err != nil {
				log.Printf("error: %s: %s\n", d.Driver, err)
				continue
			}
			*d.Source = s
			log.Println(s)
		}
		err := d.RunTests(report{d.Driver})
		if err != nil {
			log.Printf("error: %s: %s\n", d.Driver, err)
		}
	}
}
