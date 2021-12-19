package alpacaAcc

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

type AlpacaClientContainer struct {
	Client *alpaca.Client
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
func Init() *alpaca.Client {
	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_KEY_ID"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("SECRET_KEY"))
	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)
	Client := alpaca.NewClient(common.Credentials())
	return Client
}
