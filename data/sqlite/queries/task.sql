-- name: ListTasks :many
SELECT id, name, done, created, type 
FROM tasks
WHERE type = ?;

-- name: CreateItem :exec
INSERT INTO tasks (
    name, done, created, type 
) VALUES (
    ?, ?, ?, ?
);

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks
SET name = ?
WHERE id = ?

--
-- CREATE TABLE task (
--     id INTEGER PRIMARY KEY NOT NULL,
--     name TEXT NOT NULL,
--     done INTEGER NULL,
--     created INTEGER NOT NULL,
--     type INTEGER NOT NULL
-- )
--