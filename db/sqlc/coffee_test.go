package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomCoffee(t *testing.T) Coffee{
	arg := CreateCoffeeParams{
		Type: util.RandomCoffee(),
		Quantity: int32(util.RandomInt(1,10)),
		BuyedAt: util.RandomDate(),
		StockedAt: util.RandomDate(),
		IsOutstocked: false,
	}

	coffee, err := testQueries.CreateCoffee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coffee)

	require.Equal(t, arg.Type, coffee.Type)
	require.Equal(t, arg.Quantity, coffee.Quantity)
	require.Equal(t, arg.BuyedAt.Time.In(time.UTC), coffee.BuyedAt.Time.In(time.UTC))
	require.Equal(t, arg.StockedAt.Time.In(time.UTC), coffee.StockedAt.Time.In(time.UTC))
	require.Equal(t, arg.IsOutstocked, coffee.IsOutstocked)

	require.NotZero(t, coffee.ID)

	return coffee
}

func deleteIDCoffee(coffee Coffee) error{
	err := testQueries.DeleteCoffee(context.Background(), coffee.ID)

	return err
}

func TestCreateCoffee(t *testing.T) {
	coffee := createRandomCoffee(t)
	deleteIDCoffee(coffee)
}

func TestGetCoffee(t *testing.T) {
	coffee := createRandomCoffee(t)
	coffeeQ, err := testQueries.GetCoffee(context.Background(), coffee.ID)
	require.NoError(t, err)
	require.NotEmpty(t, coffeeQ)

	require.Equal(t, coffee.Type, coffeeQ.Type)
	require.Equal(t, coffee.Quantity, coffeeQ.Quantity)
	require.Equal(t, coffee.BuyedAt.Time.In(time.UTC), coffeeQ.BuyedAt.Time.In(time.UTC))
	require.Equal(t, coffee.StockedAt.Time.In(time.UTC), coffeeQ.StockedAt.Time.In(time.UTC))
	require.Equal(t, coffee.IsOutstocked, coffeeQ.IsOutstocked)

	deleteIDCoffee(coffee)
}

func TestUpdateCoffee(t *testing.T) {
	coffee := createRandomCoffee(t)

	arg := UpdateCoffeeParams{
		ID: coffee.ID,
		Type: coffee.Type,
		Quantity: int32(util.RandomInt(1,30)),
		BuyedAt: coffee.BuyedAt,
		StockedAt: util.RandomDate(),
		IsOutstocked: false,
	}

	coffeeQ, err := testQueries.UpdateCoffee(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, coffeeQ)
	require.Equal(t, arg.Quantity, coffeeQ.Quantity)
	require.Equal(t, arg.BuyedAt.Time.In(time.UTC), coffeeQ.BuyedAt.Time.In(time.UTC))
	require.Equal(t, arg.StockedAt.Time.In(time.UTC), coffeeQ.StockedAt.Time.In(time.UTC))
	require.Equal(t, arg.IsOutstocked, coffeeQ.IsOutstocked)

	deleteIDCoffee(coffee)
}

func TestDeleteCoffee(t *testing.T) {
	coffee := createRandomCoffee(t)
	err := deleteIDCoffee(coffee)
	require.NoError(t, err)

	coffeeQ, err := testQueries.GetCoffee(context.Background(), coffee.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, coffeeQ)
}