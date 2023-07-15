package alpacaAcc

import (
	"fmt"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func TestInit(t *testing.T) {
	client := Init()
	fmt.Println(client)
	if common.Credentials().ID == "" {
		t.Errorf("Running w/ credentials [%v]\n", common.Credentials().ID)
	} else if common.Credentials().Secret == "" {
		t.Errorf("Running w/ credentials [%v]\n", common.Credentials().Secret)

	}
}
