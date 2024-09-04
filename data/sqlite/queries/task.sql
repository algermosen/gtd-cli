-- name: ListTasks :many
SELECT id, name, done, created, type 
FROM tasks
WHERE type = ?;

-- name: CreateTask :one
INSERT INTO tasks (
    name, done, created, type 
) VALUES (
    ?, ?, ?, ?
) RETURNING id;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;

-- name: UpdateTask :exec
UPDATE tasks
SET name = ?
WHERE id = ?;

-- name: CheckTask :exec
UPDATE tasks
SET done = DATE()
WHERE id = ?;

-- name: CheckAllTasks :exec
UPDATE tasks
SET done = DATE()
WHERE done is null;

-- name: DeleteAllTasks :exec
DELETE FROM tasks;