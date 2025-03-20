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

func RandomFirstName() string {
	var name string
	rNum := 1 + r.Int63n(5-1+1)

	switch rNum {
	case 1:
		name = "Marcelo"
	case 2:
		name = "Joseph"
	case 3:
		name = "Tanaka"
	case 4:
		name = "Jane"
	case 5:
		name = "John"
	default:
		log.Fatal("ERROR. Something went wrong generating a random number")
	}

	return name
}

func RandomLastName() string {
	var name string
	rNum := 1 + r.Int63n(5-1+1)

	switch rNum {
	case 1:
		name = "Doe"
	case 2:
		name = "Souza"
	case 3:
		name = "Yamamoto"
	case 4:
		name = "Reichwald"
	case 5:
		name = "Leandoer"
	default:
		log.Fatal("ERROR. Something went wrong generating a random number")
	}

	return name
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

func PwdGen() string {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
	lower := "abcdefghijklmnopqrstuvwxyz";
	num := "0123456789";
	spec := "!@#$%^&*()-_=+<>?";

	full := upper + lower + num + spec
	var pwd string 

	pwd += string(upper[RandomInt(0, int64(len(upper) - 1))])
	pwd += string(lower[RandomInt(0, int64(len(lower) - 1))])
	pwd += string(num[RandomInt(0, int64(len(num) - 1))])
	pwd += string(spec[RandomInt(0, int64(len(spec) - 1))])

	for i := len(pwd); i < 9; i++ {
		pwd += string(full[RandomInt(0, int64(len(full)) - 1)])
	}

	return pwd
}