package sqltest

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBRunner struct {
	db *sqlx.DB
}

func (run *DBRunner) RunExec(tst *Test) error {
	_, err := run.db.Exec(tst.Test)
	return err
}

func (run *DBRunner) RunQuery(tst *Test) ([]string, [][]string, error) {
	rows, err := run.db.Query(tst.Test)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	lenCols := len(cols)
	dest := make([]interface{}, lenCols)
	for i := 0; i < lenCols; i++ {
		dest[i] = new(sql.RawBytes)
	}

	var results [][]string
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, nil, err
		}
		var resultRow []string
		for _, v := range dest {
			if rb, ok := v.(*sql.RawBytes); ok {
				resultRow = append(resultRow, string(*rb))
			} else {
				resultRow = append(resultRow, "-?-")
			}
		}
		results = append(results, resultRow)
	}

	return cols, results, nil
}

func (dbr *DBRunner) Connect(driver, source string) error {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		return err
	}
	dbr.db = db
	return nil
}
