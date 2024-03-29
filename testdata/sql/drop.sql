--
-- Test DROP TABLE with FOREIGN KEY constraints
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}
-- {{if eq Dialect "maho-badger"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1 CASCADE;

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl1 (
    c1 int,
    c2 int PRIMARY KEY,
    c3 int
);

CREATE TABLE tbl2 (
    c4 int PRIMARY KEY,
    c5 int,
    c6 int REFERENCES tbl1 ON DELETE RESTRICT ON UPDATE RESTRICT
);

INSERT INTO tbl1 VALUES
    (100, 1, 10),
    (200, 2, 20),
    (300, 3, 30),
    (400, 4, 10),
    (500, 5, 20);

INSERT INTO tbl2 VALUES
    (10, 100, 1),
    (20, 200, 2),
    (30, 300, 3);

ALTER TABLE tbl1 ADD FOREIGN KEY (c3) REFERENCES tbl2 ON DELETE RESTRICT ON UPDATE RESTRICT;

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (90, 900, 9);

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (40, 400, 4),
    (50, 500, 5),
    (60, 600, 6);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

INSERT INTO tbl1 VALUES
    (600, 6, 30);

INSERT INTO tbl2 VALUES
    (40, 400, 4),
    (50, 500, 5),
    (60, 600, 6),
    (70, 700, 1),
    (80, 800, 2),
    (90, 900, 3);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

{{Fail .Test}}
DELETE FROM tbl1 WHERE c2 = 6;

{{Fail .Test}}
UPDATE tbl1 SET c2 = 44 WHERE c2 = 4;

INSERT INTO tbl1 VALUES
    (700, 7, 10),
    (800, 8, 20),
    (900, 9, 30);

SELECT * FROM tbl1;

{{Fail .Test}}
DELETE FROM tbl1 WHERE c2 > 3;

{{Fail .Test}}
UPDATE tbl1 SET c2 = c2 * 10 WHERE c2 > 5;

SELECT * FROM tbl1;

{{Fail .Test}}
DROP TABLE tbl1;

{{Fail .Test}}
DROP TABLE tbl2;

DROP TABLE tbl2 CASCADE;

DELETE FROM tbl1 WHERE c2 = 6;

UPDATE tbl1 SET c2 = 44 WHERE c2 = 4;

DELETE FROM tbl1 WHERE c2 > 3;

UPDATE tbl1 SET c2 = c2 * 10 WHERE c2 > 5;

INSERT INTO tbl1 VALUES
    (4, 44, 444),
    (555, 55, 5),
    (66, 666, 6);

SELECT * FROM tbl1;

DROP TABLE tbl1 CASCADE;

