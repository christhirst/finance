package alpacaAcc

import (
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata/stream"
)

type AlpacaClientContainer struct {
	Client *marketdata.Client
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
func Initc() alpaca.Client {
	clientOp := alpaca.ClientOpts{
		ApiKey:    os.Getenv("API_KEY_ID"),
		ApiSecret: os.Getenv("SECRET_KEY"),
		BaseURL:   "https://paper-api.alpaca.markets"}

	client := alpaca.NewClient(clientOp)
	algo = alpacaClientContainer{
		tradeClient: alpaca.NewClient(alpaca.ClientOpts{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
			BaseURL:   baseURL,
		}),
		dataClient: marketdata.NewClient(marketdata.ClientOpts{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		}),
		streamClient: stream.NewStocksClient(feed,
			stream.WithCredentials(apiKey, apiSecret),
		),
		feed:          feed,
		movingAverage: movingaverage.New(windowSize),
		stock:         stock,
	}
	return client
}
