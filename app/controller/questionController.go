package controller

import (
	"errors"
	"kalle83/feedback-roulette/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// https://www.agiratech.com/building-restful-api-service-golang-using-gin-gonic
var (
	errNotExist        = errors.New("Users are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the user insertion")
	errUpdationFailed  = errors.New("Error in the user updation")
	errDeletionFailed  = errors.New("Error in the user deletion")
)

type questionController struct {
	QuestionService model.QuestionService
}

func NewQuestionController(questionService model.QuestionService) model.QuestionController {
	return &questionController{QuestionService: questionService}
}

// GetAllUser Endpoint
func (qc *questionController) GetAllQuestions(c *gin.Context) {
	questions, err := qc.QuestionService.FindAll()
	if err != nil {
		log.Error(err)
		c.AbortWithError(500, err)
	}
	c.JSON(http.StatusOK, questions)
}
