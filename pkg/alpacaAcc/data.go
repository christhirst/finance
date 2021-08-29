package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

type StockData struct {
	lastPrice float32
}

func GetLiveData(c *alpaca.Client, stock string) {

}

func GetHistData(c *alpaca.Client, stock string) []alpaca.Bar {
	numBars := 2
	bar, _ := c.GetSymbolBars(stock, alpaca.ListBarParams{Timeframe: "minute", Limit: &numBars})
	fmt.Println(bar)
	return bar
}
