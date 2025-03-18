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
    from_employee = CASE WHEN $2 IS NOT NULL THEN $2 ELSE from_employee END,
    coffee = CASE WHEN $3 IS NOT NULL THEN $3 ELSE coffee END,
    made_at = CASE WHEN $4 IS NOT NULL THEN $4 ELSE made_at END
WHERE id = $1
RETURNING *;

-- name: DeleteLogs :exec
DELETE FROM logs WHERE id = $1;