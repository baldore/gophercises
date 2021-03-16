package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

// Steps
// 1. Read the CSV. It could be set via a flag, with default to problems.csv
// 2. Show the number of right answers
// 3. Add a timer. Defaults to 30 seconds. Users should be asked to
//    press Enter before the timer starts.
//
// Bonus
// 1. Add string trimming and cleanup to help ensure that correct
//    answers with extra whitespace, capitalization, etc are not considered
//    incorrect.
//    Hint: Check out the strings package.
// 2. Add an option (a new flag) to shuffle the quiz order each time it is
//    run.

func main() {
	var csvFile string

	flag.StringVar(&csvFile, "csv", "problems.csv", "CSV file to read from")
	flag.Parse()

	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Error: file %q not found.", csvFile)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			log.Println("end!")

			break
		}

		if err != nil {
			log.Println(err)

			break
		}

		log.Println(record)
	}
}
