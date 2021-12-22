package runner

import (
	"sync"
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

/* func TestRunner(t *testing.T) {
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	stratList := []string{"GoldenCross"}
	client := alpacaAcc.Init()
	Runner(client, stockList, stratList)
} */

func TestAnalyticRunner(t *testing.T) {
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	stratList := []string{"GoldenCross"}
	client := alpacaAcc.Init()
	runs := 10
	daysback := 500
	position := make(chan confData, runs)
	var wg sync.WaitGroup
	for _, stock := range stockList {
		startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()
		bars, _ := alpacaAcc.GetHistData(client, stock, &startTime, &endTime, daysback)
		for _, strat := range stratList {
			go func(strat string) {
				analyser(bars, stock, strat, position, runs, &wg)
			}(strat)
		}
	}

}
