package csvtest

import (
	"encoding/csv"
	"os"
	"log"
)

func CsvExport() {
	records  := [][]string{
		{"id", "name", "zip_code"},
	}
	f, err := os.OpenFile("csvtest/file/test.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush() //バッファ内の」データを書込み
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}




}