package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func allsignals(stock string) {

	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(60*60*2), 0)
	numBars := 10
	longbar := 200
	start := "2022-01-01"
	t := time.Now()

	// Format the date as "year-month-day"
	end := t.Format("2006-01-02")

	clientCon := Initc()

	ee, err := clientCon.TradeClient.GetCalendar(&start, &end)
	fmt.Println(err)
	fmt.Println(len(ee))
	barslength :=len(ee[len(ee)-longbar:])
	fmt.Println(barslength)

	barss, err := GetHistData(clientCon.DataClient, stock, &startTime, &endTime, numBars)
	fmt.Println(len(barss))

	ClientCont := Initc()
	daysback := 200
	longAv := 130
	shortAv := 50

	daysback, err = Tradingdays(ClientCont.DataClient, daysback)
	if err != nil {
		log.Error().Err(err).Int("daysback", daysback).Msg("")
	}
	shortAv, err = Tradingdays(ClientCont.DataClient, shortAv)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	barsd, err := GetHistData(ClientCont.DataClient, stock, &startTime, &endTime, daysback+longAv)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	bars := barsd[stock]
	GoldenCross(bars[longAv-1:], shortAv)

}
