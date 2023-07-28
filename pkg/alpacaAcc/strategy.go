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

func Derivation(lbars []marketdata.Bar, sStart, sEnd int) int {
	longLength := len(lbars)
	longEnd := lbars[longLength-1]
	longBegin := lbars[0]
	longPriceDelta := (longEnd.Close - longBegin.Close) * 24 * 60 * 60 * float64(longLength)
	longDeltat := longEnd.Timestamp.Sub(longBegin.Timestamp).Seconds() * longBegin.Close
	longPriceDeriv := longPriceDelta / longDeltat

	shortBegin := lbars[sStart]
	shortEnd := lbars[sEnd-1]
	shortPriceDelta := (shortEnd.Close - shortBegin.Close) * 24 * 60 * 60 * float64(sEnd-sStart)
	shortDeltat := shortEnd.Timestamp.Sub(shortBegin.Timestamp).Seconds() //* shortBegin.Close
	shortPriceDeriv := shortPriceDelta / shortDeltat
	fmt.Println(longPriceDeriv)
	fmt.Println(longPriceDeriv)
	fmt.Println(longBegin)
	fmt.Println(longEnd)
	fmt.Println(shortPriceDeriv)
	fmt.Println(shortBegin)
	fmt.Println(shortEnd)
	//shortDeltat := shortPeriod
	if true {
	}

	return 2
}
