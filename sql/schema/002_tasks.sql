-- +goose Up
CREATE TABLE tasks(
	id UUID PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	task TEXT NOT NULL, 
	description TEXT,
	completed BOOLEAN NOT NULL,
	project_id UUID REFERENCES projects(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tasks;
