package sqltestdb

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBRunner struct {
	db *sqlx.DB
}

func (run *DBRunner) RunExec(tst *Test) (int64, error) {
	result, err := run.db.Exec(tst.Test)
	if err == nil && result != nil && (tst.Statement == "INSERT" || tst.Statement == "UPDATE" ||
		tst.Statement == "DELETE" || tst.Statement == "COPY") {

		n, err := result.RowsAffected()
		if err == nil {
			return n, nil
		}
	}
	return -1, err
}

func (run *DBRunner) RunQuery(tst *Test) (QueryResult, error) {
	rows, err := run.db.Query(tst.Test)
	if err != nil {
		return QueryResult{}, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return QueryResult{}, err
	}

	lenCols := len(cols)
	dest := make([]interface{}, lenCols)
	for i := 0; i < lenCols; i++ {
		dest[i] = new(sql.RawBytes)
	}

	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return QueryResult{}, err
	}

	var types []string
	for _, ct := range colTypes {
		types = append(types, ct.DatabaseTypeName())
	}

	var results [][]string
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return QueryResult{}, err
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

	return QueryResult{
		Columns:   cols,
		TypeNames: types,
		Rows:      results,
	}, nil
}

func (dbr *DBRunner) Connect(driver, source string) error {
	db, err := sqlx.Connect(driver, source)
	if err != nil {
		return err
	}
	dbr.db = db
	return nil
}
