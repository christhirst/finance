package alpacaAcc

import (
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()
	_, err := client.GetAccount()
	if err != nil {
		t.Error(err.Error())
	}
}
