-- name: CreateLogs :one
INSERT INTO logs(
    "from_employee", "coffee", "made_at"
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetLogs :one
SELECT * FROM logs
WHERE id = $1 LIMIT 1;

-- name: ListLogs :many
SELECT * FROM logs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateLogs :one
UPDATE logs
SET
    from_employee = $2,
    coffee = $3,
    made_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteLogs :exec
DELETE FROM logs WHERE id = $1;