package main

import (
	"fmt"

	"github.com/BenAyedSeeksAI/project-mongo-db/db"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A simple CLI app to say hello",
	Long:  "A simple CLI app built using Cobra for learning purposes.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI app!")
	},
}
var GetCarsCommand = &cobra.Command{
	Use:   "get-cars",
	Short: "display list of cars",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here are the available cars.")
		db.DBGetCars()
	},
}

// Execute initializes and starts the application
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	// Add subcommands to the root command here
	RootCmd.AddCommand(GetCarsCommand)
}
