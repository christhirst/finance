package mockaccount

import (
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
)

func TestAddBuy(t *testing.T) {
	max := 50
	for i := 2; i <= max; i++ {
		//AddBuy(s, fq, fp)

	}

}

func TestMockPortfolio(t *testing.T) {
	MockPortfolio := new(MockPortfolio)
	MockPortfolio.Pos = make(map[string]alpaca.Position)
}
