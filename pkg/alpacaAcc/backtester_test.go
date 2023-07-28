package alpacaAcc

import (
	"testing"
)

/* func TestAlldaysofyear(t *testing.T) {
	years, diff := alldaysofyear(2023)
	fmt.Println(years)
	fmt.Println(diff)
	t.Error()
} */

func TestGetdatebefore(t *testing.T) {
	clientCon := Init()

	t.Run("Init connection", func(t *testing.T) {
		got := "2023-03-07"
		days := 100
		want := 147
		daybefore, diff, err := getdatebefore(clientCon, got, days)
		if diff != want {
			t.Error(diff, daybefore, err)
		}
	})
	//t.Error()
}

func TestAllsignals(t *testing.T) {
	stock := "AAPL"
	allsignals(stock, 5*12)
	t.Error()
}
