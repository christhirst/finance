package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

func Trader(Client *alpaca.Client, stock string, strat []string) {
	account, err := Client.GetAccount()

	if err != nil {
		panic(err)
	}

	// TODO: Buy at signal
	
		daysback := 200
		longAv := 100
		shortAv := 50
		//+ one for minus one day
		startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
		daysback, shortAv = Tradingdays(Client, daysback), Tradingdays(Client, shortAv)
		bars := GetHistData(Client, stock, &startTime, &endTime, daysback+longAv)

		position, _ := Client.GetAsset(stock)
		fmt.Println(position)

		if strat[0] == "GoldenCross" {
			var adjSide alpaca.Side
			quantity := decimal.NewFromFloat(float64(100))
			longAv = len(bars) - daysback - 1
			if GoldenCross(bars, daysback, shortAv, longAv) == 1 {
				adjSide = alpaca.Side("buy")
				order(*Client, adjSide, quantity, &stock, account)
			} else if GoldenCross(bars, daysback, shortAv, longAv) == -1 {

				adjSide = alpaca.Side("sell")
				order(*Client, adjSide, quantity, &stock, account)

			}
		}

		/* 	if strat[1] == "engulfBullCandle" {
			if patternreconition.BullishEngulfingCandle(bars, 1) {
				fmt.Println("ddd")

			}
		} */

	}


