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

func getdatebefore(Client AlpacaClientContainer, day string, beforeDays int) (t time.Time, diff int) {
	//only the format
	t, err := time.Parse("2006-01-02", day)
	if err != nil {
		log.Error().Err(err).Msg("Parsing of day failed")
		return
	}

	// Subtract 10 days from the date
	newD := float64(beforeDays) * 1.5
	newI := int(newD) + 5
	t = t.AddDate(0, 0, -newI)
	//only the format
	tt := t.Format("2006-01-02")
	//startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now()

	days, err := Client.TradeClient.GetCalendar(&tt, &day)
	if err != nil {
		log.Error().Err(err).Msg("Getting calenderdates failed")
		return
	}
	diff = len(days[len(days)-beforeDays:])

	return

}

func allsignals(stock string, month int) {
	now := time.Now()
	startTime, endTime := now.AddDate(0, -month, 0), now.Add(-15*time.Minute)
	//.Unix()-int64(60*60*25), 0)
	fmt.Println(startTime)
	fmt.Println(endTime)
	// Format the date as "year-month-day"

	ClientCont := Initc()
	daysback := 500
	longAv := 130
	shortAv := 50
	minBack := 15

	daysback, err := Tradingdays(ClientCont.DataClient, daysback, minBack)
	if err != nil {
		log.Error().Err(err).Int("daysback", daysback).Msg("")
	}
	shortAv, err = Tradingdays(ClientCont.DataClient, shortAv, 15)
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
