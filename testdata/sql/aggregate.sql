--
-- Test SELECT ... w/ aggregate functions
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int, c3 int);

INSERT INTO tbl1 VALUES
    (0, 1, 2),
    (0, 4, 5),
    (0, 7, 8),
    (1, 10, 11),
    (1, 13, 14),
    (2, 16, 17),
    (2, 19, 20),
    (2, 22, 23),
    (2, 25, 26),
    (3, 28, 29),
    (4, 31, 32);

SELECT count(*) AS count FROM tbl1;

SELECT count(*) AS count FROM tbl1 WHERE c1 = 5;

SELECT count(c1) FROM tbl1;

SELECT count(c1) FROM tbl1 WHERE c1 = 5;

SELECT sum(c1) FROM tbl1 GROUP BY c1;

SELECT sum(c1) FROM tbl1 WHERE c1 = 5;

SELECT sum(c1) FROM tbl1 WHERE c1 = 5 GROUP BY c1;

SELECT max(c2), min(c3) FROM tbl1;

SELECT max(c2), min(c3) FROM tbl1 WHERE c1 = 5;

SELECT max(c2), min(c3) FROM tbl1 GROUP BY c1;

SELECT max(c2), min(c3) FROM tbl1 WHERE c1 = 5 GROUP BY c1;
