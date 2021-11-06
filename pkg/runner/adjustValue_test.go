package runner

import (
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func Testanalyser(t *testing.T) {
	client := alpacaAcc.Init()
	position := make(chan confData)
	for {
		time.Sleep(60 * time.Second)
		for _, stock := range stockList {
			for _, strat := range strats {
				go analyser(client, stock, strat, position)
			}
		}
	}

}
