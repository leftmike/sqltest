--
-- Test ALTER TABLE ... DROP
--
-- {{if eq Dialect "maho-kvrows"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (
    c1 int primary key,
    c2 int constraint con1 check (c2 >= 10),
    c3 int unique,
    c4 int constraint con2 not null,
    c5 int constraint con3 default -1,
    unique (c2, c3)
);

{{Fail .Test}}
ALTER TABLE tbl1 DROP CONSTRAINT con99;

ALTER TABLE tbl1 DROP CONSTRAINT IF EXISTS con99;

INSERT INTO tbl1 VALUES
    (1, 10, 11, 100, DEFAULT);

SELECT * FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (2, -10, 12, 200, 2);

ALTER TABLE tbl1 DROP CONSTRAINT con1;

INSERT INTO tbl1 VALUES
    (2, -10, 12, 200, 2);

SELECT * FROM tbl1;

{{Fail .Test}}
INSERT INTO tbl1 VALUES
    (3, 30, 13, NULL, 3);

ALTER TABLE tbl1 ALTER c4 DROP NOT NULL;

INSERT INTO tbl1 VALUES
    (3, 30, 13, NULL, 3);

SELECT * FROM tbl1;

INSERT INTO tbl1 VALUES
    (4, 40, 14, NULL, DEFAULT);

ALTER TABLE tbl1 ALTER c5 DROP DEFAULT;

INSERT INTO tbl1 VALUES
    (5, 50, 15, NULL, DEFAULT);

SELECT * FROM tbl1;
