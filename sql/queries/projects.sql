-- name: CreateProject :one
INSERT INTO projects(
	id, 
	created_at,
	updated_at, 
	name, 
	time_spent, 
	time_limit_type, 
	time_limit, 
	completed
)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(), 
	$1, 
	0, 
	$2,
	0,
	false
)
RETURNING *;

-- name: DeleteProjects :exec
DELETE FROM projects;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;

-- name: GetProjects :many 
SELECT * FROM projects ORDER BY created_at; 

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;

-- name: UpdateTime :one
UPDATE projects
SET updated_at = NOW(), 
time_spent = $1
WHERE id=$2
RETURNING *;
