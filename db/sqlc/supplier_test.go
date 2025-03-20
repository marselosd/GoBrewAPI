package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomSupplier(t *testing.T) Supplier {
	arg := CreateSupplierParams{
		Name: util.RandomFirstName(),
		Company:  "Amazon",
		Password:  util.PwdGen(),
		CreatedAt: util.RandomDate(),
	}

	supplier, err := testQueries.CreateSupplier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, supplier)

	require.Equal(t, arg.Name, supplier.Name)
	require.Equal(t, arg.Company, supplier.Company)
	require.Equal(t, arg.Password, supplier.Password)
	require.Equal(t, arg.CreatedAt.Time.In(time.UTC), supplier.CreatedAt.Time.In(time.UTC))

	require.NotZero(t, supplier.ID)

	return supplier
}

func deleteIDSupplier(supplier Supplier) error {
	err := testQueries.DeleteSupplier(context.Background(), supplier.ID)
	return err
}

func TestCreateSupplier(t *testing.T) {
	supplier := createRandomSupplier(t)
	deleteIDSupplier(supplier)
}

func TestGetSupplier(t *testing.T) {
	supplier := createRandomSupplier(t)
	supplierQ, err := testQueries.GetSupplier(context.Background(), supplier.ID)
	require.NoError(t, err)
	require.NotEmpty(t, supplierQ)

	require.Equal(t, supplierQ.Name, supplier.Name)
	require.Equal(t, supplierQ.Company, supplier.Company)
	require.Equal(t, supplierQ.Password, supplier.Password)
	require.Equal(t, supplierQ.CreatedAt.Time.In(time.UTC), supplier.CreatedAt.Time.In(time.UTC))
	
	deleteIDSupplier(supplierQ)
}

func TestUpdateSupplier(t *testing.T) {
	supplier := createRandomSupplier(t)

	arg := UpdateSupplierParams{
		ID: supplier.ID,
		Name: supplier.Name,
		Company:  supplier.Company,
		Password:  util.PwdGen(),
		CreatedAt: supplier.CreatedAt,
	}

	supplierQ, err := testQueries.UpdateSupplier(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, supplierQ)

	require.Equal(t, supplierQ.Name, arg.Name)
	require.Equal(t, supplierQ.Company, arg.Company)
	require.Equal(t, supplierQ.Password, arg.Password)
	require.Equal(t, supplierQ.CreatedAt.Time.In(time.UTC), arg.CreatedAt.Time.In(time.UTC))

	deleteIDSupplier(supplierQ)
}

func TestDeleteSupplier(t *testing.T) {
	supplier := createRandomSupplier(t)
	err := deleteIDSupplier(supplier)
	require.NoError(t, err)

	supplierQ, err := testQueries.GetSupplier(context.Background(), supplier.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, supplierQ)
}
