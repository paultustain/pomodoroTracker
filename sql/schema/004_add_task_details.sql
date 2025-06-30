-- +goose Up 
ALTER TABLE tasks 
	ADD COLUMN description TEXT, 
	ADD COLUMN time_limit_type TEXT,
	ADD COLUMN time_limit INT DEFAULT 0; 

-- +goose Down
ALTER TABLE tasks
	DROP COLUMN description, 
	DROP COLUMN time_limit_type, 
	DROP COLUMN time_limit;

