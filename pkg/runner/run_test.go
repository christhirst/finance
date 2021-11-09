package runner

import (
	"testing"

	"github.com/christhirst/finance/pkg/alpacaAcc"
)

func TestRunner(t *testing.T) {
	stockList := []string{"AAPL"} // "MSFT", "AMZN", "GOOGL", "JD"}
	stratList := []string{"GoldenCross"}
	client := alpacaAcc.Init()
	Runner(client, stockList, stratList)
}
