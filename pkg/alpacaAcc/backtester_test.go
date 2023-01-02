package alpacaAcc

import "testing"

func TestAlldaysofyear(t *testing.T) {
	alldaysofyear(2023)
	t.Error()
}

func TestGetdatebefore(t *testing.T) {
	clientCon := Initc()

	getdatebefore(clientCon.DataClient, "2023-01-02", 15)
	t.Error()
}

func TestAllsignals(t *testing.T) {
	stock := "AAPL"
	allsignals(stock)
	t.Error()
}
