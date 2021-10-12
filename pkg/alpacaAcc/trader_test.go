package alpacaAcc

import "testing"

func TestTrader(t *testing.T) {
	client := Init()
	stratList := []string{"GoldenCross"}
	stockList := []string{"AAPL", "MSFT", "AMZN", "GOOGL", "JD"}

	for _, stock := range stockList {
		Trader(client, stock, stratList)
	}

}
