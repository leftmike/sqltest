--
-- Test INSERT INTO using DEFAULT
--
-- Sqlite3 does not support the DEFAULT keyword for a column value
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 varchar(64) default 'abc', c2 int default 123);

INSERT INTO tbl1 (c2) VALUES (456);

INSERT INTO tbl1 (c1) VALUES ('def');

SELECT c1, c2 FROM tbl1;

INSERT INTO tbl1 (c2, c1) VALUES (DEFAULT, DEFAULT);

SELECT c1, c2 FROM tbl1;

INSERT INTO tbl1 VALUES
    ('1st', DEFAULT),
    (NULL, 789),
    (DEFAULT, NULL);

SELECT c1, c2 FROM tbl1;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (
    c1 int default 1,
    c2 int default 2 not null,
    c3 int not null);

INSERT INTO tbl2 VALUES
    (10, 20, 30),
    (20, 30, 40),
    (NULL, DEFAULT, 50),
    (DEFAULT, 40, 60);

SELECT c1, c2, c3 FROM tbl2;

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (DEFAULT, NULL, 10);

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (NULL, DEFAULT, NULL);

{{Fail .Test}}
INSERT INTO tbl2 (c1) VALUES (10);

{{Fail .Test}}
INSERT INTO tbl2 VALUES (10, 20, DEFAULT);

INSERT INTO tbl2 (c3) VALUES (70);

SELECT c1, c2, c3 FROM tbl2;
