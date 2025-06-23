-- +goose Up
ALTER TABLE projects ADD COLUMN time_spent INT NOT NULL;

-- +goose Down 
ALTER TABLE projects DROP COLUMN time_spent;
