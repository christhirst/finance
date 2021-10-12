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

func Runner(Client *alpaca.Client, stockList []string, strat []string) {
	for {
		time.Sleep(60 * time.Second)
		for _, stock := range stockList {
			go alpacaAcc.Trader(Client, stock, strat)

		}
	}
}
