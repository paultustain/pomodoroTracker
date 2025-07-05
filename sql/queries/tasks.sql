-- name: CreateTask :one
INSERT INTO tasks(id, created_at, updated_at, task, completed, project_id)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(), 
	$1,
	false, 
	$2
)
RETURNING *;

-- name: GetProjectTasks :many
SELECT * FROM tasks WHERE project_id = $1 ORDER BY created_at, (completed is true) ASC;

-- name: GetAllOpen :many 
SELECT * FROM tasks WHERE completed IS false;

-- name: GetAllTasks :many 
SELECT * FROM tasks;

-- name: CompleteTask :one
UPDATE tasks
SET updated_at = NOW(),
completed = NOT completed 
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id=$1;
