package runner

import (
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

func analyser(bars []alpaca.Bar, stock string, strat string, position chan confData, runs int) {
	//+ one for minus one day
	sum := 0.0
	min := 10
	MockPortfolio := new(mockaccount.MockPortfolio)
	MockPortfolio.Pos = make(map[string]alpaca.Position)

	//MockPortfolio.Pos = make(map[string]alpaca.Position)

	for i := 0; i <= runs; i++ {
		MockPortfolio.Pos[stock] = alpaca.Position{
			Qty: helper.FloatToDecimal(0),
		}
		randlongAv := helper.RandomInRange(min+1, 100)
		randshortAv := i + 10
		go func(b []alpaca.Bar, rs int, rl int, ch <-chan confData) {
			for i := 0; i <= len(bars)-rl; i++ {
				if strat == "GoldenCross" {
					//var adjSide alpaca.Side
					//sicherheit mehr shares
					if alpacaAcc.GoldenCross(b[i:rl+i], rs) == 1 {
						//adjSide = alpaca.Side("buy")
						//fake buy
						MockPortfolio.AddBuy(stock, 1, b[i : rl+i][len(b[i:rl+i])-1].Close)
						MockPortfolio.Cash = MockPortfolio.Cash + b[i : rl+i][len(b[i:rl+i])-1].Close
						//order(*Client, adjSide, quantity, &stock, account)
					} else if alpacaAcc.GoldenCross(b[i:rl+i], rs) == -1 {
						//adjSide = alpaca.Side("sell")
						//fake sell
						MockPortfolio.AddBuy(stock, -1, b[i : rl+i][len(b[i:rl+i])-1].Close)
						MockPortfolio.Cash = MockPortfolio.Cash - b[i : rl+i][len(b[i:rl+i])-1].Close
						//order(*Client, adjSide, quantity, &stock, account)
					}
				}
				/* 	if strat[1] == "engulfBullCandle" {
					if patternreconition.BullishEngulfingCandle(bars, 1) {
						fmt.Println("ddd")

					}
				} */
			}

			position <- confData{
				stock,
				randlongAv,
				randshortAv,
				float64(MockPortfolio.Cash),
			}

		}(bars, randshortAv, randlongAv, position)
		pp := <-position
		if sum < pp.gain {
			sum = pp.gain
		}

	}

}
