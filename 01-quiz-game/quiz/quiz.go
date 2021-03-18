package quiz

import (
	"fmt"
	"strings"
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
