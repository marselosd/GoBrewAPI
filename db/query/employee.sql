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
    firstname = $2,
    lastname = $3,
    password = $4,
    role = $5,
    created_at = $6,
    is_admin = $7
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee WHERE id = $1;