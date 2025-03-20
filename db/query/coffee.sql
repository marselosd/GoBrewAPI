-- name: CreateCoffee :one
INSERT INTO coffee(
    "type", "quantity", "buyed_at", "stocked_at", "is_outstocked"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCoffee :one
SELECT * FROM coffee
WHERE id = $1 LIMIT 1;

-- name: ListCoffee :many
SELECT * FROM coffee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateQuantityCoffee :one
UPDATE coffee
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCoffee :exec
DELETE FROM coffee WHERE id = $1;