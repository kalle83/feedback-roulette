package utility

import (
	"fmt"
	"kalle83/feedback-roulette/app/model"
)

func ListAllQuestion(questions []model.Question) {
	fmt.Println("============================")
	for i, question := range questions {
		fmt.Printf("Position: %d \t ID: %d \t Text: %s\n", i, question.Id, question.Text)
	}
	fmt.Println("============================")
}

func ListQuestion(question *model.Question) {
	fmt.Println("============================")
	fmt.Printf("Position: 1 \t ID: %d \t Text: %s\n", question.Id, question.Text)
	fmt.Println("============================")
}
