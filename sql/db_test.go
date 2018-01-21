package sql_test

import (
	"testing"

	"github.com/leftmike/sqltest/pkg/gosql"
	"github.com/leftmike/sqltest/pkg/sqltest"
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
	d := gosql.Drivers["sqlite3"]
	err := d.RunTests(report{d.Driver, t})
	if err != nil {
		t.Error(err)
	}
}
