--
-- Test SELECT and GROUP BY
--
-- Based on Google's Cloud Platform examples: https://cloud.google.com/bigquery/docs/reference/standard-sql/query-syntax#appendix-a-examples-with-sample-data
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

DROP TABLE IF EXISTS Roster;

CREATE TABLE Roster (LastName {{TEXT 128}}, SchoolID int);

INSERT INTO Roster VALUES
    ('Adams', 50),
    ('Buchanan', 52),
    ('Coolidge', 52),
    ('Davis', 51),
    ('Eisenhower', 77);

DROP TABLE IF EXISTS PlayerStats;

CREATE TABLE PlayerStats (LastName {{TEXT 128}}, OpponentID int, PointsScored int);

INSERT INTO PlayerStats VALUES
    ('Adams', 51, 3),
    ('Buchanan', 77, 0),
    ('Coolidge', 77, 1),
    ('Adams', 52, 4),
    ('Buchanan', 50, 13);

DROP TABLE IF EXISTS TeamMascot;

CREATE TABLE TeamMascot (SchoolId int, Mascot {{TEXT 128}});

INSERT INTO TeamMascot VALUES
    (50, 'Jaguars'),
    (51, 'Knights'),
    (52, 'Lakers'),
    (53, 'Mustangs');

SELECT * FROM Roster JOIN TeamMascot ON Roster.SchoolID = TeamMascot.SchoolID;

SELECT * FROM Roster CROSS JOIN TeamMascot;

SELECT * FROM Roster FULL JOIN TeamMascot ON Roster.SchoolID = TeamMascot.SchoolID;

SELECT * FROM Roster LEFT JOIN TeamMascot ON Roster.SchoolID = TeamMascot.SchoolID;

SELECT * FROM Roster RIGHT JOIN TeamMascot ON Roster.SchoolID = TeamMascot.SchoolID;

SELECT LastName, SUM(PointsScored) FROM PlayerStats GROUP BY LastName;
