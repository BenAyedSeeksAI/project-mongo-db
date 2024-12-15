package main

import (
	"fmt"

	"github.com/BenAyedSeeksAI/project-mongo-db/controller"

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
	Run:   controller.GetCarsCmd,
}
var AddCarCommand = &cobra.Command{
	Use:   "add-car",
	Short: "add a car object",
	Run:   controller.AddCarCmd,
}

// Execute initializes and starts the application
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	// Add subcommands to the root command here
	RootCmd.AddCommand(GetCarsCommand)
	RootCmd.AddCommand(AddCarCommand)
}
