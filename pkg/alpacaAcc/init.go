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

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_Key_ID"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("Secret_Key"))
	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	return alpaca.NewClient(common.Credentials())
}

/*
acct, err := alpacaClient.GetAccount()
if err != nil {
	panic(err)
}
fmt.Println(acct.Status)

fmt.Println(alpacaClient.GetLatestQuote("AAPL"))
bars := alpacaClient.GetBars(
	"AAPL", v2.Day, v2.Raw, time.Now().Add(-7*48*time.Hour), time.Now().Add(-20*time.Minute), 7)
var barset []v2.Bar

for bar := range bars {
	if bar.Error != nil {
		panic(bar.Error)
	}
	barset = append(barset, bar.Bar)
}
fmt.Println(barset[1])

stream.DataStreamURL = "https://stream.data.alpaca.markets"
stream.UseFeed("iex")

if err := stream.SubscribeTradeUpdates(tradeUpdateHandler); err != nil {
	fmt.Println(err)
	panic(err)
}

if err := stream.SubscribeTrades(tradeHandler, "AAPL"); err != nil {
	panic(err)
}
if err := stream.SubscribeQuotes(quoteHandler, "MSFT"); err != nil {
	panic(err)
}
if err := stream.SubscribeBars(barHandler, "IBM"); err != nil {
	panic(err)
}

select {}

// Cancel any open orders so they don't interfere with this script
alpacaClient.CancelAllOrders()

quantity := decimal.NewFromFloat(float64(100))

//fee, _ := decimal.NewFromString(".035")
//taxRate, _ := decimal.NewFromString(".08875")
adjSide := alpaca.Side("buy")
fmt.Println(*acct)
var sym = "AAPL"
_, err = alpacaClient.PlaceOrder(alpaca.PlaceOrderRequest{
	AccountID:   acct.ID,
	AssetKey:    &sym,
	Qty:         quantity,
	Side:        adjSide,
	TimeInForce: "gtc",
	Type:        "market", // [L] Change to alpaca.Limit
	// [L] Uncomment line below
	//LimitPrice:    &limitPrice,
	//TimeInForce:   alpaca.Day,
	//ClientOrderID: alp.currOrder,
})

if err == nil {
	fmt.Printf("Market order of | %d %s %s | completed.\n", quantity, sym, adjSide)
} else {
	fmt.Printf("Order of | %d %s %s | did not go through.\n", quantity, sym, adjSide)
}
*/
