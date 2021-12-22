package mockaccount

import (
	"math"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/helper"
)

func MockBar(index float64, randLevel int, strength float64) alpaca.Bar {
	var bar alpaca.Bar
	bar.Close = float32(math.Sin(index)*float64(helper.Random(randLevel))*float64(strength) + 20)
	return bar

}

func CreateMockBars(l int, randomLevel int, strength float64) []alpaca.Bar {
	var bars []alpaca.Bar
	for i := 0; i < l; i++ {
		bars = append(bars, MockBar(float64(i), randomLevel, strength))
	}
	return bars
}
