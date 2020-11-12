--
-- Test UPDATE using DEFAULT
--
-- Sqlite3 does not support the DEFAULT keyword for a column value
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int default 1,
    c2 int default 2 not null,
    c3 int not null);

INSERT INTO tbl1 VALUES
    (10, 20, 30),
    (20, 30, 40),
    (30, 40, 50),
    (40, 50, 60),
    (50, 60, 70);

SELECT c1, c2, c3 FROM tbl1;

UPDATE tbl1 SET c3 = 700 WHERE c1 = 50;

{{Fail .Test}}
UPDATE tbl1 SET c3 = NULL WHERE c1 = 40;

{{Fail .Test}}
UPDATE tbl1 SET c2 = NULL WHERE c1 = 40;

UPDATE tbl1 SET c1 = NULL WHERE c1 = 10;

UPDATE tbl1 SET c1 = DEFAULT WHERE c1 = 20;

{{Fail .Test}}
UPDATE tbl1 SET c3 = DEFAULT WHERE c1 = 40;

{{Fail .Test}}
UPDATE tbl1 SET c3 = NULL WHERE c1 = 40;

{{Fail .Test}}
UPDATE tbl1 SET c2 = NULL WHERE c1 = 40;

UPDATE tbl1 SET c2 = DEFAULT WHERE c1 = 40;

SELECT c1, c2, c3 FROM tbl1;
