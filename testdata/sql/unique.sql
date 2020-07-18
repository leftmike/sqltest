--
-- Test UNIQUE constraints
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int primary key,
    c2 int unique,
    c3 int
);

INSERT INTO tbl1 VALUES
    (0, 10, 0),
    (1, 20, 0),
    (2, 30, 0),
    (3, 40, 0);

SELECT * FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (4, 10, 0);

SELECT * FROM tbl1;

INSERT INTO tbl1 VALUES
    (4, 50, 0),
    (5, 60, 0),
    (6, 70, 0);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (7, 80, 0),
    (8, 90, 0),
    (9, 10, 0),
    (10, 110, 0);

SELECT * FROM tbl1;

UPDATE tbl1 SET c2 = c2 + 1 WHERE c1 < 4;

SELECT * FROM tbl1;

{{Fail .Test}}
UPDATE tbl1 SET c2 = c2 + 10;

SELECT * FROM tbl1;

UPDATE tbl1 SET c1 = c1 * 10;

SELECT * FROM tbl1;

UPDATE tbl1 SET c2 = c2 + 1 WHERE c1 < 4;

SELECT * FROM tbl1;

{{Fail .Test}}
UPDATE tbl1 SET c2 = c2 + 10 WHERE c1 = 40;

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (40, 51, 0);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (41, 50, 0);

UPDATE tbl1 SET c3 = 1;

SELECT * FROM tbl1;
