--
-- Test subqueries
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int PRIMARY KEY, c2 int, c3 int, c4 int);

INSERT INTO tbl1 VALUES
    (1, 10, 11, 100),
    (2, 20, 22, 200),
    (3, 30, 33, 300),
    (4, 40, 44, 400),
    (5, 50, 55, 500),
    (6, 60, 66, 600),
    (7, 70, 77, 700),
    (8, 80, 88, 800),
    (9, 90, 99, 900);

SELECT * FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int PRIMARY KEY, c2 TEXT);

INSERT INTO tbl2 VALUES
    (10, 'ten'),
    (20, 'twenty'),
    (30, 'thirty'),
    (40, 'forty'),
    (50, 'fifty');

SELECT * FROM tbl2;

SELECT col2, col1 FROM (SELECT c4, c2 FROM tbl1 WHERE c1 >= 4) AS t1 (col1, col2);

-- {{Fail .Test}}
SELECT col2, col1, c2 FROM (SELECT c4, c2 FROM tbl1) AS t1 (col1, col2);

-- {{Fail .Test}}
SELECT col2, col1 FROM (SELECT c4, c2 FROM tbl1 WHERE c1 >= 4) AS t1 (col1, col2) WHERE c2 <= 70;

SELECT col2, col1 FROM (SELECT c4, c2 FROM tbl1 WHERE c1 >= 4) AS t1 (col1, col2) WHERE col2 <= 70;

SELECT col2, col1
    FROM (SELECT c4 + c3, tbl2.c2 FROM tbl1 JOIN tbl2 ON tbl1.c2 = tbl2.c1) AS t1 (col1, col2)
    WHERE col1 % 2 = 1;

SELECT * FROM (VALUES (1, 2, 3), (4, 5, 6), (7, 8, 9)) AS t1 (c1, c2, c3);

SELECT *
    FROM (SELECT c3, c2 + c1 FROM (VALUES (1, 2, 3), (4, 5, 6), (7, 8, 9)) AS t1 (c1, c2, c3))
    AS t2 (col1, col2);

SELECT c1, c2, EXISTS(SELECT 1 FROM tbl2 WHERE c1 = tbl1.c2) AS e3 FROM tbl1;

SELECT c1, c2, EXISTS(SELECT 1 FROM tbl2 WHERE c1 = 10) AS e3 FROM tbl1;

SELECT c1, c2, EXISTS(SELECT 1 FROM tbl2 WHERE c1 = 60) AS e3 FROM tbl1;

SELECT c1, c2, EXISTS(SELECT 1 FROM tbl2 WHERE tbl1.c2 > 50) AS e3 FROM tbl1;

SELECT c1, c2 FROM tbl1 WHERE c2 IN (SELECT c1 FROM tbl2);

SELECT c1, c2 FROM tbl1 WHERE c2 NOT IN (SELECT c1 FROM tbl2);

SELECT c1, c2 IN (SELECT c1 FROM tbl2) AS e1 FROM tbl1;

SELECT c1, c2 NOT IN (SELECT c1 FROM tbl2) AS e1 FROM tbl1;

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 int PRIMARY KEY, c2 TEXT, c3 int);

INSERT INTO tbl3 VALUES
    (10, 'ten', 11),
    (20, NULL, 22),
    (30, 'thirty', 33),
    (40, NULL, NULL),
    (50, 'fifty', 55);

SELECT * FROM tbl3;

SELECT c1, c3 FROM tbl1 WHERE c3 IN (SELECT c3 FROM tbl3);

SELECT c1, c3 FROM tbl1 WHERE c3 NOT IN (SELECT c3 FROM tbl3);

SELECT c1, c3 IN (SELECT c3 FROM tbl3) AS e1 FROM tbl1;

SELECT c1, c3 NOT IN (SELECT c3 FROM tbl3) AS e1 FROM tbl1;

-- ANY

SELECT c1, c2 FROM tbl1 WHERE c2 >= ANY (SELECT c1 + 20 FROM tbl2);

SELECT c1, c2 < ANY (SELECT c1 - 10 FROM tbl2) AS e1 FROM tbl1;

-- ALL

SELECT c1, c2 FROM tbl1 WHERE c2 >= ALL (SELECT c1 + 20 FROM tbl2);

SELECT c1, c2 < ALL (SELECT c1 + 40 FROM tbl2) AS e1 FROM tbl1;
