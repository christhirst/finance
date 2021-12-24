package runner

import (
	"sync"
	"testing"

	"github.com/christhirst/finance/pkg/mockaccount"
)

func TestAnalyser(t *testing.T) {
	sum := 0.0
	stratList := []string{"GoldenCross"}
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	runs := 10
	var wg sync.WaitGroup

	chanCount := len(stratList) * len(stockList) * (runs + 1)
	position := make(chan confData, chanCount)

	for _, stock := range stockList {
		bars := mockaccount.CreateMockBars(500, 3, 20)

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
