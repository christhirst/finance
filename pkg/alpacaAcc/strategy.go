package alpacaAcc

import (
	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/gitpod/mycli/pkg/indicator"
)

func GoldenCross(bars []alpaca.Bar, daysback int, shortAvD int, longAvD int) int {
	shortBars := bars[(len(bars) - shortAvD - daysback) : len(bars)-daysback]
	longBars := bars[(len(bars) - longAvD - daysback) : len(bars)-daysback]

	shortBarsDaybefore := bars[(len(bars) - shortAvD - daysback - 1) : len(bars)-daysback-1]
	longBarsDaybefore := bars[(len(bars) - longAvD - daysback - 1) : len(bars)-daysback-1]

	shortAv := indicator.Avarage(shortBars)
	longAv := indicator.Avarage(longBars)
	shortAvDaybefore := indicator.Avarage(shortBarsDaybefore)
	longAvDaybefore := indicator.Avarage(longBarsDaybefore)

	if longAv <= shortAv && longAvDaybefore >= shortAvDaybefore {
		return 1
	}
	if longAv >= shortAv && longAvDaybefore <= shortAvDaybefore {
		return -1
	}
	return 0
}

// TODO: complete
