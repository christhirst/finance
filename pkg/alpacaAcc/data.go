package alpacaAcc

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type StockData struct {
	lastPrice float32
}

func GetLiveData(stock string) {
	ii := os.Getenv("API_KEY_ID")
	oo := os.Getenv("SECRET_KEY")
	var tradeCount, quoteCount, barCount int32
	// modify these according to your needs
	tradeHandler := func(t stream.Trade) {
		atomic.AddInt32(&tradeCount, 1)
	}
	quoteHandler := func(q stream.Quote) {
		atomic.AddInt32(&quoteCount, 1)
	}
	barHandler := func(b stream.Bar) {
		atomic.AddInt32(&barCount, 1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Creating a client that connexts to iex
	c := stream.NewStocksClient(
		"iex",
		// configuring initial subscriptions and handlers
		stream.WithTrades(tradeHandler, "SPY"),
		stream.WithQuotes(quoteHandler, "AAPL", "SPY"),
		stream.WithBars(barHandler, "AAPL", "SPY"),
		stream.WithCredentials(ii, oo),

		// use stream.WithDailyBars to subscribe to daily bars too
		// use stream.WithCredentials to manually override envvars
		// use stream.WithHost to manually override envvar
		// use stream.WithLogger to use your own logger (i.e. zap, logrus) instead of log
		// use stream.WithProcessors to use multiple processing gourotines
		// use stream.WithBufferSize to change buffer size
		// use stream.WithReconnectSettings to change reconnect settings
	)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("trades:", tradeCount, "quotes:", quoteCount, "bars:", barCount)
		}
	}()

	if err := c.Connect(ctx); err != nil {
		log.Fatalf("could not establish connection, error: %s", err)
	}
	fmt.Println("established connection")

}

func GetHistData(c *alpaca.Client, stock string, startdt *time.Time, enddt *time.Time, numBars int) ([]alpaca.Bar, error) {
	bar, err := c.GetSymbolBars(stock, alpaca.ListBarParams{Timeframe: "day", StartDt: startdt, EndDt: enddt, Limit: &numBars})
	return bar, err
}

func Tradingdays(Client *alpaca.Client, days int) (int, error) {
	startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now()
	bars, err := GetHistData(Client, "AAPL", &startTime, &endTime, days)
	return len(bars), err
}
