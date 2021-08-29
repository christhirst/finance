package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gitpod/mycli/pkg/indicator"
)

func GoldenCross(c *alpaca.Client, symbol string, shortAv int, longAv int) bool {

	barsShort := GetHistData(c, symbol, shortAv)
	barsLong := GetHistData(c, symbol, longAv)
	if indicator.Avarage(barsLong) > indicator.Avarage(barsShort) {
		fmt.Println(indicator.Avarage(barsShort))
		return true
	}
	return false

}

func DeathCross(s string, a float32) bool {

	return true
}
