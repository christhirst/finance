package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gitpod/mycli/pkg/indicator"
)

func GoldenCross(bars []alpaca.Bar, daysback int, shortAvD int, longAvD int) bool {
	fmt.Println(len(bars))
	shortBars := bars[(len(bars) - shortAvD):]
	longBars := bars[(len(bars) - longAvD):]

	shortAv := indicator.Avarage(shortBars)
	longAv := indicator.Avarage(longBars)
	if longAv <= shortAv && bars[len(shortBars)-1].Close > bars[len(shortBars)-1].Close {
		fmt.Println(longAv)
		return true
	}
	return false

}

func DeathCross(s string, a float32) bool {

	return true
}
