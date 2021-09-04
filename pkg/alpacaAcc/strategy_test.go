package alpacaAcc

import (
	"fmt"
	"testing"
	"time"
)

func TestGoldenCross(t *testing.T) {

	symbol := "AAPL"
	client := Init()

	daysback := 10
	shortAv := 1
	longAv := 10
	startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback)*24*60*60), 0), time.Now()
	bars := GetHistData(client, symbol, &startTime, &endTime, daysback+longAv)
	daysback = len(bars)

	for i := 1; i < daysback; i++ {
		fmt.Println(shortAv)
		fmt.Println(longAv)
		daysback -= i
		fmt.Println(GoldenCross(bars, daysback, shortAv, longAv))
		if GoldenCross(bars, daysback, shortAv, longAv) {
			t.Errorf("The Av %t", GoldenCross(bars, daysback, shortAv, longAv))
		}
	}
	t.Errorf("The Av %t", GoldenCross(bars, daysback, shortAv, longAv))
}
