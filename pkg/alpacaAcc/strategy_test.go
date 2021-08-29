package alpacaAcc

import (
	"testing"
)

func TestGoldenCross(t *testing.T) {

	symbol := "AAPL"
	client := Init()
	shortAv := 50
	longAv := 100
	t.Errorf("The Av %t", GoldenCross(client, symbol, shortAv, longAv))

}
