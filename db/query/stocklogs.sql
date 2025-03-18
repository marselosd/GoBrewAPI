-- name: CreateStockLogs :one
INSERT INTO stocklogs(
    "from_supplier", "from_employee", "coffee", "made_at"
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetStockLogs :one
SELECT * FROM stocklogs
WHERE id = $1 LIMIT 1;

-- name: ListStockLogs :many
SELECT * FROM stocklogs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateStockLogs :one
UPDATE stocklogs
SET
    from_supplier = CASE WHEN $2 IS NOT NULL THEN $2 ELSE from_supplier END,
    from_employee = CASE WHEN $3 IS NOT NULL THEN $3 ELSE from_employee END,
    coffee = CASE WHEN $4 IS NOT NULL THEN $4 ELSE coffee END,
    made_at = CASE WHEN $5 IS NOT NULL THEN $5 ELSE made_at END
WHERE id = $1
RETURNING *;

-- name: DeleteStockLogs :exec
DELETE FROM stocklogs WHERE id = $1;