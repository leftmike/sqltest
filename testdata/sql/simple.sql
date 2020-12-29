--
-- A simple test to start with.
--

DROP TABLE IF EXISTS tbl;

CREATE TABLE tbl (c1 INTEGER, c2 INTEGER, c3 INTEGER);

INSERT INTO tbl VALUES
    (1, 2, 3),
    (7, 8, 9),
    (4, 5, 6);

/*
SELECT d1, d2, d3 FROM table;
*/

SELECT c1, c2, c3 FROM tbl;
