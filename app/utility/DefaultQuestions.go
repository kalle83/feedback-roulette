package utility

import "kalle83/feedback-roulette/app/model"

func GetDefaultQuestions() (questions []model.Question) {

	questions = []model.Question{
		{
			Text: "Das hat mir nicht gefallen!",
		},
		{
			Text: "Das kam zu kurz!",
		},
		{
			Text: "Das könnte man besser machen!",
		},
		{
			Text: "Das war super!",
		},
		{
			Text: "Das nehme ich mit!",
		},
		{
			Text: "Das wünsche ich mir für die Zukunft",
		},
	}
	return questions
}
