package runner

import (
	"fmt"
	"os"

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

func analyser(bars []alpaca.Bar, stock string, strat string, position chan confData) {

	//+ one for minus one day
	sum := 0.0
	min := 10
	for i := 0; i <= 1; i++ {
		mockPosition := mockaccount.MockPosition{
			Pos: alpaca.Position{Qty: helper.FloatToDecimal(0)},
		}
		position <- confData{
			"ee",
			2,
			2,
			2,
		}
		randlongAv := helper.RandomInRange(min+1, 200)
		randshortAv := helper.RandomInRange(min, randlongAv)
		go func(b []alpaca.Bar, rs int, rl int, ch <-chan confData) {
			for i := 0; i <= len(bars)-rl; i++ {
				fmt.Println(i)

				if strat == "GoldenCross" {
					//var adjSide alpaca.Side
					//sicherheit mehr shares
					//quantity := decimal.NewFromFloat(float64(100))
					fmt.Println("len(bars)")
					fmt.Fprintf(os.Stdout, "index %d", i)
					fmt.Fprintf(os.Stdout, "barlen %d", len(b))
					fmt.Fprintf(os.Stdout, "rl %d", rl)
					fmt.Fprintf(os.Stdout, "bi:rl+i %d", len(b[i:rl+i]))
					fmt.Fprintf(os.Stdout, "b[i:] %d !!", len(b[i:]))
					if alpacaAcc.GoldenCross(b[i:rl+i], rs) == 1 {
						//adjSide = alpaca.Side("buy")
						//fake buy

						mockPosition.AddQty(1)
						//order(*Client, adjSide, quantity, &stock, account)
					} else if alpacaAcc.GoldenCross(b[i:rl+i], rs) == -1 {
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

			}

		}(bars, randshortAv, randlongAv, position)
		pp := <-position
		if sum < pp.gain {
			sum = pp.gain
		}

	}

}
