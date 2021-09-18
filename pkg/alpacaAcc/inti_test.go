package alpacaAcc

import (
	"os"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/common"
)

func TestInit(t *testing.T) {
	client := Init()
	if client == nil {
		t.Errorf("Expected object not %s", os.Setenv(common.EnvApiKeyID, os.Getenv("API_Key_ID")))
	}

}
