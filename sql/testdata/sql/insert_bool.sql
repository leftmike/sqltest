--
-- Test INSERT INTO
--
-- sqlite3 and mysql don't have a boolean type
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int not null, c2 bool not null);

INSERT INTO tbl1 VALUES (456, true);

SELECT c1, c2 FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 (c1) VALUES (789);

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 bool default true, c2 int default 123);

INSERT INTO tbl2 (c2) VALUES (456);

SELECT c1, c2 FROM tbl2;

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 bool);

INSERT INTO tbl3 VALUES
    (TRUE),
    (true),
    ('t'),
    ('true'),
    ('y'),
    ('yes'),
    ('on'),
    ('1'),
    (FALSE),
    (false),
    ('f'),
    ('false'),
    ('n'),
    ('no'),
    ('off'),
    ('0');

SELECT c1 from tbl3;

DROP TABLE IF EXISTS tbl4;

CREATE TABLE tbl4 (c1 bool);

INSERT INTO tbl4 VALUES
    (true AND true),
    (true AND false),
    (false AND true),
    (false AND false),
    (true OR true),
    (true OR false),
    (false OR true),
    (false OR false),
    (NOT true),
    (NOT false),
    (NOT true AND true),
    (NOT true AND false),
    (NOT false AND true),
    (NOT false AND false),
    (true AND NOT true),
    (true AND NOT false),
    (false AND NOT true),
    (false AND NOT false),
    (1 = 2),
    (1 = 1),
    (2 = 1),
    (123 >= 123),
    (123 >= 1234),
    (123 >= -123),
    (-123 >= 123),
    (123 >= 12),
    (123 > 123),
    (123 > 1234),
    (123 > -123),
    (-123 > 123),
    (123 > 12),
    (123 <= 123),
    (123 <= 1234),
    (123 <= -123),
    (-123 <= 123),
    (123 <= 12),
    (123 < 123),
    (123 < 1234),
    (123 < -123),
    (-123 < 123),
    (123 < 12),
    (123 != 123),
    (123 != 1234),
    (123 != -123),
    (-123 != 123),
    (123 != 12);

{{Fail .Test}}
INSERT INTO tbl4 VALUES
    (1 = 'abc');

SELECT c1 FROM tbl4;
