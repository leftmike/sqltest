--
-- Test INSERT INTO with || and concat expressions
--
-- sqlite3 doesn't have concat and mysql doesn't have ||
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 varchar(128));

INSERT INTO tbl1 VALUES
    ('abc' || 'def'),
    (concat('ABC', 'DEF')),
    ('abc' || 123),
    (123 || 'abc'),
    (concat('ABC', 123)),
    (concat(123, 'ABC')),
    ('abc' || true),
    (false || 'abc');

SELECT c1 FROM tbl1;
