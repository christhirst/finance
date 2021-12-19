package runner

import (
	"fmt"
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
	position := make(chan confData, 20)
	startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()

	for _, stock := range stockList {
		bars, _ := alpacaAcc.GetHistData(client, stock, &startTime, &endTime, daysback)

		for _, strat := range stratList {
			go analyser(bars, stock, strat, position, runs, &wg)

		}
		fmt.Printf("position: %v\n", wg)
	}
	wg.Wait()
	fmt.Printf("position: %v\n", len(position))
	close(position)
	fmt.Printf("position: %v\n", len(position))
	for p := range position {
		t.Errorf("%f", p.gain)
		if sum < p.gain {
			t.Errorf("%f", p.gain)
			sum = p.gain
		}
		if p.gain != 0 {
			t.Errorf("%f", p.gain)
			t.Errorf("%d", p.longAv)
			t.Errorf("%d", p.shortAv)
		}
	}
	t.Errorf("%v", "mmm")
}
