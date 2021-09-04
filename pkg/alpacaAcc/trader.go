package alpacaAcc

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func Trader(c *alpaca.Client, stockList []string, strat []string) {

	for _, stock := range stockList {
		daysback := 200
		startTime, endTime := time.Unix(time.Now().Unix()-int64(daysback*24*60*60), 0), time.Now()
		bars := GetHistData(c, stock, &startTime, &endTime, daysback)
		if strat[0] == "GoldenCross" {
			GoldenCross(bars, 0, 50, daysback)
		}
	}

}
