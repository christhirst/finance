package runner

import (
	"sync"
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func TestAnalyser(t *testing.T) {
	stratList := []string{"GoldenCross"}
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	daysback := 500
	runs := 10
	var wg sync.WaitGroup
	client := alpacaAcc.Init()
	position := make(chan confData)
	startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()

	for _, stock := range stockList {
		bars := alpacaAcc.GetHistData(client, stock, &startTime, &endTime, daysback)
		for _, strat := range stratList {
			go analyser(bars, stock, strat, position, runs, wg)
		}
	}
	wg.Wait()
	//time.Sleep(time.Second * 50)
	ss := <-position

	t.Errorf("%f", ss.gain)
	t.Errorf("%d", ss.longAv)
	t.Errorf("%d", ss.shortAv)

}
