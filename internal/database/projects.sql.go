// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: projects.sql

package database

import (
	"context"
)

const createProject = `-- name: CreateProject :one
INSERT INTO projects(id, created_at, updated_at, name, completed)
VALUES (
	gen_random_uuid(),
	NOW(), 
	NOW(), 
	$1, 
	false
)
RETURNING id, created_at, updated_at, name, completed
`

func (q *Queries) CreateProject(ctx context.Context, name string) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, name)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Completed,
	)
	return i, err
}
