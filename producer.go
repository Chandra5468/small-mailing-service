package main

import (
	"encoding/csv"
	"os"
)

func loadReceipient(filepath string, ch chan<- Receipient) error {
	// Read emails.csv
	defer close(ch)
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, record := range records[1:] { // first record is skipped as it has headers
		// fmt.Println(record)
		// Send to consumers -> Channels
		ch <- Receipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil
}
