package alpacaAcc

import (
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata/stream"
)

type bucket struct {
	list        []string
	qty         int
	adjustedQty int
	equityAmt   float64
}

type stockField struct {
	name string
	pc   float64
}

type AlpacaClientContainer struct {
	TradeClient  *alpaca.Client
	DataClient   *marketdata.Client
	StreamClient *stream.StocksClient
	long         bucket
	short        bucket
	allStocks    []stockField
	blacklist    []string
	feed         string
	//movingAverage *movingaverage.MovingAverage
	//lastOrder     string
	//stock         string
}

func Init() AlpacaClientContainer {
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
	}

	return algo
}
