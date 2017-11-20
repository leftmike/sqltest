--
-- Test INSERT INTO
--
-- mysql has implicit default values for not null columns
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 varchar(64) not null, c2 int not null);

INSERT INTO tbl1 VALUES ('ABC', 456);

{{Fail .Test}}
INSERT INTO tbl1 (c1) VALUES ('ABC');

{{Fail .Test}}
INSERT INTO tbl1 (c2) VALUES (123);

INSERT INTO tbl1 (c1, c2) VALUES ('DEF', 789);

INSERT INTO tbl1 (c2, c1) VALUES (123, 'GHI');

SELECT * FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 varchar(64) default 'abc', c2 int default 123);

INSERT INTO tbl2 (c2) VALUES (456);

INSERT INTO tbl2 (c1) VALUES ('def');

SELECT * FROM tbl2;

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 varchar, c2 int, c3 varchar, c4 int);

INSERT INTO tbl3 (c4, c2, c3, c1) VALUES
    (10, 20, 'a', 'bb'),
    (20, 30, 'bb', 'ccc'),
    (30, 40, 'ccc', 'dddd'),
    (40, 50, 'dddd', 'eeeee'),
    (50, 60, 'eeeee', 'ffffff');

SELECT * from tbl3;
