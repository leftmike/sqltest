--
-- Test CHECK constraints with INSERT
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

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (-1, 10, 1);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (0, 9, 1);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (20, 10, 30);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (4, 10, 3);

SELECT * FROM tbl1;
