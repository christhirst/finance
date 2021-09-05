package alpacaAcc

import (
	"fmt"
	"testing"
	"time"
)

func TestGoldenCross(t *testing.T) {

	symbol := "AAPL"
	client := Init()

	daysback := 0
	shortAv := 40
	longAv := 100

	startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
	backTime := time.Unix(time.Now().Unix()-int64((daysback)*24*60*60), 0)
	bars := GetHistData(client, symbol, &startTime, &endTime, daysback+longAv)
	backBars := GetHistData(client, symbol, &backTime, &endTime, daysback)
	daysback = len(backBars)

	longAv = len(bars) - 1 - daysback
	for i := 0; i <= daysback; i++ {
		if (GoldenCross(bars, daysback-i, shortAv, longAv) == 1) && (bars[len(bars)-daysback+i-1].Close < bars[len(bars)-daysback+i-2].Close) {
			fmt.Println("######################")
			t.Errorf("Next Price is lower %d", GoldenCross(bars, daysback-i, shortAv, longAv))
		}
	}
}
