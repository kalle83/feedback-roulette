package handler

import (
	"fmt"
	"kalle83/feedback-roulette/app/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type server struct {
	Port               int
	QuestionController model.QuestionController
}

func NewServer(port int, questionController model.QuestionController) model.Server {

	s := server{
		Port:               port,
		QuestionController: questionController,
	}

	return &s
}

func (s *server) Start() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "healthy: true")
	})

	router.Use(static.Serve("/", static.LocalFile("./ui/dist", false)))
	v1 := router.Group("/api/v1")
	{
		v1.GET("/question", s.QuestionController.GetAllQuestions)
		v1.POST("/question", s.QuestionController.CreateQuestion)
		v1.DELETE("/question/:id", s.QuestionController.DeleteQuestion)
		// api.GET("/question/:id", user.GetUser)
		// api.PUT("/question/:id", user.UpdateUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "{error: no route}")
		// c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":" + fmt.Sprint(s.Port))
}
