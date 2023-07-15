package patternreconition

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

func bearish_candlestick(bar marketdata.Bar) bool {
	return (bar.Close < bar.Open)
}

func BullishEngulfingCandle(bars []marketdata.Bar, index int) bool {
	curent_day := bars[index]
	previous_day := bars[index-1]

	if bearish_candlestick(previous_day) && curent_day.Close > previous_day.Open && curent_day.Open < previous_day.Close {
		return true
	}
	return false
}
