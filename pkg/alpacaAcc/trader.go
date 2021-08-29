package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gitpod/mycli/pkg/indicator"
)

func Trader(c *alpaca.Client, stockList []string, strat []string) {

	for _, stock := range stockList {
		bars := GetHistData(c, stock)
		if strat[0] == "GoldenCross" {
			GoldenCross(stock, indicator.Avarage(bars))
		}
	}

}
