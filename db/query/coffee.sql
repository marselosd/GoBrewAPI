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

-- name: UpdateCoffee :one
UPDATE coffee
SET
    quantity = CASE WHEN $2 IS NOT NULL THEN $2 ELSE quantity END,
    buyed_at = CASE WHEN $3 IS NOT NULL THEN $3 ELSE buyed_at END,
    stocked_at = CASE WHEN $4 IS NOT NULL THEN $4 ELSE stocked_at END,
    is_outstocked = CASE WHEN $5 IS NOT NULL THEN $5 ELSE is_outstocked END
WHERE id = $1
RETURNING *;

-- name: DeleteCoffee :exec
DELETE FROM coffee WHERE id = $1;