package alpacaAcc

import (
	"context"
	"fmt"

	"os"
	"sync/atomic"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Msg("could not establish connection, error")

	}
	fmt.Println("established connection")

}

func GetHistData(Client marketdata.Client, stock string, startdt *time.Time, enddt *time.Time, numBars int) (map[string][]marketdata.Bar, error) {
	bar, err := Client.GetMultiBars([]string{stock}, marketdata.GetBarsParams{
		Start:      time.Date(2021, 8, 9, 13, 30, 0, 0, time.UTC),
		End:        time.Date(2022, 3, 9, 13, 30, 1, 0, time.UTC),
		TotalLimit: numBars,
	})
	if err != nil {
		log.Error().Err(err).Msg("Unable to fetch data")
	}
	//GetSymbolBars(stock, alpaca.ListBarParams{Timeframe: "day", StartDt: startdt, EndDt: enddt, Limit: &numBars})
	return bar, err
}

func Tradingdays(Client marketdata.Client, days int, min int) (int, error) {
	startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now().Add(-15*time.Minute)
	bars, err := Client.GetMultiBars([]string{"AAPL"}, marketdata.GetBarsParams{
		Start: startTime,
		End:   endTime,
	})
	return len(bars), err
}

func GetHistDatas(Client marketdata.Client, startTime, endTime time.Time) (int, error) {
	bars, err := Client.GetMultiBars([]string{"AAPL"}, marketdata.GetBarsParams{
		Start: startTime,
		End:   endTime,
	})
	return len(bars), err
}
