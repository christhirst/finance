package alpacaAcc

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()
	_, err := client.GetAccount()
	if err != nil {
		t.Errorf("Expected object not %s", os.Getenv("API_Key_ID")[:4])
	}
	t.Error(client.GetAccount())

}
