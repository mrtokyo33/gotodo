package cmd

import (
	"os"

	"github.com/mrtokyo33/todo/pkg/repositories"
	"github.com/mrtokyo33/todo/pkg/services"
	"github.com/spf13/cobra"
)

var (
	taskManager *services.TaskManager
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "A command-line interface tool to effectively manage your daily tasks",
	Long: `gotodo is a lightweight and efficient command-line application for task management 
	It allows users to quickly add list complete and remove tasks directly from the terminal 
	Perfect for anyone looking for a straightforward way to organize their to-do items`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		repo := repositories.NewJSONRepository(".gotodo.json")

		taskManager = services.NewTaskManager(repo)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
