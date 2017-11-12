--
-- Test CREATE TABLE
--

DROP TABLE IF EXISTS tbl1

CREATE TABLE tbl1 (c1 int, c2 INT, c3 integer, c4 smallint, c5 bigint)

DROP TABLE IF EXISTS tbl2

CREATE TABLE tbl2 (c1 int2, c2 int4, c3 int8)

DROP TABLE IF EXISTS tbl3

CREATE TABLE tbl3 (c1 double precision, c2 real)

DROP TABLE IF EXISTS tbl3a

CREATE TABLE tbl3a (c1 double {{if eq $.Dialect "postgres"}}precision{{end}}, c2 real)

DROP TABLE IF EXISTS tbl4

CREATE TABLE tbl4 (c1 bool, c2 boolean)

DROP TABLE IF EXISTS tbl5

CREATE TABLE tbl5 (c1 char, c2 char(200), c3 varchar(5))

DROP TABLE IF EXISTS tbl6

{{eq $.Dialect "sqlite3" | not | Fail $.Test}}
CREATE TABLE tbl6 (c1 badtype)
