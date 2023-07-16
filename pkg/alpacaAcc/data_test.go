package alpacaAcc

import (
	"net"
	"os"
	"testing"
	"time"
)

func TestGetLiveData(t *testing.T) {
	stocklist := []string{"AAPL"}
	GetLiveData(stocklist[0])
	t.Error()

}

func TestGetHistData(t *testing.T) {
	/* start := time.Date(2021, 8, 9, 13, 30, 0, 0, time.UTC)
	end := time.Date(2022, 3, 9, 13, 30, 1, 0, time.UTC) */
	stocklist := []string{"AAPL"}
	numBars := 10
	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(30*60*1), 0)
	clientCon := Init()
	for _, stock := range stocklist {
		now := startTime
		then := endTime
		bar, err := GetHistData(*clientCon.DataClient, stock, now, then, numBars)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			t.Error(err.Error())
		}
		if clientCon.DataClient == nil {
			t.Errorf("Getting Account faild: %s", os.Getenv("API_Key_ID"))
		}
		if bar == nil {
			t.Error("No Data could be fetched", stock, now, then, err)
		}
		if len(bar) < 1 {
			t.Errorf("Slice < 1")
		}
	}
}
