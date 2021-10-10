package helper

import (
	"fmt"
	"log"
	"os"
)

func companylist() ([]byte, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	fmt.Println("####")
	dat, err := os.ReadFile("/workspace/finance/files/companies.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dat)

	return dat, err

}
