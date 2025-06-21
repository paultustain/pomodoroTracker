-- name: CreateProject :one
INSERT INTO projects(id, created_at, updated_at, name, completed)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(), 
	$1, 
	false
)
RETURNING *;

-- name: DeleteProjects :exec
DELETE FROM projects;
