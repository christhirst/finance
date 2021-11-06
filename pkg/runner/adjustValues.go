package runner

import (
	"math/rand"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/alpacaAcc"
	"github.com/christhirst/finance/pkg/helper"
	"github.com/christhirst/finance/pkg/mockaccount"
)

func adjustToStocks(s string) {

	//Runner(Client, stockList, strat)
}

type confData struct {
	symbol  string
	longAv  int
	shortAv int
	gain    float64
}

func analyser(Client *alpaca.Client, stock string, strat string, position chan confData) {

	//+ one for minus one day
	sum := 0.0
	min := 10
	daysback := rand.Intn(500) + min
	longAv := rand.Intn(daysback+min) + min

	for {
		mockPosition := mockaccount.MockPosition{
			Pos: alpaca.Position{Qty: helper.FloatToDecimal(0)},
		}

		startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
		bars := alpacaAcc.GetHistData(Client, stock, &startTime, &endTime, daysback+longAv)
		longAv = len(bars) - daysback - 1
		shortAv := rand.Intn(longAv)

		for i := min; i <= daysback; i++ {
			go func([]alpaca.Bar, int, int, <-chan confData) {
				position <- confData{
					"ee",
					2,
					2,
					2,
				}
				if strat == "GoldenCross" {
					//var adjSide alpaca.Side
					//sicherheit mehr shares
					//quantity := decimal.NewFromFloat(float64(100))
					longAv = len(bars) - daysback - 1
					if alpacaAcc.GoldenCross(bars, daysback, shortAv, longAv) == 1 {
						//adjSide = alpaca.Side("buy")
						//fake buy

						mockPosition.AddQty(1)
						//order(*Client, adjSide, quantity, &stock, account)
					} else if alpacaAcc.GoldenCross(bars, daysback, shortAv, longAv) == -1 {
						//adjSide = alpaca.Side("sell")
						//fake sell
						mockPosition.AddQty(-1)
						//order(*Client, adjSide, quantity, &stock, account)

					}
				}
				/* 	if strat[1] == "engulfBullCandle" {
					if patternreconition.BullishEngulfingCandle(bars, 1) {
						fmt.Println("ddd")

					}
				} */

			}(bars, shortAv, longAv, position)

		}
		pp := <-position
		if sum < pp.gain {
			sum = pp.gain
		}

	}

}
