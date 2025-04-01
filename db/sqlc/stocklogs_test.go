package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStocklog(t *testing.T) Stocklog {
	supplier := createRandomSupplier(t)
	employee := createRandomEmployee(t)
	coffee := createRandomCoffee(t)

	arg := CreateStockLogsParams{
		FromSupplier: supplier.ID,
		FromEmployee: employee.ID,
		Coffee:       coffee.ID,
		Quantity:	  int32(util.RandomInt(1,20)),
		MadeAt:       util.RandomDate(),
	}

	stocklog, err := testQueries.CreateStockLogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stocklog)

	require.Equal(t, arg.FromSupplier, stocklog.FromSupplier)
	require.Equal(t, arg.FromEmployee, stocklog.FromEmployee)
	require.Equal(t, arg.Coffee, stocklog.Coffee)
	require.Equal(t, arg.Quantity, stocklog.Quantity)
	require.Equal(t, arg.MadeAt.Time.In(time.UTC), stocklog.MadeAt.Time.In(time.UTC))

	require.NotZero(t, stocklog.ID)

	return stocklog
}

func deleteIDStocklog(stocklog Stocklog) error {
	err := testQueries.DeleteStockLogs(context.Background(), stocklog.ID)
	return err
}

func TestCreateStocklog(t *testing.T) {
	stocklog := createRandomStocklog(t)

	deleteIDSupplier(func() Supplier {
		supplier, err := testQueries.GetSupplier(context.Background(), stocklog.FromSupplier)
		require.NoError(t, err)
		return supplier
	}())

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), stocklog.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), stocklog.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDStocklog(stocklog)
}

func TestGetStocklog(t *testing.T) {
	stocklog := createRandomStocklog(t)
	stocklogQ, err := testQueries.GetStockLogs(context.Background(), stocklog.ID)
	require.NoError(t, err)
	require.NotEmpty(t, stocklogQ)

	require.Equal(t, stocklogQ.FromSupplier, stocklog.FromSupplier)
	require.Equal(t, stocklogQ.FromEmployee, stocklog.FromEmployee)
	require.Equal(t, stocklogQ.Coffee, stocklog.Coffee)
	require.Equal(t, stocklogQ.Quantity, stocklog.Quantity)
	require.Equal(t, stocklogQ.MadeAt.Time.In(time.UTC), stocklog.MadeAt.Time.In(time.UTC))

	deleteIDSupplier(func() Supplier {
		supplier, err := testQueries.GetSupplier(context.Background(), stocklog.FromSupplier)
		require.NoError(t, err)
		return supplier
	}())

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), stocklog.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), stocklog.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDStocklog(stocklog)
}

func TestUpdateStocklog(t *testing.T) {
	stocklog := createRandomStocklog(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), stocklog.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), stocklog.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	employee := createRandomEmployee(t)
	coffee := createRandomCoffee(t)

	arg := UpdateStockLogsParams{
		ID:           stocklog.ID,
		FromSupplier: stocklog.FromSupplier,
		FromEmployee: employee.ID,
		Coffee:       coffee.ID,
		Quantity:	  int32(util.RandomInt(1,20)),
		MadeAt:       util.RandomDate(),
	}

	stocklogQ, err := testQueries.UpdateStockLogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stocklogQ)

	require.Equal(t, stocklogQ.FromSupplier, arg.FromSupplier)
	require.Equal(t, stocklogQ.FromEmployee, arg.FromEmployee)
	require.Equal(t, stocklogQ.Coffee, arg.Coffee)
	require.Equal(t, stocklogQ.Quantity, arg.Quantity)
	require.Equal(t, stocklogQ.MadeAt.Time.In(time.UTC), arg.MadeAt.Time.In(time.UTC))

	deleteIDSupplier(func() Supplier {
		supplier, err := testQueries.GetSupplier(context.Background(), stocklog.FromSupplier)
		require.NoError(t, err)
		return supplier
	}())

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), stocklog.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), stocklog.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDStocklog(stocklog)
}

func TestDeleteStocklog(t *testing.T) {
	stocklog := createRandomStocklog(t)

	deleteIDSupplier(func() Supplier {
		supplier, err := testQueries.GetSupplier(context.Background(), stocklog.FromSupplier)
		require.NoError(t, err)
		return supplier
	}())

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), stocklog.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), stocklog.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	err := deleteIDStocklog(stocklog)
	require.NoError(t, err)

	stocklogQ, err := testQueries.GetStockLogs(context.Background(), stocklog.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, stocklogQ)
}
