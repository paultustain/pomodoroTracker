-- name: CreateTask :one
INSERT INTO tasks(id, created_at, updated_at, task, completed, project_id)
VALUES (
	$1, 
	NOW(), 
	NOW(), 
	$2,
	$3, 
	$4
)
RETURNING *;

-- name: GetTasks :many
SELECT * FROM tasks WHERE project_id = $1;

-- name: GetProjectTasks :many
SELECT * FROM tasks;

-- name: GetAllOpen :many 
SELECT * FROM tasks WHERE completed IS false;

-- name: GetAllTasks :many 
SELECT * FROM tasks;
