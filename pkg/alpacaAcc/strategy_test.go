package alpacaAcc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/helper"
)

func TestGoldenCross(t *testing.T) {
	symbol := "AAPL"
	client := Init()
	for i := 0; i <= 3; i++ {
		min := 10
		daysback := rand.Intn(500) + min + 1
		longAv := helper.RandomInRange(min, daysback)
		startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
		bars := GetHistData(client, symbol, &startTime, &endTime, daysback+longAv)
		daysback = Tradingdays(client, daysback)

		longAv = len(bars) - daysback - 1
		shortAv := rand.Intn(longAv - 1)
		for i := 0; i <= daysback; i++ {
			GoldenCross(bars[i:longAv+i], shortAv)
		}
	}
}
