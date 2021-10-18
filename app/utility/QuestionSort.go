package utility

import (
	"kalle83/feedback-roulette/app/model"
	"sort"
)

func SortQuestionById(questions []model.Question) {
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].Id < questions[j].Id
	})
}

func SortQuestionByText(questions []model.Question) {
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].Text < questions[j].Text
	})
}

func GetHighestId(questions []model.Question) int64 {
	SortQuestionById(questions)
	lastQuestionOfSlice := questions[len(questions)-1]

	return int64(lastQuestionOfSlice.Id)
}
