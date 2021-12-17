package mockaccount

import (
	"fmt"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/helper"
)

func TestAddBuy(t *testing.T) {
	max := 50
	for i := 2; i <= max; i++ {
		max := helper.RandomDeci(i).Mul(helper.RandomDeci(i))
		if max.Equals(helper.RandomDeci(i)) {
			fmt.Println(max)
			t.Errorf("%d", max)
		}
	}

}

func TestMockPortfolio(t *testing.T) {
	MockPortfolio := new(MockPortfolio)
	MockPortfolio.Pos = make(map[string]alpaca.Position)
}
