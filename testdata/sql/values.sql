--
-- Test VALUES
--
-- {{if eq Dialect "sqlite3"}}{{Skip}}{{end}}

VALUES (1, 2, 3), (4, NULL, 5);

VALUES (1.2, 3), (4, 5.6);

{{Fail .Test}}
VALUES (1, 2), ('three', 4);

{{Fail .Test}}
VALUES ('one', 2), (3, 4);

{{Fail .Test}}
VALUES (1, 'two'), (3, 4);

{{Fail .Test}}
VALUES (1, 2), (3, 'four');

