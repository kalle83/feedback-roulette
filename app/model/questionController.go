package model

import "github.com/gin-gonic/gin"

type QuestionController interface {
	GetAllQuestions(c *gin.Context)
}
