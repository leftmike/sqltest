--
-- Test INSERT INTO
--
-- sqlite3 and mysql don't have a boolean type
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int not null, c2 bool not null);

INSERT INTO tbl1 VALUES (456, true);

SELECT * FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 (c1) VALUES (789);

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 bool default true, c2 int default 123);

INSERT INTO tbl2 (c2) VALUES (456);

SELECT * FROM tbl2;
