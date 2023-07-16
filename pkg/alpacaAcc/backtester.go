package alpacaAcc

import (
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/rs/zerolog/log"
)

/*
	 func alldaysofyears(year int) (time.Time, time.Time) {
		// Set the start date to January 1st of the current year
		startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

		// Set the end date to December 31st of the current year
		endDate := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)

		// Iterate over each day in the range from startDate to endDate

		for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
			fmt.Println(d)
		}
		return startDate, endDate
	}
*/
func getdatebefore(Client AlpacaClientContainer, endday string, beforeDays int) (time.Time, int, error) {
	//only the format
	endDay, err := time.Parse("2006-01-02", endday)
	if err != nil {
		log.Error().Err(err).Msg("Parsing of day failed")
		return time.Time{}, 0, err
	}
	newD := float64(beforeDays) * 1.5
	newI := int(newD) + 5
	startDay := endDay.AddDate(0, 0, -newI-beforeDays)

	req := &alpaca.GetCalendarRequest{
		Start: startDay,
		End:   endDay,
	}
	tradingDays, err := Client.TradeClient.GetCalendar(*req)
	if err != nil {
		log.Error().Err(err).Msg("Getting calenderdates failed")
		return time.Time{}, 0, err
	}
	delta := len(tradingDays) - beforeDays
	ll := tradingDays[delta-1].Date
	dDay, err := time.Parse("2006-01-02", ll)
	if err != nil {
		log.Error().Err(err).Msg("Parsing of day failed")
		return time.Time{}, 0, err
	}
	diff := len(tradingDays[delta:])
	if diff != beforeDays {
		log.Fatal()
	}

	end, err := time.Parse("2006-01-02", endday)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	days := int(end.Sub(dDay).Hours() / 24)

	return dDay, days, nil

}

func allsignals(stock string, month int) {
	now := time.Now()
	startTime, endTime := now.AddDate(0, -month, 0), now.Add(-15*time.Minute)
	// Format the date as "year-month-day"

	ClientCont := Init()
	daysback := 500
	longAv := 130
	shortAv := 50

	startTime, _, err := getdatebefore(ClientCont, endTime.Format("2006-01-02"), daysback+longAv)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	barsd, err := GetHistDatas(*ClientCont.DataClient, stock, startTime, endTime, daysback+longAv)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	bars := barsd[stock]
	qty := 1

	for i := 0; i <= len(bars)-longAv-1; i++ {
		o := longAv + i
		oo := o + longAv
		if bars[o:oo][len(bars[o:oo])-1].Close == 0 {
			fmt.Println(bars[o:oo][len(bars[o:oo])-1].Close)
			break
		}
		checkedDay := bars[oo-1 : oo][0]
		//fmt.Println(checkedDay)
		signal := GoldenCross(bars[o-1:oo], shortAv)
		if signal == 1 {
			equityAmt := float64(qty) * checkedDay.Close
			ClientCont.long[stock].Order(stock, qty, equityAmt)
		}

	}

	fmt.Println(ClientCont.long[stock].equityAmt)
	ii := ClientCont.valuePositions(stock, bars[len(bars)-1].Close)
	fmt.Println(ii)
}

//1128 1258
