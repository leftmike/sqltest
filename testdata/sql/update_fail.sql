--
-- Test UPDATE failing with primary key
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int primary key, c2 int, c3 int);

INSERT INTO tbl1 VALUES
    (0, 0, 0),
    (2, 2, 2),
    (4, 4, 4),
    (6, 6, 6),
    (8, 8, 8),
    (10, 10, 10);

SELECT * FROM tbl1;

UPDATE tbl1 SET c1 = c1 + 1, c2 = c2 + 1;

SELECT * FROM tbl1;

UPDATE tbl1 SET c1 = c1 - 1, c2 = c2 - 1;

SELECT * FROM tbl1;

{{Fail .Test}}
UPDATE tbl1 SET c1 = c1 + 2, c2 = c2 + 2;
