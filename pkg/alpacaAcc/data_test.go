package alpacaAcc

import (
	"testing"
	"time"
)

func TestGetLiveData(t *testing.T) {

}

func TestGetHistData(t *testing.T) {

	stocklist := []string{"AAPL"}
	numBars := 50
	startTime, endTime := time.Unix(time.Now().Unix()-int64(50*24*60*60), 0), time.Now()
	for _, stock := range stocklist {
		now := startTime
		then := endTime
		bar := GetHistData(Init(), stock, &now, &then, numBars)
		if bar == nil {
			t.Error("Strock not found")
		}
		if len(bar) < 1 {
			t.Errorf("Slice < 1")
		}
	}
}