package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomMachine(t *testing.T) Machine {
	coffee := createRandomCoffee(t)

	arg := CreateMachineParams{
		Sector:          "Main",
		Company:         "Kaisha",
		CoffeeID:        coffee.ID,
		Quantity:        sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		LastRestockedAt: util.RandomDate(),
	}

	machine, err := testQueries.CreateMachine(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, machine)

	require.Equal(t, arg.Sector, machine.Sector)
	require.Equal(t, arg.Company, machine.Company)
	require.Equal(t, arg.CoffeeID, machine.CoffeeID)
	require.Equal(t, arg.LastRestockedAt.Time.In(time.UTC), machine.LastRestockedAt.Time.In(time.UTC))

	require.NotZero(t, machine.ID)

	return machine
}

func deleteIDMachine(machine Machine) error {
	err := testQueries.DeleteMachine(context.Background(), machine.ID)
	return err
}

func TestCreateMachine(t *testing.T) {
	machine := createRandomMachine(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), machine.CoffeeID)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDMachine(machine)
}

func TestGetMachine(t *testing.T) {
	machine := createRandomMachine(t)
	machineQ, err := testQueries.GetMachine(context.Background(), machine.ID)
	require.NoError(t, err)
	require.NotEmpty(t, machineQ)

	require.Equal(t, machineQ.Sector, machine.Sector)
	require.Equal(t, machineQ.Company, machine.Company)
	require.Equal(t, machineQ.CoffeeID, machine.CoffeeID)
	require.Equal(t, machineQ.Quantity, machine.Quantity)
	require.Equal(t, machineQ.LastRestockedAt.Time.In(time.UTC), machine.LastRestockedAt.Time.In(time.UTC))

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), machine.CoffeeID)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDMachine(machineQ)
}

func TestUpdateMachine(t *testing.T) {
	machine := createRandomMachine(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), machine.CoffeeID)
		require.NoError(t, err)
		return coffee
	}())

	coffee := createRandomCoffee(t)

	arg := UpdateMachineParams{
		ID:              machine.ID,
		Sector:          "HR",
		Company:         machine.Company,
		CoffeeID:        coffee.ID,
		LastRestockedAt: util.RandomDate(),
	}

	machineQ, err := testQueries.UpdateMachine(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, machineQ)

	require.Equal(t, machineQ.Sector, arg.Sector)
	require.Equal(t, machineQ.Company, arg.Company)
	require.Equal(t, machineQ.CoffeeID, arg.CoffeeID)
	require.Equal(t, machineQ.LastRestockedAt.Time.In(time.UTC), arg.LastRestockedAt.Time.In(time.UTC))

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), machine.CoffeeID)
		require.NoError(t, err)
		return coffee
	}())

	deleteIDMachine(machineQ)
}

func TestDeleteMachine(t *testing.T) {
	machine := createRandomMachine(t)

	deleteIDCoffee(func() Coffee {
		coffee, err := testQueries.GetCoffee(context.Background(), machine.CoffeeID)
		require.NoError(t, err)
		return coffee
	}())

	err := deleteIDMachine(machine)
	require.NoError(t, err)

	machineQ, err := testQueries.GetMachine(context.Background(), machine.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, machineQ)
}
