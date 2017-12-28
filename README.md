# sqltest
Test SQL compatibility between different implementations.
* [SQLite](https://www.sqlite.org/)
* [PostgreSQL](https://www.postgresql.org/)
* [MySQL](https://www.sqlite.org/)
* [Maho](https://github.com/leftmike/maho)

## Running Tests

The tests are run using the Go
[testing](https://golang.org/pkg/testing/) package: `go test`.
Each test is a templatized file of SQL
statements run against a particular implementation and the output compared with the expected
output. These files are in `sql/testdata/sql`, `sql/testdata/output`, and
`sql/testdata/expected`.

The files in `testdata/*` are for testing the *implementation* of sqltest.

To control which implementation is run, use a `-run` flag; for example, to run just postgres,
use `-run=/postgres`.

Use the following flags to control the tests; they need to be proceeded by a single `-args`
argument:
* `-update`: update the expected output.
* `-testdata <directory>`: specify a different directory for testdata; the default is `testdata`.
* `-sqlite3`: data source to use for sqlite3; the default is `:memory:`.
* `-postgres`: data source to use for postgres; there is no default.
* `-mysql`: data source to use for mysql; there is no default.

I use dev/test RDS instances in AWS for testing. For example, to run against just postgres, use
the following command:
```
go test ./sql -v -run=/postgres -args -postgres "host=<host>.rds.amazonaws.com port=5432 dbname=<dbname> user=<user> password=<password>"
```

To update the expected output to be the output from postres, add `-update` after `-args`.
```
go test ./sql -v -run=/postgres -args -update -postgres ...
```

Finally, to run all of the tests against all of the supported implementations, do the following:
```
go test ./sql -args -postgres "host=<host>.rds.amazonaws.com port=5432 dbname=<dbname> user=<user> password=<password>" -mysql "<user>:<password>@tcp(<host>.rds.amazonaws.com)/<dbname>"
```

## Writing Tests

Template actions are delimited by `{{` and `}}` in the test files; see the Go
[template](https://golang.org/pkg/text/template/) package for more details.

* `{{Skip}}`: skip the rest of this test; this should go at the top of the file, typically
after testing for a specific implementation.
* `{{Dialect}}`: specifies which implementation is being tested.
* `{{Fail .Test|.Global [true|false]}}`: specify whether the next statement (`.Test`) or all
following statements (`.Global`) should succeed or fail. The initial default is that all
statements succeed.
* `{{Statement .Test|.Global <stmt>}}`: specify the kind of next statement (`.Test`) or all
following statements (`.Global`). The initial keyword of each SQL statement is used to
determine whether the statement is a SELECT or not. This is necessary to know whether or not to
expect a set of rows as the result. Use this action to override this.
* `{{Sort .Test|.Global true|false}}`: specify whether or not the set of rows from the next
statement (`.Test`) or all following statements (`.Global`) are sorted. The default is that they
are not sorted.
* `{{BINARY [<length>]}}`: the column type appropriate to the implementation.
* `{{VARBINARY [<length>]}}`: the column type appropriate to the implementation.
* `{{BLOB [<length>]}}`: the column type appropriate to the implementation.
* `{{TEXT [<length>]}}`: the column type appropriate to the implementation.

### Examples
Put the following at the top of a test file to skip the file if it is mysql that is being tested.The `--` at the start of the line is to keep it from being treated as SQL.
```
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}
```

Indicate that the next statement should fail.
```
{{Fail .Test}}
INSERT INTO tbl4 VALUES
    (1 = 'abc');
```

See `sql/testdata/sql/create.sql` and `sql/testdata/sql/insert_bool.sql` for more examples of
template usage.

## Adding SQL Implementations

See the godoc for this package.
