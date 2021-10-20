package runner

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func adjustToStocks(s string) {

	//Runner(Client, stockList, strat)
}

func analyser(Client *alpaca.Client, stock string, strat string, ch <-chan string) {
	/* 	account, err := Client.GetAccount()

	   	if err != nil {
	   		panic(err)
	   	} */

	daysback := 200

	//+ one for minus one day
	for {
		shortAv := 22
		longAv := 444

		startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
		daysback, shortAv = alpacaAcc.Tradingdays(Client, daysback), alpacaAcc.Tradingdays(Client, shortAv)
		bars := alpacaAcc.GetHistData(Client, stock, &startTime, &endTime, daysback+longAv)

		position, _ := Client.GetAsset(stock)
		fmt.Println(position)

		if strat == "GoldenCross" {
			//var adjSide alpaca.Side
			//sicherheit mehr shares
			//quantity := decimal.NewFromFloat(float64(100))
			longAv = len(bars) - daysback - 1
			if alpacaAcc.GoldenCross(bars, daysback, shortAv, longAv) == 1 {
				//adjSide = alpaca.Side("buy")
				//fake buy
				//order(*Client, adjSide, quantity, &stock, account)
			} else if alpacaAcc.GoldenCross(bars, daysback, shortAv, longAv) == -1 {
				//adjSide = alpaca.Side("sell")
				//fake sell
				//order(*Client, adjSide, quantity, &stock, account)

			}
		}

		/* 	if strat[1] == "engulfBullCandle" {
			if patternreconition.BullishEngulfingCandle(bars, 1) {
				fmt.Println("ddd")

			}
		} */

	}
}
