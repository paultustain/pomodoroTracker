-- +goose Up
CREATE TABLE projects(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL UNIQUE,
	time_spent INT NOT NULL,
	time_limit_type TEXT NOT NULL,
	time_limit INT NOT NULL DEFAULT 0,
	completed BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE projects; 

