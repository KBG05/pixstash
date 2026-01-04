-- +goose up 
ALTER TABLE users
     ALTER COLUMN password TYPE VARCHAR;

-- +goose down
ALTER TABLE users
     ALTER COLUMN password TYPE VARCHAR(32);
