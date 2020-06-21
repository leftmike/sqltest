package main

import (
	"testing"

	"github.com/leftmike/sqltest/sqltestdb"
)

type report_test struct {
	driver string
	t      *testing.T
}

func (r report_test) Report(test string, err error) error {
	if err == nil {
		r.t.Logf("%s: %s: passed", r.driver, test)
	} else if err == sqltestdb.Skipped {
		r.t.Logf("%s: %s: skipped", r.driver, test)
	} else {
		r.t.Errorf("%s: %s: failed: %s", r.driver, test, err)
	}
	return nil
}

func TestSQL(t *testing.T) {
	d := Drivers["sqlite3"]
	err := d.RunTests(report_test{d.Driver, t})
	if err != nil {
		t.Error(err)
	}
}
