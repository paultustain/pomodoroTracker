-- +goose Up
CREATE TABLE projects(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL UNIQUE,
	completed BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE projects; 

