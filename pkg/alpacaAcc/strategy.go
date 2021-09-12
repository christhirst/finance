package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/indicator"
)

func GoldenCross(bars []alpaca.Bar, daysback int, shortAvD int, longAvD int) int {
	fmt.Println(len(bars))
	fmt.Println(-longAvD - daysback - 1)
	shortBars := bars[(len(bars) - shortAvD - daysback) : len(bars)-daysback]
	longBars := bars[(len(bars) - longAvD - daysback) : len(bars)-daysback]

	shortBarsDaybefore := bars[(len(bars) - shortAvD - daysback - 1) : len(bars)-daysback-1]
	longBarsDaybefore := bars[(len(bars) - longAvD - daysback - 1) : len(bars)-daysback-1]

	shortAv := indicator.Avarage(shortBars)
	longAv := indicator.Avarage(longBars)
	shortAvDaybefore := indicator.Avarage(shortBarsDaybefore)
	longAvDaybefore := indicator.Avarage(longBarsDaybefore)

	if longAv <= shortAv && longAvDaybefore >= shortAvDaybefore {
		fmt.Println(longAv, shortAv, longAvDaybefore, shortAvDaybefore)
		return 1
	}
	if longAv >= shortAv && longAvDaybefore <= shortAvDaybefore {
		return -1
	}
	return 0
}

// TODO: complete
