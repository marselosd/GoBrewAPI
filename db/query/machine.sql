-- name: CreateMachine :one
INSERT INTO machine(
    "sector", "company", "coffee_id", "quantity", "last_restocked_at"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMachine :one
SELECT * FROM machine
WHERE id = $1 LIMIT 1;

-- name: GetMachineForUpdate :one
SELECT * FROM machine
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

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
    quantity = $5,
    last_restocked_at = $6
WHERE id = $1
RETURNING *;

-- name: UpdateMachineQuantity :one
UPDATE machine
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMachine :exec
DELETE FROM machine WHERE id = $1;

-- name: AddQuantity :one
UPDATE machine
SET quantity = quantity + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;