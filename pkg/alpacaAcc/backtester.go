package alpacaAcc

/* func alldaysofyear(year int) {
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
	//tt := t.Format("2006-01-02")
	//startTime, endTime := time.Unix(time.Now().Unix()-int64(days*24*60*60), 0), time.Now()

	ts := time.Time{}
	ds := time.Time{}

	req := &alpaca.GetCalendarRequest{
		Start: ts,
		End:   ds,
	}
	days, err := Client.TradeClient.GetCalendar(*req)
	if err != nil {
		log.Error().Err(err).Msg("Getting calenderdates failed")
		return
	}
	diff = len(days)

	return t, diff

}

func allsignals(stock string, month int) {
	now := time.Now()
	startTime, endTime := now.AddDate(0, -month, 0), now.Add(-15*time.Minute)
	//.Unix()-int64(60*60*25), 0)
	fmt.Println(startTime)
	fmt.Println(endTime)
	// Format the date as "year-month-day"

	ClientCont := Init()
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
	barsd, err := GetHistDatas(ClientCont.DataClient, stock, startTime, endTime, daysback+longAv)
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	bars := barsd[stock]
	fmt.Println(len(bars))
	for i := 0; i <= len(bars)-longAv-1; i++ {
		o := longAv + i
		oo := o + longAv
		//fmt.Println(bars[o:oo][len(bars[o:oo])-1])
		if bars[o:oo][len(bars[o:oo])-1].Close == 0 {
			fmt.Println(bars[o:oo][len(bars[o:oo])-1].Close)
			break
		}
		ii := GoldenCross(bars[o-1:oo], shortAv)
		fmt.Println(ii)
	}
}

//1128 1258
*/
