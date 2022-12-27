package alpacaAcc

import (
	"time"

	"github.com/rs/zerolog/log"
)

func allsignals() {
	ClientCont := Initc()
	daysback := 200
	longAv := 150
	shortAv := 50
	stock := "AAPL"

	startTime, endTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0), time.Now()
	daysback, err := Tradingdays(ClientCont.DataClient, daysback)
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
