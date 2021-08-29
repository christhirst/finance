package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

type StockData struct {
	lastPrice float32
}

func GetLiveData(c *alpaca.Client, stock string) {

}

func GetHistData(c *alpaca.Client, stock string, numBars int) []alpaca.Bar {
	bar, _ := c.GetSymbolBars(stock, alpaca.ListBarParams{Timeframe: "day", Limit: &numBars})
	return bar
}
