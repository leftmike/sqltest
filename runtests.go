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
)

type Runner interface {
	RunExec(tst *Test) error
	RunQuery(tst *Test) ([]string, [][]string, error)
}

type Reporter interface {
	Report(test string, err error) error
}

// RunTests runs all of the tests in a directory: <dir>/sql/*.sql contains the tests and
// <dir>/expected/*.out contains the expected output for each sql file. The output from each
// sql file will get written to <dir>/output/*.out.
func RunTests(dir string, run Runner, report Reporter) error {
	files, err := filepath.Glob(filepath.Join(dir, "sql", "*.sql"))
	if err != nil {
		return fmt.Errorf("Glob(%q) failed with %s", filepath.Join(dir, "sql", "*.sql"), err)
	}

	for _, sqlname := range files {
		err, ret := testFile(dir, sqlname, run)
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

func testFile(dir, sqlname string, run Runner) (error, error) {
	// Get filename without .sql
	basename := filepath.Base(sqlname)
	basename = basename[:strings.LastIndexByte(basename, '.')]

	sqlf, err := os.Open(sqlname)
	if err != nil {
		return fmt.Errorf("Open(%q) failed with %s", sqlname, err), nil
	}
	defer sqlf.Close()

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

		for _, l := range tst.Comments {
			fmt.Fprintln(&out, l)
		}
		for _, l := range tst.Stmts {
			fmt.Fprintln(&out, l)
		}

		if tst.IsQuery {
			err = testQuery(tst, run, &out)

		} else {
			err = run.RunExec(tst)
		}
		if err != nil {
			return nil, fmt.Errorf("%s:%d: %s", tst.Filename, tst.LineNumber, err)
		}
	}

	outname := filepath.Join(dir, "output", basename+".out")
	err = ioutil.WriteFile(outname, out.Bytes(), 0666)
	if err != nil {
		return fmt.Errorf("WriteFile(%q) failed with %s", outname, err), nil
	}

	expname := filepath.Join(dir, "expected", basename+".out")
	exp, _ := ioutil.ReadFile(expname) // Ignore the error; exp will be nil in that case.

	if exp == nil {
		return nil, fmt.Errorf("no expected output for %s", sqlname)
	}

	if string(exp) != out.String() {
		return nil, fmt.Errorf("%s and %s are different\n", outname, expname)
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

func testQuery(tst *Test, run Runner, out io.Writer) error {
	cols, rows, err := run.RunQuery(tst)
	if err != nil {
		return err
	}
	sort.Sort(resultRows(rows))

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
