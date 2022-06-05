package csvtest

import (
	"encoding/csv"
	"os"
	"log"
)

func CsvExport() {
	// data format
	records  := [][]string{
		{"id", "name", "zip_code"},
	}
	// open
	f, err := os.OpenFile("csvtest/file/test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	// write & create
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}




}