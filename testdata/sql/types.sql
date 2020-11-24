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

SELECT * FROM tbl1;

