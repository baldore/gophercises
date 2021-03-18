// Steps
// ==== 1. Read the CSV. It could be set via a flag, with default to problems.csv
// ==== 2. Show the number of right answers
// ==== 3. Add a timer. Defaults to 30 seconds. Users should be asked to
//    press Enter before the timer starts.
//
// Bonus
// ==== 1. Add string trimming and cleanup to help ensure that correct
//    answers with extra whitespace, capitalization, etc are not considered
//    incorrect.
//    Hint: Check out the strings package.
// 2. Add an option (a new flag) to shuffle the quiz order each time it is
//    run.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"./quiz"
)

func main() {
	var (
		csvFile   string
		timeLimit int
		random    bool
	)

	flag.StringVar(&csvFile, "csv", "problems.csv", "CSV file to read from")
	flag.IntVar(&timeLimit, "t", 30, "Time limit to finish the test")
	flag.BoolVar(&random, "r", false, "Randomize the questions order")
	flag.Parse()

	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatalf("Error: file %q not found.", csvFile)
	}
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	qs := recordsToQuestions(records)
	qz := quiz.New(qs)

	fmt.Print("Press Enter to start the quiz. A timer will start after.")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	doneCh := make(chan bool)

	go func() {
		qz.Start()
		doneCh <- true
	}()

	select {
	case <-doneCh:
	case <-timer.C:
		fmt.Println("\nTime's up!")
	}

	fmt.Printf("Your final score: %d", qz.Score)
}

func recordsToQuestions(r [][]string) quiz.Questions {
	qs := quiz.Questions{}
	for _, record := range r {
		qs = append(qs, quiz.Question{
			Question: record[0],
			Answer:   record[1],
		})
	}

	return qs
}
