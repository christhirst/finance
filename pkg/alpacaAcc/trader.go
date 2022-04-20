package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	"github.com/shopspring/decimal"
)

func Trader(Client alpaca.Client, stock string, strat string, longAv int, shortAv int, ErrorChan chan<- error) {
	account, err := Client.GetAccount()

	if err != nil {
		ErrorChan <- err
	}
	// TODO: Buy at signal
	daysback := 200

	//+ one for minus one day
	startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
	daysback, err = Tradingdays(Client, daysback)
	ErrorChan <- err
	shortAv, err = Tradingdays(shortAv)
	ErrorChan <- err
	barsd, err := GetHistData(stock, &startTime, &endTime, daysback+longAv)
	bars := barsd[stock]
	if err != nil {
		ErrorChan <- err
	}

	position, _ := Client.GetAsset(stock)
	fmt.Println(position)

	if strat == "GoldenCross" {
		var adjSide alpaca.Side
		quantity := decimal.NewFromFloat(float64(100))
		longAv = len(bars) - daysback - 1
		if GoldenCross(bars[longAv-1:], shortAv) == 1 {
			adjSide = alpaca.Side("buy")
			order(Client, adjSide, quantity, &stock, account, -1)
		} else if GoldenCross(bars[longAv-1:], shortAv) == -1 {
			adjSide = alpaca.Side("sell")
			order(Client, adjSide, quantity, &stock, account, -1)
		}
	}

	/* 	if strat[1] == "engulfBullCandle" {
		if patternreconition.BullishEngulfingCandle(bars, 1) {
			fmt.Println("ddd")

		}
	} */
}
