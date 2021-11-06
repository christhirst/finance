package helper

import "github.com/shopspring/decimal"

func FloatToDecimal(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)

}

