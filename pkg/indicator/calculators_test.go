package indicator

import (
	"math/rand"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func GenBars() []alpaca.Bar {
	var barsList []alpaca.Bar

	for i := 0; i < 2000; i++ {
		newBar := alpaca.Bar{
			Open:  rand.Float32() * 10000,
			Close: rand.Float32() * 10000,
		}
		barsList = append(barsList, newBar)
	}
	return barsList
}

func TestAvarage(t *testing.T) {
	newBarList := GenBars()
	average := Avarage(newBarList)
	if min(newBarList) > average || average > max(newBarList) {
		t.Errorf("Average %f is lower as min: %f or higher as max: %f", average, min(newBarList), max(newBarList))
	}

}
