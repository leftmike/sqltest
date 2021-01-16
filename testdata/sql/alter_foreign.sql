--
-- Test ALTER TABLE with FOREIGN KEY
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1 CASCADE;

DROP TABLE IF EXISTS tbl2;

DROP TABLE IF EXISTS tbl3 CASCADE;

CREATE TABLE tbl1 (
    c1 int,
    c2 int PRIMARY KEY,
    c3 int,
    c4 text
);

CREATE TABLE tbl2 (
    c5 int PRIMARY KEY,
    c6 int,
    c7 text,
    c8 int,
    UNIQUE (c6),
    UNIQUE (c8, c7)
);

CREATE TABLE tbl3 (
    c1 int,
    c2 int PRIMARY KEY,
    c3 int,
    c4 text
);

{{Fail .Test}}
ALTER TABLE tbl4 ADD FOREIGN KEY (c1) REFERENCES tbl1;

{{Fail .Test}}
ALTER TABLE tbl1 ADD FOREIGN KEY (c1) REFERENCES tbl4;

{{Fail .Test}}
ALTER TABLE tbl1 ADD FOREIGN KEY (col1) REFERENCES tbl2;

{{Fail .Test}}
ALTER TABLE tbl1 ADD FOREIGN KEY (c4) REFERENCES tbl2;

{{Fail .Test}}
ALTER TABLE tbl1 ADD FOREIGN KEY (c1) REFERENCES tbl2 (c7);

INSERT INTO tbl1 VALUES
    (10, 1, 100, 'one hundred'),
    (20, 2, 200, 'two hundred'),
    (30, 3, 333, 'three hundred thirty three');

SELECT * FROM tbl1;

SELECT COUNT(*) AS cnt
    FROM tbl1 LEFT JOIN tbl2 ON c4 = c7 AND c3 = c8
    WHERE (c7 IS NULL) OR (c8 IS NULL);

{{Fail .Test}}
ALTER TABLE tbl1 ADD FOREIGN KEY (c4, c3) REFERENCES tbl2 (c7, c8);

DELETE FROM tbl1 WHERE c3 = 333;

INSERT INTO tbl2 VALUES
    (11, 1, 'one hundred', 100),
    (22, 2, 'two hundred', 200),
    (33, 3, 'three hundred', 300),
    (44, 4, 'four hundred', 400),
    (55, 5, 'five hundred', 500);

SELECT * FROM tbl2;

ALTER TABLE tbl1 ADD FOREIGN KEY (c4, c3) REFERENCES tbl2 (c7, c8);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (30, 3, 333, 'three hundred thirty three');

INSERT INTO tbl1 VALUES
    (30, 3, 300, 'three hundred'),
    (40, 4, 400, 'four hundred'),
    (50, 5, 500, 'five hundred');

ALTER TABLE tbl2 ADD FOREIGN KEY (c6) REFERENCES tbl1;

ALTER TABLE tbl1 ADD FOREIGN KEY (c2) REFERENCES tbl2 (c6);

ALTER TABLE tbl3
    ADD FOREIGN KEY (c4, c3) REFERENCES tbl2 (c7, c8),
    ADD FOREIGN KEY (c2) REFERENCES tbl2 (c6);

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (60, 6, 600, 'six hundred');

{{Fail .Test}}
INSERT INTO tbl2 VALUES
    (66, 6, 'six hundred', 600);

SELECT * FROM tbl1;

SELECT * FROM tbl2;

DROP TABLE IF EXISTS tbl1 CASCADE;

INSERT INTO tbl2 VALUES
    (66, 6, 'six hundred', 600);

UPDATE tbl2 SET c8 = c8 * 10;

SELECT * FROM tbl2;

DROP TABLE IF EXISTS tbl2 CASCADE;

DROP TABLE IF EXISTS tbl3 CASCADE;
