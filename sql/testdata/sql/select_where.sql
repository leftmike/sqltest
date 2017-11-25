--
-- Test SELECT ... WHERE w/out FROM
--
-- mysql doesn't like this syntax
-- {{if eq Dialect "mysql"}}{{Skip}}{{end}}

SELECT 1 AS c1, 2 AS c2 WHERE 1 = 1;

SELECT 1 AS c1, 2 AS c2 WHERE 0 = 1;

SELECT 1 + 2 - 3 as c1, 12 / 4 * 3 as c2;

SELECT 1 - 2 + 3 as c1, 12 * 4 / 3 as c2;

SELECT 300 * 2 + 3 / - 4 as c1, 2 + 3 * 4 as c2, 2 * 3 + 4 as c3, - 2 * 3 as c4, - (2 * 3) as c5;
