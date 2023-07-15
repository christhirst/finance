package store

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func readCsv(url string, comma, comment rune) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("Error reading url: %v", err)
	}
	reader := csv.NewReader(resp.Body)
	reader.Comma = comma     //','    Set the delimiter to comma
	reader.Comment = comment //'#'  Ignore lines starting with #
	records, err := reader.ReadAll()
	if err != nil {
		log.Error().Err(err).Msgf("Error reading all records: %v", err)
	}
	for _, v := range records {
		fmt.Println(v[2])
	}

	return records, err
}
