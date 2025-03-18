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
    name = CASE WHEN $2 IS NOT NULL THEN $2 ELSE name END,
    company = CASE WHEN $3 IS NOT NULL THEN $3 ELSE company END,
    password = CASE WHEN $4 IS NOT NULL THEN $4 ELSE password END,
    created_at = CASE WHEN $5 IS NOT NULL THEN $5 ELSE created_at END
WHERE id = $1
RETURNING *;

-- name: DeleteSupplier :exec
DELETE FROM supplier WHERE id = $1;