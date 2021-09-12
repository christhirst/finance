package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	patternreconition "github.com/christhirst/finance/pkg/patternReconition"
	"github.com/shopspring/decimal"
)

func Trader(Client *alpaca.Client, stockList []string, strat []string) {
	account, err := Client.GetAccount()

	if err != nil {
		panic(err)
	}

	// TODO: Buy at signal
	for _, stock := range stockList {
		daysback := 90
		startTime, endTime := time.Unix(time.Now().Unix()-int64(daysback*24*60*60), 0), time.Now()
		bars := GetHistData(Client, stock, &startTime, &endTime, 0)
		fmt.Println("##", len(bars))

		position, _ := Client.GetAsset(stock)
		fmt.Println(position)

		if strat[0] == "GoldenCross" {
			var adjSide alpaca.Side
			quantity := decimal.NewFromFloat(float64(100))
			if GoldenCross(bars, 0, 50, daysback) == 1 {
				adjSide = alpaca.Side("buy")
				order(*Client, adjSide, quantity, &stock, account)
			} else if GoldenCross(bars, 10, 50, 0) == -1 {

				adjSide = alpaca.Side("sell")
				order(*Client, adjSide, quantity, &stock, account)

			}
		}

		if strat[1] == "engulfBullCandle" {
			if patternreconition.BullishEngulfingCandle(bars, 1) {
				fmt.Println("ddd")

			}
		}

	}

}
