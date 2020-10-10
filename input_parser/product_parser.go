package input_parser

import (
	"io"
	"log"
	"promotionengine/model"
	"strconv"
)

func InitializeProductDB(csvFileName string) (map[string]model.Product, error) {

	mapProducts := make(map[string]model.Product)
	r, err := openCsvFile(csvFileName)
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

		priceProduct, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		p := model.Product{
			ProductName: record[0],
			Price:       priceProduct,
		}
		mapProducts[record[0]] = p
	}
	return mapProducts, nil
}
