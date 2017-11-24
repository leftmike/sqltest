--
-- Test INSERT INTO using DEFAULT
--
-- Sqlite3 does not support the DEFAULT keyword for a column value
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 varchar(64) default 'abc', c2 int default 123);

INSERT INTO tbl1 (c2) VALUES (456);

INSERT INTO tbl1 (c1) VALUES ('def');

SELECT * FROM tbl1;

INSERT INTO tbl1 (c2, c1) VALUES (DEFAULT, DEFAULT);

SELECT * FROM tbl1;
