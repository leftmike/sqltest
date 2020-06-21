--
-- Test INSERT INFO
--
-- postgres will space pad strings inserted into char(n) columns so that they are n
-- characters in width
-- {{if eq Dialect "postgres"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 char(64) not null, c2 varchar(64) not null, c3 int not null);

INSERT INTO tbl1 VALUES ('ABC', 'DEF', 456);

SELECT c1, c2, c3 FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 char(64) default 'abc', c2 varchar(64) default 'def', c3 int default 123);

INSERT INTO tbl2 (c3) VALUES (456);

SELECT c1, c2, c3 FROM tbl2;
