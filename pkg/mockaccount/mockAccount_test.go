package mockaccount

import (
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

func TestAddBuy(t *testing.T) {
	MockPortfolio := MockPortfolio{Cash: 1000}

	//p := new(alpaca.Position)
	s := "AAPL"
	MockPortfolio.Pos = map[string]alpaca.Position{s: {}}
	max := 10

	fq := float64(44)
	fp := float64(55)
	for i := 0; i <= max; i++ {
		MockPortfolio.AddBuy(s, fq, fp)
	}
}

func TestMockPortfolio(t *testing.T) {
	MockPortfolio := MockPortfolio{Cash: 1000}
	MockPortfolio.Pos = make(map[string]alpaca.Position)

}
