package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func Trader(c *alpaca.Client, stockList []string, strat []string) {

	for _, stock := range stockList {
		if strat[0] == "GoldenCross" {
			GoldenCross(c, stock, 50, 200)
		}
	}

}
