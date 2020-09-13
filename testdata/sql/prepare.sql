--
-- Test PREPARE and EXECUTE
--
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int primary key, c2 int, c3 int);

PREPARE delete1 AS DELETE FROM tbl1 WHERE c1 = $1;

PREPARE insert1 AS INSERT INTO tbl1 VALUES
    ($1, $1 * 10, $1 * 100),
    ($2, $2 * 2, $2 * 20);

PREPARE update1 AS UPDATE tbl1 SET c2 = $2 WHERE c1 = $1;

PREPARE values1 AS VALUES
    ($1 + 0, $1 + 1, $1 * 1),
    ($2 + 0, $2 + 2, $2 * 2),
    ($3 + 0, $3 + 3, $3 * 3);

EXECUTE insert1 (1, 2);

SELECT * FROM tbl1;

EXECUTE insert1 (3, 4);

SELECT * FROM tbl1;

SELECT * FROM tbl1 WHERE c1 >= 0 AND c1 <= 100;

EXECUTE delete1 (3);

SELECT * FROM tbl1;

EXECUTE update1 (4, 100);

SELECT * FROM tbl1;
