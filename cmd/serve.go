package cmd

import (
	"kalle83/feedback-roulette/app/controller"
	"kalle83/feedback-roulette/app/handler"
	"kalle83/feedback-roulette/app/repository"
	"kalle83/feedback-roulette/app/service"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Feedback-Roulette API Server",
	Long:  "Start Feedback-Roulette API Server",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Serve() {

	questionRepository := repository.NewFileQuestionRepository("questions.json")
	questionService := service.NewQuestionService(questionRepository)
	questionController := controller.NewQuestionController(questionService)
	server := handler.NewServer(80, questionController)

	server.Start()

}
