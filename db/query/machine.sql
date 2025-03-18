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
    sector = CASE WHEN $2 IS NOT NULL THEN $2 ELSE sector END,
    company = CASE WHEN $3 IS NOT NULL THEN $3 ELSE company END,
    coffee_id = CASE WHEN $4 IS NOT NULL THEN $4 ELSE coffee_id END,
    last_restocked_at = CASE WHEN $5 IS NOT NULL THEN $5 ELSE last_restocked_at END
WHERE id = $1
RETURNING *;

-- name: DeleteMachine :exec
DELETE FROM machine WHERE id = $1;