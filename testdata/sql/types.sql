--
-- Test types of columns
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}
-- {{Types .Global true}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 smallint, c2 int2, c3 int, c4 int4, c5 int8, c6 bigint);

INSERT INTO tbl1 VALUES
    (1, 2, 3, 4, 5, 6),
    (-7, -8, -9, -10, -11, -12),
    (13, 14, 15, 16, 17, 18);

SELECT c1, c2, c3, c4, c5, c6 FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 double precision, c2 real);

INSERT INTO tbl2 VALUES
    (1.23, 4.5678),
    (-7.891, -2.345);

SELECT c1, c2 FROM tbl2;

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 bool);

INSERT INTO tbl3 VALUES (true), (false);

SELECT c1 FROM tbl3;

DROP TABLE IF EXISTS tbl4;

CREATE TABLE tbl4 (c1 char, c2 char(2), c3 varchar(5), c4 text);

INSERT INTO tbl4 VALUES
    ('a', 'bb', 'ccc', 'dddd'),
    ('A', 'BB', 'CCC', 'DDDD');

SELECT c1, c2, c3, c4 FROM tbl4;

DROP TABLE IF EXISTS tbl5;

CREATE TABLE tbl5 (c1 BYTEA);

INSERT INTO tbl5 VALUES ('A'), ('bb'), ('CCC');

SELECT c1 FROM tbl5;
