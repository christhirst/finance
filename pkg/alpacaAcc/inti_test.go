package alpacaAcc

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()
	_, err := client.GetAccount()
	if err != nil {
		t.Error(err.Error())
		//t.Error(os.Environ())
		t.Error(os.Getenv("API_KEY_ID"))
		t.Error(os.Getenv("SECRET_KEY"))
	}

}
