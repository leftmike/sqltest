--
-- Test FOREIGN KEY constraints
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl3;

DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int PRIMARY KEY,
    c2 int,
    c3 int
);

CREATE TABLE tbl2 (
    c4 int PRIMARY KEY,
    c5 int,
    c6 int REFERENCES tbl1 ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO tbl1 VALUES
    (1, 10, 100),
    (2, 20, 200),
    (3, 30, 300),
    (4, 40, 400),
    (5, 50, 500),
    (6, 60, 600),
    (7, 70, 700),
    (8, 80, 800);

INSERT INTO tbl2 VALUES
    (10, 100, 1),
    (20, 200, 2),
    (30, 300, 3),
    (40, 400, 1),
    (50, 500, 2),
    (60, 600, 3),
    (70, 700, 4),
    (80, 800, 5);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

DELETE FROM tbl1 WHERE c1 = 7;

DELETE FROM tbl1 WHERE c1 = 2;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

UPDATE tbl1 SET c1 = 2 WHERE c1 = 1;

UPDATE tbl1 SET c1 = 9 WHERE c1 = 5;

UPDATE tbl1 SET c2 = 400 WHERE c1 = 4;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

-- Drop table with foreign key reference first.
DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int PRIMARY KEY,
    c2 int,
    x1 int DEFAULT 0
);

CREATE TABLE tbl2 (
    c3 int PRIMARY KEY,
    c4 int REFERENCES tbl1 ON DELETE CASCADE ON UPDATE CASCADE,
    x2 int DEFAULT 0
);

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (
    c5 int PRIMARY KEY,
    c6 int REFERENCES tbl2 ON DELETE RESTRICT ON UPDATE RESTRICT,
    x3 int DEFAULT 0
);

INSERT INTO tbl1 (c1, c2) VALUES
    (1, 10),
    (2, 20),
    (3, 30),
    (4, 40),
    (5, 50),
    (6, 60);

INSERT INTO tbl2 (c3, c4) VALUES
    (10, 1),
    (11, 1),
    (12, 1),
    (20, 2),
    (30, 3),
    (31, 3);

INSERT INTO tbl3 (c5, c6) VALUES
    (100, 11),
    (101, 11),
    (200, 20),
    (201, 20),
    (202, 20),
    (300, 31);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

DELETE FROM tbl1 WHERE c1 = 10;

DELETE FROM tbl2 WHERE c3 = 30;

{{Fail .Test}}
DELETE FROM tbl2 WHERE c3 = 20;

{{Fail .Test}}
DELETE FROM tbl1 WHERE c1 = 2;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

{{Fail .Test}}
UPDATE tbl2 SET c3 = 200 WHERE c3 = 20;

UPDATE tbl1 SET c1 = 100 WHERE c1 = 1;

UPDATE tbl1 SET x1 = 100 WHERE c1 = 100;

UPDATE tbl2 SET x2 = 200 WHERE c3 = 20;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

-- Drop table with foreign key references first.
DROP TABLE IF EXISTS tbl3;

DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;
