// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql"
)

type Coffee struct {
	ID           int64         `json:"id"`
	Type         string        `json:"type"`
	Quantity     sql.NullInt32 `json:"quantity"`
	BuyedAt      sql.NullTime  `json:"buyed_at"`
	StockedAt    sql.NullTime  `json:"stocked_at"`
	IsOutstocked bool          `json:"is_outstocked"`
}

type Employee struct {
	ID        int64        `json:"id"`
	Firstname string       `json:"firstname"`
	Lastname  string       `json:"lastname"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
	IsAdmin   bool         `json:"is_admin"`
}

type Log struct {
	ID           int64        `json:"id"`
	FromEmployee int64        `json:"from_employee"`
	Coffee       int64        `json:"coffee"`
	MadeAt       sql.NullTime `json:"made_at"`
}

type Machine struct {
	ID              int64         `json:"id"`
	Sector          string        `json:"sector"`
	Company         string        `json:"company"`
	CoffeeID        int64         `json:"coffee_id"`
	Quantity        sql.NullInt32 `json:"quantity"`
	LastRestockedAt sql.NullTime  `json:"last_restocked_at"`
}

type Machinelog struct {
	ID           int64        `json:"id"`
	FromEmployee int64        `json:"from_employee"`
	ToMachine    int64        `json:"to_machine"`
	Coffee       int64        `json:"coffee"`
	Quantity     int32        `json:"quantity"`
	MadeAt       sql.NullTime `json:"made_at"`
}

type Stocklog struct {
	ID           int64        `json:"id"`
	FromSupplier int64        `json:"from_supplier"`
	FromEmployee int64        `json:"from_employee"`
	Coffee       int64        `json:"coffee"`
	Quantity     int32        `json:"quantity"`
	MadeAt       sql.NullTime `json:"made_at"`
}

type Supplier struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Company   string       `json:"company"`
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
}
