# sqltest
Test SQL compatibility between different implementations.
* [SQLite](https://www.sqlite.org/)
* [PostgreSQL](https://www.postgresql.org/)
* [MySQL](https://www.sqlite.org/)
* [Maho](https://github.com/leftmike/maho)

## Running Tests

The tests are run using `sqltest` which is written in [Go](https://golang.org/).

Each test is a templatized file of SQL
statements run against a particular implementation and the output compared with the expected
output. These files are in `testdata/sql`, `testdata/output`, and
`testdata/expected`.

The files in `sqltestdb/testdata/*` are for testing the *implementation* of sqltest.

To control which implementation is run, specify it as an argument to `sqltest`.

Use the following flags to control the tests:
* `-update`: update the expected output.
* `-testdata <directory>`: specify a different directory for testdata; the default is `testdata`.
* `-sqlite3`: data source to use for sqlite3; the default is `:memory:`.
* `-postgres`: data source to use for postgres; there is no default.
* `-mysql`: data source to use for mysql; there is no default.
* `-aws`: use an AWS RDS instance for postgres; one will be started if necessary.

I use dev/test RDS instances in AWS for testing. For example, to run against just postgres, use
the following command:
```
./sqltest -postgres "host=<host>.rds.amazonaws.com port=5432 dbname=<dbname> user=<user> password=<password>" postgres
```
or
```
./sqltest -aws postgres
```

To update the expected output to be the output from postres, add `-update`.
```
./sqltest -update -postgres ...
```

Finally, to run all of the tests against all of the supported implementations, do the following:
```
./sqltest -postgres "host=<host>.rds.amazonaws.com port=5432 dbname=<dbname> user=<user> password=<password>" -mysql "<user>:<password>@tcp(<host>.rds.amazonaws.com)/<dbname>" postgres mysql sqlite3
```

To test against a local postgres instance, I use:
```
./sqltest -postgres "host=localhost port=5432 dbname=test sslmode=disable" postgres
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

See `testdata/sql/create.sql` and `testdata/sql/insert_bool.sql` for more examples of
template usage.

## Adding SQL Implementations

Package sqltestdb is used to test SQL compatibility between different implementations.

Implementations that have a database driver for Go should be added to `./gosql.go`.

Otherwise, `RunTests` should be called directly.
