--
-- Test CREATE TABLE
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 INT, c3 integer, c4 smallint, c5 bigint);

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int2, c2 int4, c3 int8);

CREATE TABLE IF NOT EXISTS tbl2 (c1 bool);

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 double precision, c2 real);

DROP TABLE IF EXISTS tbl3a;

CREATE TABLE tbl3a (c1 double {{if eq Dialect "postgres"}}precision{{end}}, c2 real);

DROP TABLE IF EXISTS tbl4;

CREATE TABLE tbl4 (c1 bool, c2 boolean);

DROP TABLE IF EXISTS tbl5;

CREATE TABLE tbl5 (c1 char, c2 char(200), c3 varchar(5), c4 text, c5 {{TEXT 123}});

DROP TABLE IF EXISTS tbl6;

-- {{eq Dialect "sqlite3" | not | Fail .Test}}
CREATE TABLE tbl6 (c1 badtype);

DROP TABLE IF EXISTS tbl7;

CREATE TABLE tbl7 (c1 {{BINARY}}, c2 {{VARBINARY 10}}, c3 {{BLOB}});

DROP TABLE IF EXISTS tbl8;

CREATE TABLE tbl8 (c1 {{BINARY 123}}, c2 {{BLOB 456}});

DROP TABLE IF EXISTS tbl9;

CREATE TABLE tbl9 (c1 char(64) not null, c2 varchar(64) not null, c3 bool not null,
    c4 int not null)

DROP TABLE IF EXISTS tbl10;

CREATE TABLE tbl10 (c1 char(64) default 'abc', c2 varchar(64) default 'def', c3 bool default true,
    c4 int default 123)
