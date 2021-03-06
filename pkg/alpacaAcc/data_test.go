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
	stocklist := []string{"AAPL"}
	numBars := 10
	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Unix(time.Now().Unix()-int64(60*60*2), 0)
	clientCon := Initc()
	for _, stock := range stocklist {
		now := startTime
		then := endTime
		bar, err := GetHistData(clientCon.DataClient, stock, &now, &then, numBars)
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
