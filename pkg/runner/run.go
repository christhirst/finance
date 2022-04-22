package runner

import (
	"fmt"
	"sync"
	"time"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func doSomething(s string) {
	fmt.Println("doing something", s)
}

func Runner(Client alpacaAcc.AlpacaClientContainer, stockList []string, strats []string) {
	errorChan := make(chan error)
	for {
		for _, stock := range stockList {
			for _, strat := range strats {
				go alpacaAcc.Trader(Client, stock, strat, 100, 50, errorChan)
			}
		}
	}
}

func AnalyticRunner(Client alpacaAcc.AlpacaClientContainer, stockList []string, stratList []string) confData {
	daysback := 500
	runs := 50
	position := make(chan confData, runs)
	var wg sync.WaitGroup
	for {
		for _, stock := range stockList {
			startTime, endTime := time.Unix(time.Now().Unix()-int64((daysback+1)*24*60*60), 0), time.Now()
			bars, _ := alpacaAcc.GetHistData(Client.DataClient, stock, &startTime, &endTime, daysback)
			for _, strat := range stratList {
				go func(strat string) {
					analyser(bars, stock, strat, position, runs, &wg)
				}(strat)
			}
		}
	}
}
