-- name: CreateProject :one
INSERT INTO projects(id, created_at, updated_at, name, completed, time_spent)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(), 
	$1, 
	false, 
	0
)
RETURNING *;

-- name: DeleteProjects :exec
DELETE FROM projects;

-- name: DeleteProject :exec
DELETE FROM projects WHERE name = $1;

-- name: GetProjects :many 
SELECT * FROM projects; 

-- name: GetProject :one
SELECT * FROM projects WHERE name = $1;

-- name: UpdateTime :one
UPDATE projects
SET update_at = NOW(), 
time_spent = $1
WHERE name=$2
RETURNING *;
