package cmd

import "testing"

func TestInit(t *testing.T) {
	stockList := []string{"AAPL", "MSFT", "AMZN", "GOOGL", "JD"}
	stratList := []string{"GoldenCross"}
	Init(stockList, stratList)

}
