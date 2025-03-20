package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomLog(t *testing.T) Log {
	employee := createRandomEmployee(t)
	coffee := createRandomCoffee(t)

	arg := CreateLogsParams{
		FromEmployee: employee.ID,
		Coffee:  coffee.ID,
		MadeAt:  util.RandomDate(),
	}

	logs, err := testQueries.CreateLogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, logs)

	require.Equal(t, arg.FromEmployee, logs.FromEmployee)
	require.Equal(t, arg.Coffee, logs.Coffee)
	require.Equal(t, arg.MadeAt.Time.In(time.UTC), logs.MadeAt.Time.In(time.UTC))

	require.NotZero(t, logs.ID)

	return logs
}

func deleteIDLog(log Log) error {
	err := testQueries.DeleteLogs(context.Background(), log.ID)
	return err
}

func TestCreateLog(t *testing.T) {
	log := createRandomLog(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), log.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), log.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDLog(log)
}

func TestGetLog(t *testing.T) {
	log := createRandomLog(t)
	logQ, err := testQueries.GetLogs(context.Background(), log.ID)
	require.NoError(t, err)
	require.NotEmpty(t, logQ)

	require.Equal(t, logQ.FromEmployee, log.FromEmployee)
	require.Equal(t, logQ.Coffee, log.Coffee)
	require.Equal(t, logQ.MadeAt.Time.In(time.UTC), log.MadeAt.Time.In(time.UTC))
	
	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), log.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), log.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDLog(logQ)
}

func TestUpdateLog(t *testing.T) {
	log := createRandomLog(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), log.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), log.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	employee := createRandomEmployee(t)
	coffee := createRandomCoffee(t)

	arg := UpdateLogsParams{
		ID: log.ID,
		FromEmployee: employee.ID,
		Coffee:  coffee.ID,
		MadeAt:  util.RandomDate(),
	}

	logQ, err := testQueries.UpdateLogs(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, logQ)
	require.Equal(t, logQ.Coffee, arg.Coffee)
	require.Equal(t, logQ.FromEmployee, arg.FromEmployee)
	require.Equal(t, logQ.MadeAt.Time.In(time.UTC), arg.MadeAt.Time.In(time.UTC))

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), arg.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), arg.FromEmployee)
		require.NoError(t, err)
		return employee
	}())

	deleteIDLog(logQ)
}

func TestDeleteLog(t *testing.T) {
	log := createRandomLog(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), log.Coffee)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDEmployee(func() Employee {
		employee, err := testQueries.GetEmployee(context.Background(), log.FromEmployee)
		require.NoError(t, err)
		return employee
	}())
	
	err := deleteIDLog(log)
	require.NoError(t, err)

	logQ, err := testQueries.GetLogs(context.Background(), log.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, logQ)
}
