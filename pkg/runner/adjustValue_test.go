package runner

import (
	"sync"
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func TestAnalyser(t *testing.T) {
	sum := 0.0
	stratList := []string{"GoldenCross"}
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	daysback := 500
	runs := 10
	var wg sync.WaitGroup
	client := alpacaAcc.Init()
	chanCount := len(stratList) * len(stockList) * (runs + 1)

	position := make(chan confData, chanCount)
	startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()

	for _, stock := range stockList {
		bars, _ := alpacaAcc.GetHistData(client, stock, &startTime, &endTime, daysback)

		for _, strat := range stratList {
			wg.Add(1)
			go analyser(bars, stock, strat, position, runs, &wg)

		}
	}
	wg.Wait()
	close(position)
	for p := range position {
		if sum == p.gain {
			t.Errorf("%f", p.gain)
			sum = p.gain
		}
		if p.gain == 1 {
			t.Errorf("%f", p.gain)
			t.Errorf("%d", p.longAv)
			t.Errorf("%d", p.shortAv)
		}
	}
}
