package alpacaAcc

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()

	if client == nil {
		t.Errorf("Expected object not %s", os.Getenv("API_Key_ID")[:4])
	}

}
