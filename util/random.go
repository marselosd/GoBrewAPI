package util

import (
	"database/sql"
	"log"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandomCoffee() string {
	var coffeeString string
	rNum := 1 + r.Int63n(5-1+1)

	switch rNum {
	case 1:
		coffeeString = "Cappucino"
	case 2:
		coffeeString = "Moccachino"
	case 3:
		coffeeString = "Matcha"
	case 4:
		coffeeString = "Caramel"
	case 5:
		coffeeString = "Brew"
	default:
		log.Fatal("ERROR. Something went wrong generating a random number")
	}

	return coffeeString
}

func RandomDate() sql.NullTime {
	date := int(1 + r.Int31n(31-1+1))
	monthInt := int(1 + r.Int31n(12-1+1))
	year := int(2024 + r.Int31n(2025-2024+1))
	hours := int(1 + r.Int31n(24-1+1))
	minutes := int(1 + r.Int31n(59-1+1))
	seconds := int(1 + r.Int31n(59-1+1))

	month := time.Month(monthInt)

	return sql.NullTime{
		Time:  time.Date(year, month, date, hours, minutes, seconds, 0, time.UTC),
		Valid: true,
	}
}
