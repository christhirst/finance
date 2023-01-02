package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v2/marketdata"
	"github.com/rs/zerolog/log"
)

func alldaysofyear(year int) {
	// Set the start date to January 1st of the current year
	fmt.Println(year)
	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Set the end date to December 31st of the current year
	endDate := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)

	// Iterate over each day in the range from startDate to endDate

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		fmt.Println(d)
	}
	return

}

func getdatebefore(Client marketdata.Client, day string, beforeDays int) (t time.Time) {
	t, err := time.Parse("2006-01-02", day)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Subtract 10 days from the date
	newD := float64(beforeDays) * 1.5
	newI := int(newD) + 5
	t = t.AddDate(0, 0, -newI)
	tt := t.Format("2006-01-02")
	fmt.Println(t)
	//startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now()

	clientCon := Initc()

	ee, err := clientCon.TradeClient.GetCalendar(&tt, &day)
	fmt.Println(err)
	fmt.Println(len(ee))
	barslength := len(ee[len(ee)-beforeDays:])
	fmt.Println(barslength)

	return t

}

func allsignals(stock string) {

	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(60*60*2), 0)

	// Format the date as "year-month-day"

	ClientCont := Initc()
	daysback := 200
	longAv := 130
	shortAv := 50

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
