// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: projects.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createProject = `-- name: CreateProject :one
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
RETURNING id, created_at, updated_at, name, time_spent, time_limit_type, time_limit, completed
`

type CreateProjectParams struct {
	Name          string
	TimeLimitType string
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, arg.Name, arg.TimeLimitType)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.TimeSpent,
		&i.TimeLimitType,
		&i.TimeLimit,
		&i.Completed,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const deleteProjects = `-- name: DeleteProjects :exec
DELETE FROM projects
`

func (q *Queries) DeleteProjects(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteProjects)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, created_at, updated_at, name, time_spent, time_limit_type, time_limit, completed FROM projects WHERE id = $1
`

func (q *Queries) GetProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.TimeSpent,
		&i.TimeLimitType,
		&i.TimeLimit,
		&i.Completed,
	)
	return i, err
}

const getProjects = `-- name: GetProjects :many
SELECT id, created_at, updated_at, name, time_spent, time_limit_type, time_limit, completed FROM projects ORDER BY created_at
`

func (q *Queries) GetProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.TimeSpent,
			&i.TimeLimitType,
			&i.TimeLimit,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTime = `-- name: UpdateTime :one
UPDATE projects
SET updated_at = NOW(), 
time_spent = $1
WHERE id=$2
RETURNING id, created_at, updated_at, name, time_spent, time_limit_type, time_limit, completed
`

type UpdateTimeParams struct {
	TimeSpent int32
	ID        uuid.UUID
}

func (q *Queries) UpdateTime(ctx context.Context, arg UpdateTimeParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, updateTime, arg.TimeSpent, arg.ID)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.TimeSpent,
		&i.TimeLimitType,
		&i.TimeLimit,
		&i.Completed,
	)
	return i, err
}
