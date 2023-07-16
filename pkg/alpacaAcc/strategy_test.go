package alpacaAcc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/christhirst/finance/pkg/helper"
)

func TestGoldenCross(t *testing.T) {
	symbol := "AAPL"
	clientCon := Init()
	now := time.Now()
	for i := 0; i <= 3; i++ {
		endTime := now.Add(-15 * time.Minute)
		min := 10
		daysback := rand.Intn(500) + min + 1
		longAv := helper.RandomInRange(min, daysback)
		//startTime := time.Unix(time.Now().Unix()-int64((longAv+daysback+1)*24*60*60), 0)
		//endTime := time.Now()
		startDay, deltaD, err := getdatebefore(clientCon, endTime.Format("2006-01-02"), longAv+daysback)

		bars, err := GetHistData(*clientCon.DataClient, symbol, startDay, endTime, deltaD)
		if err != nil {
			t.Error(err.Error())
		}
		/* daysback, err = Tradingdays(*client.DataClient, daysback, 15)
		if err != nil {
			t.Error(err.Error())
		} */

		longAv = len(bars[symbol]) - daysback - 1
		shortAv := rand.Intn(longAv - 1)
		for i := 0; i <= daysback; i++ {
			GoldenCross(bars[symbol][i:longAv+i], shortAv)
		}
	}
}
