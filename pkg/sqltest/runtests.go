package sqltest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/andreyvit/diff"
)

var Skipped = fmt.Errorf("skipped")

type Runner interface {
	// RunExec is used to execute a single SQL statement which must not return rows.
	RunExec(tst *Test) error
	// RunQuery is used to execute a single SQL statement which returns a slice of column
	// names and a slice of rows.
	RunQuery(tst *Test) ([]string, [][]string, error)
}

type Reporter interface {
	// Report on the results of a single test.
	Report(test string, err error) error
}

// Dialect specifies the behavior of a particular SQL implementation.
type Dialect interface {
	// DriverName is the name of the dialect; eg. postgres.
	DriverName() string
	// ColumnType maps a particular type into something that the implementation understands.
	ColumnType(typ string) string
	// ColumnTypeArg maps a particular type with an argument into something that the
	// implementation understands.
	ColumnTypeArg(typ string, arg int) string
}

// DefaultDialect provides default behavior; most implementation specific dialects can be based
// on this one.
type DefaultDialect struct{}

func (_ DefaultDialect) ColumnType(typ string) string {
	return typ
}

func (_ DefaultDialect) ColumnTypeArg(typ string, arg int) string {
	return fmt.Sprintf("%s(%d)", typ, arg)
}

// RunTests runs all of the tests in a directory: <dir>/sql/*.sql contains the tests and
// <dir>/expected/*.out contains the expected output for each sql file. The output from each
// sql file will get written to <dir>/output/*.out.
func RunTests(dir string, run Runner, report Reporter, dialect Dialect, update bool) error {
	files, err := filepath.Glob(filepath.Join(dir, "sql", "*.sql"))
	if err != nil {
		return fmt.Errorf("Glob(%q) failed with %s", filepath.Join(dir, "sql", "*.sql"), err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no sql files found in %s", dir)
	}

	for _, sqlname := range files {
		err, ret := testFile(dir, sqlname, run, dialect, update)
		if err != nil {
			return err
		}
		err = report.Report(filepath.Base(sqlname), ret)
		if err != nil {
			return err
		}
	}

	return nil
}

func testFile(dir, sqlname string, run Runner, dialect Dialect, update bool) (error, error) {
	// Get filename without .sql
	basename := filepath.Base(sqlname)
	basename = basename[:strings.LastIndexByte(basename, '.')]

	sqlf, err := os.Open(sqlname)
	if err != nil {
		return fmt.Errorf("Open(%q) failed with %s", sqlname, err), nil
	}
	defer sqlf.Close()

	tmplCtx := NewTemplateContext(dialect)
	var out bytes.Buffer
	scanner := NewScanner(sqlf)
	scanner.Filename = basename + ".sql"
	for {
		tst, err := scanner.Scan()
		if err != nil {
			return err, nil
		}
		if tst == nil {
			break
		}

		fmt.Fprintln(&out, tst.Test)
		var tctx TestContext
		tst.Test, tctx, err = tmplCtx.Execute(tst.Test)
		if err != nil {
			return err, nil
		}
		if tmplCtx.Skip {
			return nil, Skipped
		}

		if strings.ToUpper(tst.Statement) == "SELECT" {
			err = testQuery(tst, run, &out, &tctx)
		} else {
			err = run.RunExec(tst)
		}
		if err != nil && !tctx.Fail {
			return nil, fmt.Errorf("%s:%d: %s", tst.Filename, tst.LineNumber, err)
		} else if err == nil && tctx.Fail {
			return nil, fmt.Errorf("%s:%d: did not fail", tst.Filename, tst.LineNumber)
		}
	}

	outname := filepath.Join(dir, "output", basename+".out")
	err = ioutil.WriteFile(outname, out.Bytes(), 0666)
	if err != nil {
		return fmt.Errorf("WriteFile(%q) failed with %s", outname, err), nil
	}

	expname := filepath.Join(dir, "expected", basename+".out")
	if update {
		err = ioutil.WriteFile(expname, out.Bytes(), 0666)
		if err != nil {
			return fmt.Errorf("WriteFile(%q) failed with %s", expname, err), nil
		}
	} else {
		exp, _ := ioutil.ReadFile(expname) // Ignore the error; exp will be nil in that case.

		if exp == nil {
			return nil, fmt.Errorf("no expected output for %s", sqlname)
		}

		expString := string(exp)
		outString := out.String()
		if expString != outString {
			return nil, fmt.Errorf("%s and %s are different\n%v", outname, expname,
				diff.LineDiff(expString, outString))
		}
	}

	return nil, nil
}

type resultRows [][]string

func (rr resultRows) Len() int {
	return len(rr)
}

func (rr resultRows) Swap(i, j int) {
	rr[i], rr[j] = rr[j], rr[i]
}

func (rr resultRows) Less(i, j int) bool {
	jrow := rr[j]
	for k, v := range rr[i] {
		switch strings.Compare(v, jrow[k]) {
		case -1:
			return true
		case 1:
			return false
		}
	}
	return false
}

func testQuery(tst *Test, run Runner, out io.Writer, tctx *TestContext) error {
	cols, rows, err := run.RunQuery(tst)
	if err != nil {
		return err
	}
	if !tctx.NoSort {
		sort.Sort(resultRows(rows))
	}

	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', tabwriter.AlignRight)

	fmt.Fprint(w, "\t")
	for _, col := range cols {
		fmt.Fprintf(w, "%s\t", col)
	}
	fmt.Fprint(w, "\n\t")
	for _, col := range cols {
		fmt.Fprintf(w, "%s\t", strings.Repeat("-", len(col)))

	}
	fmt.Fprintln(w)

	i := 1
	for _, row := range rows {
		fmt.Fprintf(w, "%d\t", i)
		for _, v := range row {
			fmt.Fprintf(w, "%s\t", v)
		}
		fmt.Fprintln(w)
		i += 1
	}
	w.Flush()
	switch len(rows) {
	case 0:
		fmt.Fprint(out, "(no rows)\n")
	case 1:
		fmt.Fprint(out, "(1 row)\n")
	default:
		fmt.Fprintf(out, "(%d rows)\n", len(rows))
	}
	return nil
}