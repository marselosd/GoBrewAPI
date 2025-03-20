-- name: CreateSupplier :one
INSERT INTO supplier(
    "name", "company", "password", "created_at"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSupplier :one
SELECT * FROM supplier
WHERE id = $1 LIMIT 1;

-- name: ListSupplier :many
SELECT * FROM supplier
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSupplier :one
UPDATE supplier
SET
    name = $2,
    company = $3,
    password = $4,
    created_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM supplier WHERE id = $1;