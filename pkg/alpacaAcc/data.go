package alpacaAcc

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

type StockData struct {
	lastPrice float32
}

func GetLiveData(c *alpaca.Client, stock string) {

}

func GetHistData(c *alpaca.Client, stock string, startdt *time.Time, enddt *time.Time, numBars int) []alpaca.Bar {
	bar, _ := c.GetSymbolBars(stock, alpaca.ListBarParams{Timeframe: "day", StartDt: startdt, EndDt: enddt, Limit: &numBars})
	return bar
}

func Tradingdays(Client *alpaca.Client, days int) int {

	startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now()
	bars := GetHistData(Client, "AAPL", &startTime, &endTime, days)

	return len(bars)
}