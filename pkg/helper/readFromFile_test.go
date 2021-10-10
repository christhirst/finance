package helper

import "testing"

func TestCompanylist(t *testing.T) {

	_, err := companylist()
	if err != nil {
		t.Error(err)
	}
}
