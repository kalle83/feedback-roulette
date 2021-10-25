package model

type QuestionService interface {
	FindAll() ([]Question, error)
	FindById(id int64) (*Question, error)
	Create(questionText string) (Question, error)
	Update(question *Question) error
	Delete(question *Question) error
	DeleteById(id int64) error
}
