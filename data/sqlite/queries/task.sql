-- name: ListTasks :many
SELECT id, name, done, created, type 
FROM tasks
WHERE type = ?;

-- name: CreateTask :exec
INSERT INTO tasks (
    name, created, type, done
) VALUES (
    ?, ?, ?, null
);

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