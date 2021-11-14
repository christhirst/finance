package alpacaAcc

import (
	"os"
	"testing"
	"time"
)

func TestGetLiveData(t *testing.T) {

}

func TestGetHistData(t *testing.T) {

	stocklist := []string{"AAPL"}
	numBars := 10
	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Now()
	for _, stock := range stocklist {
		now := startTime
		then := endTime
		bar := GetHistData(Init(), stock, &now, &then, numBars)

		if Init() == nil {
			t.Errorf("Getting Account faild: %s", os.Getenv("API_Key_ID"))
		}

		if bar == nil {
			t.Error("No Data could be fetched", stock, now, then)
		}
		if len(bar) < 1 {
			t.Errorf("Slice < 1")
		}
	}
}
