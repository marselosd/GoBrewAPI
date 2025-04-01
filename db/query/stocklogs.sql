-- name: CreateStockLogs :one
INSERT INTO stocklogs(
    "from_supplier", "from_employee", "coffee", "quantity", "made_at"
) VALUES (
    $1, $2, $3, $4, $5
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
    from_supplier = $2,
    from_employee = $3,
    coffee = $4,
    quantity = $5,
    made_at = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStockLogs :exec
DELETE FROM stocklogs WHERE id = $1;