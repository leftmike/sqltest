--
-- Test expressions with values
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int);

INSERT INTO tbl1 VALUES
    (1, 10),
    (2, 20),
    (3, 30),
    (4, 40),
    (5, 50),
    (6, 60),
    (7, 70),
    (8, 80);

SELECT
    (SELECT v2
        FROM (VALUES (tbl1.c1, tbl1.c1*2)) AS tbl3 (v1, v2)
        WHERE v1 = tbl1.c1) AS e1, c2 FROM tbl1;

SELECT
    (SELECT v2
        FROM (SELECT tbl1.c1, tbl1.c1*2) AS tbl3 (v1, v2)
        WHERE v1 = tbl1.c1) AS e1, c2 FROM tbl1;
