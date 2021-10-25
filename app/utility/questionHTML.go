package utility

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomQuestion(all bool) string {
	questions := GetDefaultQuestions()

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6)

	if all {
		for k, question := range questions {
			if k == number {
				return fmt.Sprintf("%d) \t %s \n", k+1, question.Text)
			}
		}
	}

	return questions[number].Text
}
