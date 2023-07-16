package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
	"github.com/rs/zerolog/log"
)

type bucket struct {
	symbol      string
	qty         int
	adjustedQty int
	equityAmt   float64
}

type stockField struct {
	name string
	pc   float64
}

type AlpacaClientContainer struct {
	TradeClient  TradeClient
	DataClient   *marketdata.Client
	StreamClient *stream.StocksClient
	long         map[string]*bucket
	short        map[string]*bucket
	allStocks    *[]stockField
	blacklist    *[]string
	feed         string
	//movingAverage *movingaverage.MovingAverage
	//lastOrder     string
	//stock         string
}

type TradeClient interface {
	PlaceOrder(req alpaca.PlaceOrderRequest) (*alpaca.Order, error)
	GetAccount() (*alpaca.Account, error)
	GetAsset(symbol string) (*alpaca.Asset, error)
	GetCalendar(req alpaca.GetCalendarRequest) ([]alpaca.CalendarDay, error)
}

func Init() AlpacaClientContainer {
	//aa := alpaca.Client{}
	//aa.GetCalendar()
	// You can set your API key/secret here or you can use environment variables!
	apiKey := "PK7358TZCGIMNCEJNQQS"                        //os.Getenv("API_KEY_ID")
	apiSecret := "mS3QY46BcAHyKQWWoubSoSCpoeMY3zWEjWCj0p2K" //os.Getenv("SECRET_KEY")
	if apiKey == "" || apiSecret == "" {
		log.Panic()
	}

	// Change baseURL to https://paper-api.alpaca.markets if you want use paper!
	baseURL := "https://paper-api.alpaca.markets"
	// Change feed to sip if you have proper subscription
	feed := "iex"

	algo := AlpacaClientContainer{
		TradeClient: alpaca.NewClient(alpaca.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
			BaseURL:   baseURL,
		}),
		DataClient: marketdata.NewClient(marketdata.ClientOpts{
			APIKey:    apiKey,
			APISecret: apiSecret,
		}),
		StreamClient: stream.NewStocksClient(feed,
			stream.WithCredentials(apiKey, apiSecret),
		),
		feed: feed,
		long: make(map[string]*bucket),
	}
	algo.initStocks("AAPL")

	return algo
}

func (acc AlpacaClientContainer) initStocks(sym string) {
	acc.long[sym] = &bucket{symbol: sym, qty: 0, adjustedQty: 0, equityAmt: 0}
}

func (acc AlpacaClientContainer) valuePositions(sym string, price float64) map[string]float64 {
	//req := marketdata.GetLatestQuoteRequest{Feed: acc.feed, Currency: "USD"}

	positions := map[string]float64{}
	for i, v := range acc.long {
		//lq, err := acc.DataClient.GetLatestQuote("AAPL", req)
		//fmt.Println(err)
		//fmt.Println(lq)
		/* fmt.Println(v.symbol)
		if err != nil {
			log.Error().Err(err).Msg("")
		} */
		fmt.Println(float64(v.qty) * v.equityAmt)
		fmt.Println(float64(v.qty) * price)

		positions[i] = float64(v.qty)*v.equityAmt - float64(v.qty)*price
	}

	return positions
}
