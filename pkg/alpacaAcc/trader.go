package alpacaAcc

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

func Trader(Client *alpaca.Client, stockList []string, strat []string) {
	account, err := Client.GetAccount()
	if err != nil {
		panic(err)
	}

	// TODO: Buy at signal
	for _, stock := range stockList {
		daysback := 300
		startTime, endTime := time.Unix(time.Now().Unix()-int64(daysback*24*60*60), 0), time.Now()
		bars := GetHistData(Client, stock, &startTime, &endTime, 0)

		if strat[0] == "GoldenCross" {

			adjSide := alpaca.Side("buy")
			quantity := decimal.NewFromFloat(float64(100))
			if GoldenCross(bars, 10, 50, 0) == 1 {
				order(*Client, adjSide, quantity, &stock, account)
			}
		}
		if strat[0] == "DeathCross" {
			// get position
			adjSide := alpaca.Side("sell")
			quantity := decimal.NewFromFloat(float64(100))
			if DeathCross(bars, 10, 50, 0) {
				order(*Client, adjSide, quantity, &stock, account)
			}
		}
	}

}
