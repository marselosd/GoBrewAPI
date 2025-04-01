-- name: CreateMachineLogs :one
INSERT INTO machinelogs(
    "from_employee", "to_machine", "coffee", "quantity", "made_at"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMachineLogs :one
SELECT * FROM machinelogs
WHERE id = $1 LIMIT 1;

-- name: ListMachineLogs :many
SELECT * FROM machinelogs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateMachineLogs :one
UPDATE machinelogs
SET
    from_employee = $2,
    to_machine = $3,
    coffee = $4,
    quantity = $5,
    made_at = $6
WHERE id = $1
RETURNING *;

-- name: DeleteMachineLogs :exec
DELETE FROM machinelogs WHERE id = $1;