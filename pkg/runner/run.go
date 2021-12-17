package runner

import (
	"fmt"
	"sync"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func doSomething(s string) {
	fmt.Println("doing something", s)
}

func Runner(Client *alpaca.Client, stockList []string, strats []string) {
	for {
		time.Sleep(60 * time.Second)
		for _, stock := range stockList {
			for _, strat := range strats {
				go alpacaAcc.Trader(Client, stock, strat, 100, 50)
			}
		}
	}
}

func AnalyticRunner(Client *alpaca.Client, stockList []string, stratList []string) confData {
	daysback := 500
	runs := 50
	position := make(chan confData, runs)
	var wg sync.WaitGroup
	for {
		for _, stock := range stockList {
			startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()
			bars := alpacaAcc.GetHistData(Client, stock, &startTime, &endTime, daysback)
			for _, strat := range stratList {
				go func(strat string) {
					analyser(bars, stock, strat, position, runs, &wg)
					time.Sleep(60 * time.Second)
				}(strat)
			}
		}
	}
}
