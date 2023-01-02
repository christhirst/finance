package alpacaAcc

import "testing"

func TestAlldaysofyear(t *testing.T) {
	alldaysofyear(2023)
	t.Error()
}

func TestAllsignals(t *testing.T) {
	stock := "AAPL"
	allsignals(stock)
	t.Error()
}
