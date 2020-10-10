package promotion_rules

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"promotionengine/model"
)

var rules_map = map[string]interface{}{
	"nOfSame":          NofSame,
	"combinationOfTwo": CombinationOfTwo,
}

func Parse_rules_file_csv(csv_file string) ([]model.Rules, error) {

	r, err := openCsvFile("rules.csv")
	if err != nil {
		return nil, err
	}

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		//find the rule in the map
		if _, ok := rules_map[record[0]]; ok {
			//check if the number of params is the same as expected

			//add to the list of promotion rules

		} else {
			return nil, errors.New("Rule " + record[0] + " is undefined.")
		}

	}
	return nil, nil
}

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
