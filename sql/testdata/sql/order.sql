--
-- Test SELECT ... ORDER BY
--
-- {{Sort .Global false}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int, c3 int, c4 int);

INSERT INTO tbl1 VALUES
    (3, 34, 5, 10),
    (6, 27, 8, 9),
    (27, 4, 29, 0),
    (12, 20, 14, 9),
    (15, 16, 17, 12),
    (18, 13, 20, 43),
    (9, 25, 11, 5),
    (21, 10, 23, 9),
    (24, 7, 26, 53),
    (30, 1, 32, 1),
    (0, 41, 2, 0);

-- {{Sort .Test true}}
SELECT * FROM tbl1;

SELECT * FROM tbl1 ORDER BY c1;

SELECT * FROM tbl1 ORDER by c2 ASC;

SELECT c2 FROM tbl1 ORDER BY c1;

SELECT * FROM tbl1 ORDER by c2 DESC;

SELECT c2 FROM tbl1 ORDER BY c1 DESC;

SELECT * FROM tbl1 ORDER by c4, c1;

SELECT c2, c3 FROM tbl1 ORDER by c4, c1 DESC;
