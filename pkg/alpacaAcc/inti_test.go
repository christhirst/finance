package alpacaAcc

import (
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func TestInit(t *testing.T) {
	client := Init()
	_, err := client.GetAccount()
	if err != nil {
		t.Error(err.Error())
	} else if common.Credentials().ID == "" {
		t.Errorf("Running w/ credentials [%v]\n", common.Credentials().ID)
	} else if common.Credentials().Secret == "" {
		t.Errorf("Running w/ credentials [%v]\n", common.Credentials().Secret)

	}
}
