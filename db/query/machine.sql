-- name: CreateMachine :one
INSERT INTO machine(
    "sector", "company", "coffee_id", "last_restocked_at"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetMachine :one
SELECT * FROM machine
WHERE id = $1 LIMIT 1;

-- name: ListMachine :many
SELECT * FROM machine
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateMachine :one
UPDATE machine
SET
    sector = $2,
    company = $3,
    coffee_id = $4,
    last_restocked_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteMachine :exec
DELETE FROM machine WHERE id = $1;