package alpacaAcc

import (
	"fmt"
	"testing"
)

func TestAlldaysofyear(t *testing.T) {
	alldaysofyear(2023)
	t.Error()
}

func TestGetdatebefore(t *testing.T) {
	clientCon := Init()
	getdatebefore(clientCon, "2023-01-02", 15)

	t.Run("Init connection", func(t *testing.T) {
		got := "2023-01-02"
		want := 15
		daybefore, diff := getdatebefore(clientCon, got, want)
		if diff != want {
			fmt.Println(daybefore)
			t.Error()
		}
	})
}

func TestAllsignals(t *testing.T) {
	stock := "AAPL"
	allsignals(stock, 5*12)
	t.Error()
}
