package sqltestdb_test

import (
	"testing"

	"github.com/leftmike/sqltest/sqltestdb"
)

type testRunner struct{}

func (run testRunner) RunExec(tst *sqltestdb.Test) error {
	return nil
}

func (run testRunner) RunQuery(tst *sqltestdb.Test) ([]string, [][]string, error) {
	return []string{"col1", "col2", "col3"},
		[][]string{{"1", "2", "3"}, {"4", "5", "6"}}, nil
}

type report struct {
	test string
	err  error
}

type testReporter []report

func (tr *testReporter) Report(test string, err error) error {
	*tr = append(*tr, report{test, err})
	return nil
}

type testDialect struct {
	sqltestdb.DefaultDialect
}

func (_ testDialect) DriverName() string {
	return "test"
}

func TestRunTests(t *testing.T) {
	var tr testReporter
	err := sqltestdb.RunTests("testdata", testRunner{}, &tr, testDialect{}, false)
	if err != nil {
		t.Errorf("RunTests() failed with %s", err)
	}
	for _, r := range tr {
		if r.err != nil {
			t.Errorf("%s: %s", r.test, r.err)
		}
	}
}
