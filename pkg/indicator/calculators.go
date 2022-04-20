package indicator

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
)

func Avarage(bars []marketdata.Bar) float64 {
	var sum float64
	for _, c := range bars {
		sum += c.Close
	}

	return (sum / float64(len(bars)))

}

func min(bars []marketdata.Bar) float64 {
	lowest := bars[0].Close
	for _, running := range bars {
		if running.Close < lowest {
			lowest = running.Close
		}

	}
	return lowest
}

func max(bars []alpaca.Bar) float32 {
	highest := bars[0].Close
	for _, running := range bars {
		if running.Close > highest {
			highest = running.Close
		}

	}
	return highest
}
