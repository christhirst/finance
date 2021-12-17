package helper

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

func random(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}

func RandomInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomString(count int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, count)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomDeci(i int) decimal.Decimal {
	return decimal.NewFromInt(int64(random(i)))
}
