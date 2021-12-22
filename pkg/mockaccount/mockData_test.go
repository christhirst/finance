package mockaccount

import (
	"testing"
)

var tests = []struct {
	index     float64
	randLevel int
}{
	{1, 1},
	{20, 100},
}

func TestMockBars(t *testing.T) {
	//var want alpaca.Bar
	strength := 0.5
	for _, v := range tests {
		got := MockBar(v.index, v.randLevel, strength)

		if got.Close > float32(20)*float32(v.randLevel) || got.Close <= float32(10) {
			t.Errorf("got %f want %f", got.Close, v.index)

		}
	}

}

func TestCreateMockBars(t *testing.T) {
	got := CreateMockBars(200, 2, 10)
	if got == nil {
		t.Errorf("got %s want %s", "nil", "alpaca.bar")
	}

}
