package mockaccount

import (
	"sync"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

type MockPortfolio struct {
	Mu  sync.Mutex
	Pos map[string]alpaca.Position
	//Currency string          `json:"currency"`
	Cash float32 `json:"cash"`
}

func (m *MockPortfolio) AddBuy(s string, fq float64, fp float32) {
	newQty := decimal.NewFromFloat(fq)
	newPrice := decimal.NewFromFloat(fq)
	m.Mu.Lock()
	if entry, ok := m.Pos[s]; ok {
		// Then we modify the copy
		if !newQty.Add(entry.Qty).Equals(decimal.NewFromFloat(0)) {
			entry.EntryPrice = entry.EntryPrice.Mul(entry.Qty).Add(newQty.Mul(newPrice).Div(newQty.Add(entry.Qty)))
			entry.Qty = entry.Qty.Add(newQty)
		} else {
			entry.Qty = entry.Qty.Add(decimal.NewFromFloat(0))
		}
		// Then we reassign map entry

		m.Pos[s] = entry
		m.Mu.Unlock()

	}

}
