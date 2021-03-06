package alpacaAcc

import (
	"os"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type AlpacaClientContainer struct {
	TradeClient   alpaca.Client
	DataClient    marketdata.Client
	StreamClient  stream.StocksClient
	feed          string
	movingAverage *movingaverage.MovingAverage
	lastOrder     string
	stock         string
}

/* func tradeUpdateHandler(update alpaca.TradeUpdate) {
	fmt.Println("trade update", update)
}

func tradeHandler(trade stream.Trade) {
	fmt.Println("trade", trade)
}

func quoteHandler(quote stream.Quote) {
	fmt.Println("quote", quote)
}

func barHandler(bar stream.Bar) {
	fmt.Println("bar", bar)
}
*/
func Init() marketdata.Client {

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_KEY_ID"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("SECRET_KEY"))
	clientOp := marketdata.ClientOpts{ApiKey: os.Getenv("API_KEY_ID"), ApiSecret: os.Getenv("SECRET_KEY")}
	Client := marketdata.NewClient(clientOp)
	return Client
}
func Initc() AlpacaClientContainer {

	// You can set your API key/secret here or you can use environment variables!
	apiKey := os.Getenv("API_KEY_ID")
	apiSecret := os.Getenv("SECRET_KEY")
	// Change baseURL to https://paper-api.alpaca.markets if you want use paper!
	baseURL := "https://paper-api.alpaca.markets"
	// Change feed to sip if you have proper subscription
	feed := "iex"

	algo := AlpacaClientContainer{
		TradeClient: alpaca.NewClient(alpaca.ClientOpts{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
			BaseURL:   baseURL,
		}),
		DataClient: marketdata.NewClient(marketdata.ClientOpts{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		}),
		StreamClient: stream.NewStocksClient(feed,
			stream.WithCredentials(apiKey, apiSecret),
		),
		feed: feed,
	}

	return algo
}
