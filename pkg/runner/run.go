package runner

import (
	"fmt"
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

func AnalyticRunner(Client *alpaca.Client, stockList []string, strats []string) {

	position := make(chan confData)

	for {

		for _, stock := range stockList {
			for _, strat := range strats {
				go func(strat string) {
					analyser(Client, stock, strat, position)
					time.Sleep(60 * time.Second)
				}(strat)
			}
		}
	}
}
