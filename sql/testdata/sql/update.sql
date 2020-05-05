--
-- Test UPDATE
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int primary key, c2 int, c3 int);

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

UPDATE tbl1 SET c2 = c1 + c3 WHERE c1 < 6;

SELECT * FROM tbl1;

UPDATE tbl1 SET c3 = c3 * 5 WHERE c2 % 2 = 0;

SELECT * FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int primary key, c2 int, c3 int);

INSERT INTO tbl2 VALUES
    (0, 0, 0),
    (2, 2, 2),
    (4, 4, 4),
    (6, 6, 6),
    (8, 8, 8),
    (10, 10, 10);

SELECT * FROM tbl2;

UPDATE tbl2 SET c1 = c1 + 1, c2 = c2 + 1;

SELECT * FROM tbl2;

UPDATE tbl2 SET c1 = c1 - 1, c2 = c2 - 1;

SELECT * FROM tbl2;
