package alpacaAcc

import (
	"math/rand"
	"testing"
	"time"
)

func TestGoldenCross(t *testing.T) {

	symbol := "AAPL"
	client := Init()

	for i := 0; i <= 20; i++ {

		min := 10
		daysback := rand.Intn(500) + min
		longAv := rand.Intn(daysback+min) + min

		startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
		bars := GetHistData(client, symbol, &startTime, &endTime, daysback+longAv)
		daysback = Tradingdays(client, daysback)

		longAv = len(bars) - daysback - 1
		shortAv := rand.Intn(longAv)
		for i := 0; i <= daysback; i++ {
			GoldenCross(bars, daysback-i, shortAv, longAv)
		}
	}
}
