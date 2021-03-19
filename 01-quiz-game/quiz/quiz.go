package quiz

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Question struct {
	Question string
	Answer   string
}

type Questions []Question

type Quiz struct {
	Questions Questions
	Score     int
}

func New(q Questions) Quiz {
	return Quiz{
		Questions: q,
		Score:     0,
	}
}

// Starts the quiz from the first question.
func (qz *Quiz) Start() {
	var userInput string

	for _, q := range qz.Questions {
		fmt.Printf("%s: ", q.Question)
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == q.Answer {
			fmt.Println("Correct!")
			qz.Score++
		}
	}
}

// Shuffles the questions.
func (qz *Quiz) Shuffle() {
	qs := qz.Questions

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(qs), func(i, j int) {
		qs[i], qs[j] = qs[j], qs[i]
	})
}
