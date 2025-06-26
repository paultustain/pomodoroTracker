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
SELECT * FROM tasks WHERE project_id = $1 ORDER BY created_at;

-- name: GetAllOpen :many 
SELECT * FROM tasks WHERE completed IS false;

-- name: GetAllTasks :many 
SELECT * FROM tasks;
