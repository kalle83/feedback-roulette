package cmd

import (
	"fmt"
	"kalle83/feedback-roulette/app/utility"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listAll bool

var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Show a random feedback - question",
	Long:  "This command show you a random question of the default question set",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomQuestion()
	},
}

func init() {
	questionCmd.Flags().BoolVarP(&listAll, "all", "a", false, "show all commands")
	rootCmd.AddCommand(questionCmd)
}

func getRandomQuestion() {

	questions := utility.GetDefaultQuestions()

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6)

	if listAll {
		for k, question := range questions {
			if k == number {
				color.Green(fmt.Sprintf("%d) \t %s \n", k+1, question.Text))
				continue
			}
			color.Blue(fmt.Sprintf("%d) \t %s \n", k+1, question.Text))
		}
		return
	}

	fmt.Println(questions[number].Text)
}
