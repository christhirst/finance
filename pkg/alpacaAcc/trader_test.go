package alpacaAcc

import (
	"fmt"
	"testing"
)

func TestTrader(t *testing.T) {
	client := Init()
	fmt.Println(client)

	stratList := []string{"GoldenCross"}
	stockList := []string{"AAPL", "MSFT", "AMZN", "GOOGL", "JD"}
	longAv := 30
	shortAv := 10
	for _, stock := range stockList {
		err := Trader(client, stock, stratList[0], longAv, shortAv)
		if err != nil {
			t.Error(err.Error())
		}

	}

}
