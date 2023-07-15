package alpacaAcc

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
	"github.com/rs/zerolog/log"
)

type StockData struct {
	lastPrice float32
}

func GetLiveData(stock string) {
	apiKey := os.Getenv("API_KEY_ID")
	apiSecret := os.Getenv("SECRET_KEY")
	//baseURL := "https://paper-api.alpaca.markets"

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
		stream.WithCredentials(apiKey, apiSecret),
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

func GetHistData(Client marketdata.Client, stock string, start time.Time, end time.Time, numBars int) (map[string][]marketdata.Bar, error) {
	bar, err := Client.GetMultiBars([]string{stock}, marketdata.GetBarsRequest{Start: start, End: end})
	if err != nil {
		log.Error().Err(err).Msg("Unable to fetch data")
	}
	return bar, err
}

func Tradingdays(Client marketdata.Client, days int, min int) (int, error) {
	startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now().Add(-15*time.Minute)
	bars, err := Client.GetMultiBars([]string{"AAPL"}, marketdata.GetBarsRequest{
		Start: startTime,
		End:   endTime,
	})
	return len(bars), err
}

func GetHistDatas(Client marketdata.Client, stock string, startTime, endTime time.Time, numBars int) (map[string][]marketdata.Bar, error) {
	bars, err := Client.GetMultiBars([]string{stock}, marketdata.GetBarsRequest{
		Start: startTime,
		End:   endTime,
	})
	return bars, err
}
