package alpacaAcc

import (
	"fmt"
	"testing"
	"time"
)

func TestGoldenCross(t *testing.T) {

	symbol := "AAPL"
	client := Init()

	daysback := 500
	shortAv := 10
	longAv := 300

	startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
	backTime := time.Unix(time.Now().Unix()-int64((daysback)*24*60*60), 0)
	bars := GetHistData(client, symbol, &startTime, &endTime, daysback+longAv)
	backBars := GetHistData(client, symbol, &backTime, &endTime, daysback)
	daysback = len(backBars)

	longAv = len(bars) - 1 - daysback
	for i := 0; i <= daysback; i++ {
		fmt.Println(len(bars) - longAv - daysback)
		if GoldenCross(bars, daysback-i, shortAv, longAv) && bars[daysback].Close > bars[daysback-1].Close {
			t.Errorf("Next Price is lower %t", GoldenCross(bars, daysback-i, shortAv, longAv))
		}
	}
}
