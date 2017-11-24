--
-- For testing sqltest: this does not need to be valid SQL.
--

SELECT * FROM table

-- Go ahead and insert 3 rows into the pretend table.
INSERT INTO table (col1, col2, col3, col4)
VALUES
    (1, 2, 3, 4, 5),
    (6, 7, 8, 9, 10),
    (11, 12, 13, 14, 15)

-- Delete some stuff.
DELETE FROM table col1 = 6

-- Get back some results.
SELECT * FROM table WHERE col1 > 0
