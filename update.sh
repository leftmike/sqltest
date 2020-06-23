#!/bin/bash
for sql in testdata/postgres/sql/*.sql
do
    name=`basename $sql .sql`
    psql -A -f $sql > testdata/postgres/expected/$name.out
done
