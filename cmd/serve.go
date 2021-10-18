package cmd

import (
	"kalle83/feedback-roulette/app/handler"

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

	handler.Start()

}
