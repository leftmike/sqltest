--
-- Test UNIQUE constraints with NULLs
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int primary key,
    c2 int unique,
    c3 int
);

INSERT INTO tbl1 VALUES
    (0, NULL, 0),
    (1, NULL, 0),
    (2, NULL, 0),
    (3, 40, 0);

SELECT * FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (4, 40, 0);

{{Fail .Test}}
UPDATE tbl1 SET c2 = 40 WHERE c1 = 2;

UPDATE tbl1 SET c2 = 30 WHERE c1 = 2;

SELECT * FROM tbl1;

UPDATE tbl1 SET c2 = NULL WHERE c1 = 3;

UPDATE tbl1 SET c2 = 40 WHERE c1 = 2;

SELECT * FROM tbl1;

{{Fail .Test}}
UPDATE tbl1 SET c2 = 40 WHERE c1 = 1;

INSERT INTO tbl1 VALUES
    (4, 60, 0),
    (5, NULL, 0),
    (6, 80, 0),
    (7, NULL, 0);

SELECT * FROM tbl1;
