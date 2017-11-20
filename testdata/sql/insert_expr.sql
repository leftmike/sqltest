--
-- Test INSERT INTO expressions
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int, c2 int, c3 int);

INSERT INTO tbl1 VALUES
    (1 + 2, 3 * 4, 6 / 2),
    (12 * 6 + 3 * 2, 12 * 4 + 3, 12 + 4 * 3),
    (12 * (6 + 3), (12 + 4) * 3, 8 / 17);

SELECT * FROM tbl1;

INSERT INTO tbl1 (c1, c3) VALUES (1, 2), (3, 4), (5, 6), (7, 8);

SELECT * FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int);

INSERT INTO tbl2 VALUES
    (12 + 34),
    (1234 & 5678),
    (1234 | 5678),
    (4567 / 89),
    (1 << 23),
    (45 % 13),
    (123 * 45),
    (123 - 4567),
    (123456 >> 7),
    (- 123),
    (abs(123)),
    (abs(-123));

SELECT * FROM tbl2;
