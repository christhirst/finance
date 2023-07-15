package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
)

func order(client alpaca.Client, adjSide alpaca.Side, quantity decimal.Decimal, sym string, account *alpaca.Account, mockPosition float64) (*alpaca.Position, error) {

	orderInformation := alpaca.PlaceOrderRequest{
		ClientOrderID: account.ID,
		Symbol:        sym,
		Qty:           &quantity,
		Side:          adjSide,
		TimeInForce:   "gtc",
		Type:          "market", // [L] Change to alpaca.Limit
		// [L] Uncomment line below
		//LimitPrice:    &limitPrice,
		//TimeInForce:   alpaca.Day,
		//ClientOrderID: alp.currOrder,
	}

	if mockPosition > 0 {
		_, err := client.PlaceOrder(orderInformation)
		if err != nil {
			panic(err)
		}
		return client.GetPosition(sym)
	} else {
		return client.GetPosition(sym)
	}

}
