--
-- Test creating and dropping indexes
--

DROP TABLE IF EXISTS tbl1;

CREATE TABLE tbl1 (c1 int PRIMARY KEY, c2 char(20), c3 text, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl1a;

{{Fail .Test}}
CREATE TABLE tbl1a (c1 int PRIMARY, c2 char(20), c3 text, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl1b;

CREATE TABLE tbl1b (c1 int, c2 char(20) PRIMARY KEY, c3 text, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl1c;

CREATE TABLE tbl1c (c1 int, c2 char(20), c3 text PRIMARY KEY, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl1d;

CREATE TABLE tbl1d (c1 int, c2 char(20), c3 text, c4 {{BINARY}} PRIMARY KEY)

DROP TABLE IF EXISTS tbl2;

CREATE TABLE tbl2 (c1 int UNIQUE DEFAULT 123, c2 char(20), c3 text, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl2b;

CREATE TABLE tbl2b (c1 int, c2 char(20) UNIQUE, c3 text, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl2c;

CREATE TABLE tbl2c (c1 int, c2 char(20), c3 text UNIQUE, c4 {{BINARY}})

DROP TABLE IF EXISTS tbl2d;

CREATE TABLE tbl2d (c1 int, c2 char(20), c3 text, c4 {{BINARY}} UNIQUE)

DROP TABLE IF EXISTS tbl3;

CREATE TABLE tbl3 (c1 int, c2 int, c3 text,
    PRIMARY KEY (c1))

DROP TABLE IF EXISTS tbl3a;

CREATE TABLE tbl3a (c1 int, c2 int, c3 text,
    PRIMARY KEY (c1, c2));

DROP TABLE IF EXISTS tbl3b;

CREATE TABLE tbl3b (c1 int, c2 int, c3 text,
    PRIMARY KEY (c3, c2, c1));

DROP TABLE IF EXISTS tbl3c;

CREATE TABLE tbl3c (c1 int, c2 int, c3 text,
    PRIMARY KEY (c1),
    UNIQUE (c2, c3),
    UNIQUE (c1, c2));

DROP TABLE IF EXISTS tbl4;

CREATE TABLE tbl4 (c1 int, c2 int, c3 text);

CREATE INDEX idx1 ON tbl4 (c1);

CREATE INDEX IF NOT EXISTS idx1 ON tbl4 (c1);

{{Fail .Test}}
CREATE INDEX idx1 ON tbl4 (c1);

CREATE INDEX idx2 ON tbl4 (c1, c2);

CREATE UNIQUE INDEX idx3 ON tbl4 (c2, c3);
