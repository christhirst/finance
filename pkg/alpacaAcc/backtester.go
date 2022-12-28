package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func allsignals() {
	stock := "AAPL"
	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(60*60*2), 0)
	numBars := 10

	start := "2022-12-20"
	end := "2022-12-27"
	clientCon := Initc()

	ee, err := clientCon.TradeClient.GetCalendar(&start, &end)
	fmt.Println(err)
	fmt.Println(ee)
	barss, err := GetHistData(clientCon.DataClient, stock, &startTime, &endTime, numBars)
	fmt.Println(barss)

	ClientCont := Initc()
	daysback := 200
	longAv := 150
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
