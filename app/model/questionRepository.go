package model

type QuestionRepository interface {
	FindAll() ([]Question, error)
	FindById(id int64) (*Question, error)
	Save(question *Question) (*Question, error)
	Update(*Question) (*Question, error)
	Delete(*Question) error
}
