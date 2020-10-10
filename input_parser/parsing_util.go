package input_parser

import (
	"encoding/csv"
	"log"
	"os"
)

func openCsvFile(csvFile string) (*csv.Reader, error) {
	// Open the file
	csvfile, err := os.Open(csvFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return nil, err
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	return r, nil
}
