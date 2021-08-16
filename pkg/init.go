package pkg

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/shopspring/decimal"
)

type alpacaClientContainer struct {
	client     *alpaca.Client
	tickSize   int
	tickIndex  int
	baseBet    float64
	currOrder  string
	lastPrice  float64
	stock      string
	position   int64
	equity     float64
	marginMult float64
	seconds    int
}

func Init() {

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_Key_ID"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("Secret_Key"))
	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")

	alpacaClient := alpaca.NewClient(common.Credentials())
	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}

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




}
