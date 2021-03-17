package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Steps
// ==== 1. Read the CSV. It could be set via a flag, with default to problems.csv
// ==== 2. Show the number of right answers
// ==== 3. Add a timer. Defaults to 30 seconds. Users should be asked to
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
	var (
		// flags
		csvFile   string
		timeLimit int

		score int
	)

	flag.StringVar(&csvFile, "csv", "problems.csv", "CSV file to read from")
	flag.IntVar(&timeLimit, "t", 30, "Time limit to finish the test")
	flag.Parse()

	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Error: file %q not found.", csvFile)
	}
	defer f.Close()

	fmt.Print("Press Enter to start the quiz. A timer will start after.")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	doneCh := make(chan bool)

	go func() {
		score = askQuestions(f)
		doneCh <- true
	}()

	select {
	case <-doneCh:
	case <-timer.C:
		fmt.Println("\nTime's up!")
	}

	fmt.Printf("Your final score: %d", score)
}

func askQuestions(f io.Reader) int {
	var score int

	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Println(err)

			break
		}

		question := record[0]
		answer := record[1]

		var userInput string
		fmt.Printf("%s: ", question)
		fmt.Scanln(&userInput)

		if userInput == answer {
			score++
		}
	}

	return score
}
