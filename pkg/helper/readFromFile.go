package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func companylist() ([]byte, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	fmt.Println("####")
	dat, err := os.ReadFile(filepath.Join(path, "../../") + "/files/companies.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dat)

	return dat, err

}
