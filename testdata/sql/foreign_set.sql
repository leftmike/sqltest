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
    c6 int REFERENCES tbl1 ON DELETE SET DEFAULT ON UPDATE SET NULL DEFAULT 1
);

CREATE TABLE tbl3 (
    c4 int PRIMARY KEY,
    c5 int,
    c6 int REFERENCES tbl1 ON DELETE SET NULL ON UPDATE SET DEFAULT DEFAULT 4 NOT NULL
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
    (30, 300, 1),
    (40, 400, 2),
    (50, 500, 6),
    (60, 600, 7),
    (70, 700, 8);

INSERT INTO tbl3 VALUES
    (10, 100, 3),
    (20, 200, 3),
    (30, 300, 4),
    (40, 400, 4),
    (50, 500, 5),
    (60, 600, 6),
    (70, 700, 7);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

DELETE FROM tbl1 WHERE c1 = 2;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

UPDATE tbl1 SET c1 = 10 WHERE c1 = 1;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

UPDATE tbl2 SET c6 = 10 WHERE c4 <= 40;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

{{Fail .Test}}
DELETE FROM tbl1 WHERE c1 = 10;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

{{Fail .Test}}
DELETE FROM tbl1 WHERE c1 = 3;

UPDATE tbl1 SET c1 = 30 WHERE c1 = 3;

SELECT * FROM tbl1;

SELECT * FROM tbl2;

SELECT * FROM tbl3;

-- Drop table with foreign key reference first.
DROP TABLE IF EXISTS tbl3;

DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;
