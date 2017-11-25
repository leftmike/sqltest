--
-- Test SELECT
--

SELECT 1 AS c1, 2 AS c2;

SELECT 1 + 2 AS c1, 2 + 3 AS c2, 4 * 5 AS c3;

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int, c3 int);

INSERT INTO tbl1 VALUES
    (0, 1, 2),
    (3, 4, 5),
    (6, 7, 8),
    (9, 10, 11),
    (12, 13, 14),
    (15, 16, 17),
    (18, 19, 20),
    (21, 22, 23),
    (24, 25, 26),
    (27, 28, 29),
    (30, 31, 32);

SELECT * FROM tbl1;

SELECT tbl1.* FROM tbl1;

SELECT c2 FROM tbl1;

SELECT c3, c1 FROM tbl1;

SELECT c3, c1 FROM tbl1 WHERE c2 < 10;

SELECT c2 FROM tbl1 WHERE c1 > 6 AND c3 <23;

SELECT c2, c3, c1 FROM tbl1 WHERE c1 = 24;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int, c2 int, c3 int);

INSERT INTO tbl2 VALUES
    (6, 60, 600),
    (7, 70, 700),
    (8, 80, 800),
    (9, 90, 900),
    (10, 100, 1000),
    (11, 110, 1100),
    (12, 120, 1200);

SELECT * FROM tbl2 WHERE c3 = c1 * c2;

SELECT tbl2.* FROM tbl2 WHERE c1 * 10 = c2 AND c2 * 10 = c3;
