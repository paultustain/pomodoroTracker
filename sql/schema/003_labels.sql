-- +goose Up 
CREATE TABLE labels(
	id UUID PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL, 
	label TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE labels;
