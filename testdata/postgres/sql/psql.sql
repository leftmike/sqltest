--
-- Test psql output format
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 text primary key, integer_column int, c3 text);

INSERT INTO tbl1 VALUES
    ('one', 1, NULL),
    ('two', 2, 'two two two two two two two'),
    ('two hundred', 200, 'two hundred');

SELECT * FROM tbl1;
