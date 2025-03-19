package db

import (
	"GoBrewAPI/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateCoffee(t *testing.T) {

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

}