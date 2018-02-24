--
-- Test DELETE
--

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

DELETE FROM tbl1 WHERE c1 < 6;

SELECT * FROM tbl1;

DELETE FROM tbl1 WHERE c2 = 16;

SELECT * FROM tbl1;

DELETE FROM tbl1;

SELECT * FROM tbl1;
