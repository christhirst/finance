package mockaccount

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/shopspring/decimal"
)

type MockPosition struct {
	Pos alpaca.Position
}

func (m MockPosition) AddQty(f float64) {
	deci := decimal.NewFromFloat(f)
	m.Pos.Qty = m.Pos.Qty.Add(deci)
}
