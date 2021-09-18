package alpacaAcc

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()
	_, err := client.GetAccount()
	if err != nil {
		t.Errorf("Expected object not %s", os.Getenv("GOROOT_1_17_X64"))
		t.Error(os.Getenv("API_Key_ID")[2])
	}
	if 0 < 2 {
		t.Error(os.Environ())
	}

}
