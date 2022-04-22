package helper

import (
	"encoding/csv"
	"os"

	"github.com/rs/zerolog/log"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Error().Err(err).Msg("File does not exist")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	if len(records) == 0 {

	}
	if err != nil {
		log.Error().Err(err).Msg("Unable to read file")
	}

	return records
}
