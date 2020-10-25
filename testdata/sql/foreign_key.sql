--
-- Test FOREIGN KEY constraints
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int,
    c2 int PRIMARY KEY,
    c3 int
);

CREATE TABLE tbl2 (
    c4 int PRIMARY KEY,
    c5 int,
    c6 int REFERENCES tbl1
);

INSERT INTO tbl1 VALUES
    (10, 1, 100),
    (20, 2, 200),
    (30, 3, 300),
    (40, 4, 400),
    (50, 5, 500);

INSERT INTO tbl2 VALUES
    (10, 100, 1),
    (20, 200, 2),
    (30, 300, 3);

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (90, 900, 9);

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (40, 400, 4),
    (50, 500, 5),
    (60, 600, 6);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

INSERT INTO tbl1 VALUES
    (60, 6, 600);

INSERT INTO tbl2 VALUES
    (40, 400, 4),
    (50, 500, 5),
    (60, 600, 6),
    (70, 700, 1),
    (80, 800, 2),
    (90, 900, 3);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

-- Drop table with foreign key reference first.
DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl1;
