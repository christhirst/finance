package indicator

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func Avarage(bars []alpaca.Bar) float32 {
	var sum float32
	var timesum int
	for t, c := range bars {
		sum += c.Close
		timesum += t
	}

	return (sum / float32(timesum+1))

}

func min(bars []alpaca.Bar) float32 {
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
