package mockaccount

import (
	"math"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/christhirst/finance/pkg/helper"
)

func MockBar(index float64, randLevel int, strength float64) marketdata.Bar {
	var bar marketdata.Bar
	bar.Close = float64(math.Sin(index)*float64(helper.Random(randLevel))*float64(strength) + 20)
	return bar

}

func CreateMockBars(l int, randomLevel int, strength float64) []marketdata.Bar {
	var bars []marketdata.Bar
	for i := 0; i < l; i++ {
		bars = append(bars, MockBar(float64(i), randomLevel, strength))
	}
	return bars
}
