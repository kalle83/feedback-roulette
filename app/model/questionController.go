package model

import "github.com/gin-gonic/gin"

type QuestionController interface {
	GetAllQuestions(c *gin.Context)
	CreateQuestion(c *gin.Context)
	DeleteQuestion(c *gin.Context)
	// GetQuestionById(c *gin.Context)
}
