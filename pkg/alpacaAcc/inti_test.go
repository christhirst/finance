package alpacaAcc

import (
	"testing"
)

func TestInit(t *testing.T) {
	client := Init()
	if client == nil {
		t.Errorf("Expected object not %s", "nil")
	}

}
