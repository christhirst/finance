package alpacaAcc

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/christhirst/finance/pkg/indicator"
	"github.com/rs/zerolog/log"
)

func GoldenCross(lbars []marketdata.Bar, shortAv int) int {

	if len(lbars) < shortAv {
		fmt.Println("##Panic##")
		fmt.Println(len(lbars))
		fmt.Println(shortAv)
		log.Panic()
	}
	longBarsOnDay := lbars[1:]
	longBarsBeforeDaybefore := lbars[:len(lbars)-1]

	shortBarsOnDay := lbars[len(lbars)-shortAv:]
	shortBarsBeforeDaybefore := lbars[len(lbars)-shortAv-1 : len(lbars)-1]

	shortAvf := indicator.Avarage(shortBarsOnDay)
	shortAvDaybefore := indicator.Avarage(shortBarsBeforeDaybefore)

	longAv := indicator.Avarage(longBarsOnDay)
	longAvDaybefore := indicator.Avarage(longBarsBeforeDaybefore)

	if longAv <= shortAvf && longAvDaybefore >= shortAvDaybefore {
		return 1
	}
	if longAv >= shortAvf && longAvDaybefore <= shortAvDaybefore {
		return -1
	}
	return 0
}
