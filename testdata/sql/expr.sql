--
-- Test expressions
--

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

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 text, c2 int);

INSERT INTO tbl2 VALUES
    ('one', 1),
    ('two', 2),
    ('three', 3),
    ('four', 4),
    ('five', 5),
    ('six', 6),
    ('seven', 7),
    ('eight', 8);

SELECT (SELECT c1 FROM tbl2 WHERE c2 = tbl1.c1) AS e1, c2 FROM tbl1;

SELECT (SELECT c1 FROM tbl2 WHERE c2 = tbl1.c1 AND c1 != 'one') AS e1, c2 FROM tbl1;

SELECT (SELECT * FROM (SELECT c1 FROM tbl2 WHERE c2 = tbl1.c1) AS tbl3) AS e1, c2 FROM tbl1;
