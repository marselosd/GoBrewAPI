-- name: CreateEmployee :one
INSERT INTO employee(
    "firstname", "lastname", "password", "role", "created_at", "is_admin"
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE id = $1 LIMIT 1;

-- name: ListEmployee :many
SELECT * FROM employee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEmployee :one
UPDATE employee
SET
    firstname = CASE WHEN $2 IS NOT NULL THEN $2 ELSE firstname END,
    lastname = CASE WHEN $3 IS NOT NULL THEN $3 ELSE lastname END,
    password = CASE WHEN $4 IS NOT NULL THEN $4 ELSE password END,
    role = CASE WHEN $5 IS NOT NULL THEN $5 ELSE role END,
    created_at = CASE WHEN $6 IS NOT NULL THEN $6 ELSE created_at END,
    is_admin = CASE WHEN $7 IS NOT NULL THEN $7 ELSE is_admin END
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee WHERE id = $1;