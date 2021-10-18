package controller

import (
	"errors"
	"kalle83/feedback-roulette/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
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

// GetAllUser Endpoint
func GetAllQuestions(c *gin.Context) {
	questions := []model.Question{}
	c.JSON(http.StatusOK, gin.H{"status": "success", "question": &questions})
}
