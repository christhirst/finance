package runner

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func doSomething(s string) {
	fmt.Println("doing something", s)
}

func Runner(f func(*alpaca.Client, []string, []string), Client *alpaca.Client) {

	for {
		time.Sleep(2 * time.Second)
		go doSomething("from polling 1")
	}
}
