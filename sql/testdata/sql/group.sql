--
-- Test GROUP BY
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int, c3 int, c4 int);

INSERT INTO tbl1 VALUES
    (1, 2, 3, 4),
    (1, 2, 33, 44),
    (2, 1, 3, NULL),
    (2, 2, 3, 4),
    (2, 3, 4, NULL),
    (3, 4, 5, 6),
    (4, 5, 6, 7),
    (4, 1, NULL, 5),
    (4, 1, 5, 6),
    (5, 1, 6, 7),
    (5, 2, NULL, NULL);

SELECT c1 FROM tbl1 GROUP BY c1;

SELECT c1 + c2 AS c FROM tbl1 GROUP BY c1 + c2;

SELECT c1 + c2 AS d1, (c1 + c2) * 10 AS d2 FROM tbl1 GROUP BY c1 + c2;

SELECT count(*) AS count_all, c1 + c2 AS c, count(c3), count(c4) FROM tbl1 GROUP BY c1 + c2;

SELECT c1, c2, count(*) AS count_all FROM tbl1 GROUP BY c1, c2;

SELECT c1, c2, count(*) AS count_all FROM tbl1 GROUP BY c2, c1;
