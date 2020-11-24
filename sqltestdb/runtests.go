package sqltestdb

import (
	"bufio"
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

type QueryResult struct {
	Columns   []string
	TypeNames []string
	Rows      [][]string
}

type Runner interface {
	// RunExec is used to execute a single SQL statement which must not return rows.
	RunExec(tst *Test) (int64, error)
	// RunQuery is used to execute a single SQL statement which returns a slice of column
	// names and a slice of rows.
	RunQuery(tst *Test) (QueryResult, error)
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
func RunTests(dir string, run Runner, report Reporter, dialect Dialect, update, psql bool) error {
	files, err := filepath.Glob(filepath.Join(dir, "sql", "*.sql"))
	if err != nil {
		return fmt.Errorf("Glob(%q) failed with %s", filepath.Join(dir, "sql", "*.sql"), err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no sql files found in %s", dir)
	}

	for _, sqlname := range files {
		err, ret := testFile(dir, sqlname, run, dialect, update, psql)
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

func testFile(dir, sqlname string, run Runner, dialect Dialect, update, psql bool) (error, error) {
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
		var tst *Test
		tst, err = scanner.Scan()
		if err != nil {
			return err, nil
		}
		if tst == nil {
			break
		}

		if !psql {
			fmt.Fprintln(&out, tst.Test)
		}

		var tctx TestContext
		tst.Test, tctx, err = tmplCtx.Execute(tst.Test)
		if err != nil {
			return err, nil
		}
		if tmplCtx.Skip {
			return nil, Skipped
		}

		stmt := strings.ToUpper(tst.Statement)
		if stmt == "SELECT" || stmt == "SHOW" || stmt == "EXPLAIN" || stmt == "VALUES" {
			err = testQuery(tst, run, &out, &tctx, psql)
		} else {
			var n int64
			n, err = run.RunExec(tst)
			if psql && stmt != "" {
				if n >= 0 {
					if stmt == "INSERT" {
						fmt.Fprintf(&out, "INSERT 0 %d\n", n)
					} else {
						fmt.Fprintf(&out, "%s %d\n", stmt, n)
					}
				} else {
					fmt.Fprintln(&out, stmt)
				}
			}
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
		expString := readFile(expname)

		if expString == "" {
			return nil, fmt.Errorf("no expected output for %s", sqlname)
		}

		outString := out.String()
		if expString != outString {
			return nil, fmt.Errorf("%s and %s are different\n%v", outname, expname,
				diff.LineDiff(expString, outString))
		}
	}

	return nil, nil
}

func readFile(name string) string {
	f, err := os.Open(name)
	if err != nil {
		return ""
	}
	defer f.Close()

	var out bytes.Buffer
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Fprintln(&out, s.Text())
	}

	err = s.Err()
	if err != nil {
		return ""
	}

	return out.String()
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

func testQuery(tst *Test, run Runner, out io.Writer, tctx *TestContext, psql bool) error {
	qr, err := run.RunQuery(tst)
	if err != nil {
		return err
	}
	if !tctx.NoSort {
		sort.Sort(resultRows(qr.Rows))
	}

	if psql {
		// Unaligned output like psql -A
		psqlOutput(out, qr.Columns, qr.Rows)
	} else {
		tabwriterOutput(out, qr.Columns, qr.Rows)
	}
	switch len(qr.Rows) {
	case 0:
		fmt.Fprint(out, "(no rows)\n")
	case 1:
		fmt.Fprint(out, "(1 row)\n")
	default:
		fmt.Fprintf(out, "(%d rows)\n", len(qr.Rows))
	}
	return nil
}

func psqlOutput(out io.Writer, cols []string, rows [][]string) {
	for cdx, col := range cols {
		if cdx > 0 {
			fmt.Fprint(out, "|")
		}
		fmt.Fprint(out, col)
	}
	fmt.Fprintln(out)

	for _, row := range rows {
		for vdx, val := range row {
			if vdx > 0 {
				fmt.Fprint(out, "|")
			}
			fmt.Fprint(out, val)
		}
		fmt.Fprintln(out)
	}
}

func tabwriterOutput(out io.Writer, cols []string, rows [][]string) {
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
}
