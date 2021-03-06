--
-- Test CHECK constraints with UPDATE
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int primary key check (c1 >= 0),
    c2 int check (c2 >= 10),
    c3 int,
    check (c2 > c1),
    check (c3 > c1)
);

INSERT INTO tbl1 VALUES
    (0, 10, 1),
    (1, 10, 2),
    (2, 10, 3),
    (3, 10, 4);

SELECT * FROM tbl1;

UPDATE tbl1 SET c2 = 11 WHERE c1 = 0;

{{Fail .Test}}
UPDATE tbl1 SET c2 = 0 WHERE c1 = 0;

{{Fail .Test}}
UPDATE tbl1 SET c1 = -1 WHERE c3 = 3;

{{Fail .Test}}
UPDATE tbl1 SET c2 = 0 WHERE c1 = 3;

{{Fail .Test}}
UPDATE tbl1 SET c3 = c1;

SELECT * FROM tbl1;
