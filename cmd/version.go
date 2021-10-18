package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vrsionCmd)
}

var vrsionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of feedback roulette",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},
}
