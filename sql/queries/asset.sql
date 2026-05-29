-- name: CreateAsset :one
INSERT INTO assets (
    user_id, original_name, file_hash, file_size, storage_path, status
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetAsset :one
SELECT * FROM assets
WHERE id=$1;
