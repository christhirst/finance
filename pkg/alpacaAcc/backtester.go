package alpacaAcc

import (
	"fmt"
	"time"

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

func getdatebefore(day string, beforeDays int) (t time.Time) {
	t, err := time.Parse("2006-01-02", day)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Subtract 10 days from the date
	newD := float64(beforeDays) * 1.1
	newI := int(newD)
	t = t.AddDate(0, 0, -newI)
	fmt.Println(t)
	return

}

func allsignals(stock string) {

	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(60*60*2), 0)
	numBars := 10
	longbar := 200
	start := "2022-01-01"
	t := time.Now()

	// Format the date as "year-month-day"

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
