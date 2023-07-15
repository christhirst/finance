package mockaccount

import (
	"fmt"
	"sync"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
)

type MockPortfolio struct {
	Mu  sync.Mutex
	Pos map[string]alpaca.Position
	//Currency string          `json:"currency"`
	Cash float64 `json:"cash"`
}

func (m *MockPortfolio) AddBuy(s string, fq float64, fp float64) {
	newQty := decimal.NewFromFloat(fq)
	newPrice := decimal.NewFromFloat(fp)
	m.Mu.Lock()
	fmt.Println("rr")
	if entry, ok := m.Pos[s]; ok {
		// Then we modify the copy
		if !newQty.Add(entry.Qty).Equals(decimal.NewFromFloat(0)) {
			entry.AvgEntryPrice = entry.AvgEntryPrice.Mul(entry.Qty).Add(newQty.Mul(newPrice).Div(newQty.Add(entry.Qty)))
			entry.Qty = entry.Qty.Add(newQty)
		} else {
			entry.Qty = entry.Qty.Add(decimal.NewFromFloat(0))
		}
		// Then we reassign map entry

		m.Pos[s] = entry
		m.Mu.Unlock()
		fmt.Println("dd")
	}

}
