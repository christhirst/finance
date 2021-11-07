package helper

import "math/rand"

func random(i int) int {

	return rand.Intn(i)
}

func RandomInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
