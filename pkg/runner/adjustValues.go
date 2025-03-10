package runner

import (
	"sync"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
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

func analyser(bars []marketdata.Bar, stock string, strat string, position chan confData, runs int, wg *sync.WaitGroup) {
	defer wg.Done()
	//+ one for minus one day
	min := 10
	MockPortfolio := new(mockaccount.MockPortfolio)
	MockPortfolio.Pos = make(map[string]alpaca.Position)

	for i := 0; i <= runs; i++ {
		wg.Add(1)
		MockPortfolio.Mu.Lock()
		MockPortfolio.Pos[stock] = alpaca.Position{
			Qty: helper.FloatToDecimal(0),
		}
		MockPortfolio.Mu.Unlock()
		randlongAv := helper.RandomInRange(min+1, 100)
		randshortAv := i + 10

		go func(b []marketdata.Bar, rs int, rl int, ch <-chan confData, wg *sync.WaitGroup) {
			//todo |----| bars is long; it has to be used over the hole bar range
			for i := 0; i <= len(bars)-rl; i++ {
				if rs > rl-5 {
					break
				}
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
			go func() {
				defer wg.Done()
			}()
		}(bars, randshortAv, randlongAv, position, wg)

	}

}
