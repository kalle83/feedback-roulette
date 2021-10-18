package service

import (
	"errors"
	"kalle83/feedback-roulette/app/model"

	log "github.com/sirupsen/logrus"
)

type questionService struct {
	repo model.QuestionRepository
}

func NewQuestionService(r model.QuestionRepository) model.QuestionService {
	return &questionService{repo: r}
}

func (qs *questionService) FindAll() (questions []model.Question, err error) {

	questions, err = qs.repo.FindAll()
	if err != nil {
		log.Error(err)
		return nil, errors.New("CouldNotRetriveDataFromReop")
	}

	//Write test if question are 0

	return questions, nil
}

func (qs *questionService) FindById(id int64) (*model.Question, error) {
	question, err := qs.repo.FindById(id)
	if err != nil {
		log.Error("CouldNotRetriveQuestionFromRepo")
		return nil, err
	}

	if question == nil {
		log.Errorf("No Question with ID %d", id)
		return nil, err
	}

	return question, nil
}

func (qs *questionService) Create(questionText string) (model.Question, error) {

	question, err := qs.repo.Save(&model.Question{Text: questionText})
	if err != nil {
		log.Error(err)
		return model.Question{}, errors.New("couldNotStoreQuestion")
	}
	return *question, nil
}

func (qs *questionService) Update(question *model.Question) error {

	_, err := qs.repo.Update(question)
	if err != nil {
		return err
	}

	return nil
}

func (qs *questionService) Delete(question *model.Question) error {

	err := qs.repo.Delete(question)
	if err != nil {
		return err
	}

	return nil
}
