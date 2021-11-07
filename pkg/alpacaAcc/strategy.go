package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/christhirst/finance/pkg/indicator"
)

func GoldenCross(lbars []alpaca.Bar, sbarsCount int) int {
	fmt.Println("####")
	fmt.Println(len(lbars))
	fmt.Println(sbarsCount)
	longBarsOnDay := lbars[(1):]
	longBarsBeforeDaybefore := lbars[:len(lbars)-1]

	shortBarsOnDay := lbars[(sbarsCount - 1):]
	shortBarsBeforeDaybefore := lbars[sbarsCount : len(lbars)-1]

	shortAv := indicator.Avarage(shortBarsOnDay)
	longAv := indicator.Avarage(longBarsOnDay)
	shortAvDaybefore := indicator.Avarage(shortBarsBeforeDaybefore)
	longAvDaybefore := indicator.Avarage(longBarsBeforeDaybefore)

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
