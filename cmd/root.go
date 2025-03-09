package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cirr",
	Short: "cirr retrieves cloud IP ranges from various providers",
	Long:  `A CLI tool to fetch and display IP ranges from cloud providers like AWS, Azure, and Google Cloud.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to cirr! Use --help for available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
