package repository

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kalle83/feedback-roulette/app/model"
	"kalle83/feedback-roulette/app/utility"
	"os"

	log "github.com/sirupsen/logrus"
)

type fileQuestionRepository struct {
	Path string
}

func NewFileQuestionRepository(path string) model.QuestionRepository {
	return &fileQuestionRepository{Path: path}
}

func (qr *fileQuestionRepository) FindAll() (questions []model.Question, err error) {

	log.Debug("Start Reading...")
	jsonFile, err := os.Open(qr.Path)
	if err != nil {
		log.Error(err)
		return nil, errors.New("CouldNotReadFile")
	}

	log.Debug("Successfully Openend ", qr.Path)
	defer jsonFile.Close()

	byteArray, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(err)
		return nil, errors.New("CouldNotReadByteArray")
	}

	err = json.Unmarshal(byteArray, &questions)
	if err != nil {
		log.Error(err)
		return nil, errors.New("CouldNotUnmarshalByteArray: verify JSON format")
	}

	log.Debugf("Read %d Questions from: %s", len(questions), qr.Path)

	return questions, nil
}

func (qr *fileQuestionRepository) FindById(id int64) (*model.Question, error) {

	questions, err := qr.FindAll()
	if err != nil {
		log.Error(err)
		return nil, errors.New("CouldNotLoadQuestions")
	}

	for _, question := range questions {
		if question.Id == id {
			log.Debugf("Found question with id %d", id)
			return &question, nil
		}
	}

	log.Debugf("Could Not found a question with id %d", id)
	return nil, nil
}

func (qr *fileQuestionRepository) Update(question *model.Question) (*model.Question, error) {

	questions, err := qr.FindAll()
	if err != nil {
		return nil, err
	}

	for k, q := range questions {
		if q.Id == question.Id {
			log.Debugf("Found question with ID: %d", q.Id)
			questions[k].Text = question.Text
			log.Debugf("Update Question with ID: %d", q.Id)
		}
	}

	qr.writeQuestionJson(questions)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (qr *fileQuestionRepository) Save(question *model.Question) (*model.Question, error) {

	questions, err := qr.FindAll()
	if err != nil {
		return nil, err
	}

	highestID := utility.GetHighestId(questions)
	if highestID == -1 {
		return nil, errors.New("CouldNotFindHighestID")
	}

	question.Id = highestID + 1
	questions = append(questions, *question)

	err = qr.writeQuestionJson(questions)
	if err != nil {
		return nil, err
	}

	log.Infof("Store new Question with ID %d and Text %s", question.Id, question.Text)

	return question, nil
}

func (qr *fileQuestionRepository) Delete(question *model.Question) error {

	questions, err := qr.FindAll()
	if err != nil {
		return err
	}

	for k, q := range questions {
		if q.Id == question.Id {
			log.Debugf("Found question with ID: %d", q.Id)

			questions[k] = questions[len(questions)-1]
			newQuestions := questions[:len(questions)-1]

			qr.writeQuestionJson(newQuestions)
			if err != nil {
				return err
			}
			log.Debugf("Remove Question with ID: %d", q.Id)
		}
	}
	return nil
}

func (qr *fileQuestionRepository) writeQuestionJson(questions []model.Question) error {

	file, err := json.MarshalIndent(questions, "", "\t")
	if err != nil {
		log.Error(err)
		return errors.New("CouldNotMarshalQuestions")
	}

	err = ioutil.WriteFile(qr.Path, file, 0644)
	if err != nil {
		log.Error(err)
		return errors.New("CouldNotWriteQuestionFile")
	}
	log.Debugf("Write Questions File with %d entries", len(questions))
	return nil
}
