package helper

import (
	"fmt"
	"os"
)

func companylist() []byte {

	dat, err := os.ReadFile("files/companies.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dat)

	return dat

}
