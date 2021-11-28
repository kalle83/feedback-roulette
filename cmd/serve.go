package cmd

import (
	"kalle83/feedback-roulette/app/controller"
	"kalle83/feedback-roulette/app/handler"
	"kalle83/feedback-roulette/app/repository"
	"kalle83/feedback-roulette/app/service"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Feedback-Roulette API Server",
	Long:  "Start Feedback-Roulette API Server",
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Panic("could not read port")
		}

		Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntP("port", "p", 80, "Set your port")
}

func Serve(port int) {

	questionRepository := repository.NewFileQuestionRepository("questions.json")
	questionService := service.NewQuestionService(questionRepository)
	questionController := controller.NewQuestionController(questionService)
	server := handler.NewServer(port, questionController)

	server.Start()

}
