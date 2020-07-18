--
-- Test PRIMARY KEYs
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int primary key,
    c2 int,
    c3 int
);

INSERT INTO tbl1 VALUES
    (0, 10, 0),
    (1, 20, 0),
    (2, 30, 0),
    (3, 40, 0);

SELECT * FROM tbl1;

UPDATE tbl1 SET c1 = c1 * 10;

SELECT * FROM tbl1;
