package store

import "testing"

func TestAddEneeetryToDbss(t *testing.T) {
	url := "https://raw.githubusercontent.com/christhirst/finance/main/files/companies.csv"
	comma := ','
	comment := '#'
	readCsv(url, comma, comment)
	t.Error()
}
