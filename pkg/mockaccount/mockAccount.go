package mockaccount

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

type MockPortfolio struct {
	Pos map[string]alpaca.Position
	//Currency string          `json:"currency"`
	Cash float32 `json:"cash"`
}

func (m MockPortfolio) AddBuy(s string, fq float64, fp float32) {
	newQty := decimal.NewFromFloat(fq)
	newPrice := decimal.NewFromFloat(fq)
	if entry, ok := m.Pos[s]; ok {
		// Then we modify the copy
		entry.EntryPrice = entry.EntryPrice.Mul(entry.Qty).Add(newQty.Mul(newPrice))
		entry.Qty = entry.Qty.Add(newQty)
		// Then we reassign map entry
		m.Pos[s] = entry
	}

}
