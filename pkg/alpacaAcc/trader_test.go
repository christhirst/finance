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
	errorChan := make(chan error, len(stockList)*2)

	longAv := 30
	shortAv := 10
	for _, stock := range stockList {
		Trader(client, stock, stratList[0], longAv, shortAv, errorChan)
	}
	close(errorChan)
	for err := range errorChan {
		if err != nil {
			t.Error(err.Error())
		}
	}

}
