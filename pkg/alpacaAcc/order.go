package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

func order(client alpaca.Client, adjSide alpaca.Side, quantity decimal.Decimal, sym *string, account *alpaca.Account) {

	orderInformation := alpaca.PlaceOrderRequest{
		AccountID:   account.ID,
		AssetKey:    sym,
		Qty:         quantity,
		Side:        adjSide,
		TimeInForce: "gtc",
		Type:        "market", // [L] Change to alpaca.Limit
		// [L] Uncomment line below
		//LimitPrice:    &limitPrice,
		//TimeInForce:   alpaca.Day,
		//ClientOrderID: alp.currOrder,
	}

	_, err := client.PlaceOrder(orderInformation)
	if err != nil {
		panic(err)
	}
}
