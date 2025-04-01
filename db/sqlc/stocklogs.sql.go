// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: stocklogs.sql

package db

import (
	"context"
	"database/sql"
)

const createStockLogs = `-- name: CreateStockLogs :one
INSERT INTO stocklogs(
    "from_supplier", "from_employee", "coffee", "quantity", "made_at"
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, from_supplier, from_employee, coffee, quantity, made_at
`

type CreateStockLogsParams struct {
	FromSupplier int64        `json:"from_supplier"`
	FromEmployee int64        `json:"from_employee"`
	Coffee       int64        `json:"coffee"`
	Quantity     int32        `json:"quantity"`
	MadeAt       sql.NullTime `json:"made_at"`
}

func (q *Queries) CreateStockLogs(ctx context.Context, arg CreateStockLogsParams) (Stocklog, error) {
	row := q.db.QueryRowContext(ctx, createStockLogs,
		arg.FromSupplier,
		arg.FromEmployee,
		arg.Coffee,
		arg.Quantity,
		arg.MadeAt,
	)
	var i Stocklog
	err := row.Scan(
		&i.ID,
		&i.FromSupplier,
		&i.FromEmployee,
		&i.Coffee,
		&i.Quantity,
		&i.MadeAt,
	)
	return i, err
}

const deleteStockLogs = `-- name: DeleteStockLogs :exec
DELETE FROM stocklogs WHERE id = $1
`

func (q *Queries) DeleteStockLogs(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStockLogs, id)
	return err
}

const getStockLogs = `-- name: GetStockLogs :one
SELECT id, from_supplier, from_employee, coffee, quantity, made_at FROM stocklogs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStockLogs(ctx context.Context, id int64) (Stocklog, error) {
	row := q.db.QueryRowContext(ctx, getStockLogs, id)
	var i Stocklog
	err := row.Scan(
		&i.ID,
		&i.FromSupplier,
		&i.FromEmployee,
		&i.Coffee,
		&i.Quantity,
		&i.MadeAt,
	)
	return i, err
}

const listStockLogs = `-- name: ListStockLogs :many
SELECT id, from_supplier, from_employee, coffee, quantity, made_at FROM stocklogs
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStockLogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStockLogs(ctx context.Context, arg ListStockLogsParams) ([]Stocklog, error) {
	rows, err := q.db.QueryContext(ctx, listStockLogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stocklog
	for rows.Next() {
		var i Stocklog
		if err := rows.Scan(
			&i.ID,
			&i.FromSupplier,
			&i.FromEmployee,
			&i.Coffee,
			&i.Quantity,
			&i.MadeAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStockLogs = `-- name: UpdateStockLogs :one
UPDATE stocklogs
SET
    from_supplier = $2,
    from_employee = $3,
    coffee = $4,
    quantity = $5,
    made_at = $6
WHERE id = $1
RETURNING id, from_supplier, from_employee, coffee, quantity, made_at
`

type UpdateStockLogsParams struct {
	ID           int64        `json:"id"`
	FromSupplier int64        `json:"from_supplier"`
	FromEmployee int64        `json:"from_employee"`
	Coffee       int64        `json:"coffee"`
	Quantity     int32        `json:"quantity"`
	MadeAt       sql.NullTime `json:"made_at"`
}

func (q *Queries) UpdateStockLogs(ctx context.Context, arg UpdateStockLogsParams) (Stocklog, error) {
	row := q.db.QueryRowContext(ctx, updateStockLogs,
		arg.ID,
		arg.FromSupplier,
		arg.FromEmployee,
		arg.Coffee,
		arg.Quantity,
		arg.MadeAt,
	)
	var i Stocklog
	err := row.Scan(
		&i.ID,
		&i.FromSupplier,
		&i.FromEmployee,
		&i.Coffee,
		&i.Quantity,
		&i.MadeAt,
	)
	return i, err
}
