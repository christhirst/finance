package helper

import (
	"fmt"
	"testing"
)



func TestReadCsvFile(t *testing.T) {
	records := readCsvFile("../../files/companies.csv")
	fmt.Println(records)

}
