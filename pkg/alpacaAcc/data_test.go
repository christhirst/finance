package alpacaAcc

import (
	"testing"
)

func TestGetLiveData(t *testing.T) {

}

func TestGetHistData(t *testing.T) {

	stocklist := []string{"AAPL"}
	numBars := 50
	for _, stock := range stocklist {
		bar := GetHistData(Init(), stock, numBars)
			if bar == nil {
			t.Error("Strock not found")
		}
		if len(bar) < 1 {
			t.Errorf("Slice < 1")
		}
	}
}
