package main

import (
	"flag"
	"log"

	"github.com/leftmike/sqltest/pkg/awsrds"
	"github.com/leftmike/sqltest/pkg/gosql"
	"github.com/leftmike/sqltest/pkg/sqltest"
)

var (
	useAWS = flag.Bool("aws", false, "use an AWS RDS database")
)

type report struct {
	driver string
}

func (r report) Report(test string, err error) error {
	if err == nil {
		log.Printf("%s: %s: passed\n", r.driver, test)
	} else if err == sqltest.Skipped {
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
		d, ok := gosql.Drivers[arg]
		if !ok {
			log.Printf("invalid driver: %s\n", arg)
			continue
		}
		if *d.Source == "" && *useAWS {
			s, err := awsrds.EnsurePostgresRDS("sqltest-postgresql")
			if err != nil {
				log.Printf("error: %s: %s\n", d.Driver, err)
				continue
			}
			*d.Source = s
		}
		err := d.RunTests(report{d.Driver})
		if err != nil {
			log.Printf("error: %s: %s\n", d.Driver, err)
		}
	}
}
