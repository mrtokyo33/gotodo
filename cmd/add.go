package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Quickly create a new task",
	Long:  `Quickly create a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println("ERROR: Task Manager not initialized")
			return
		}

		title := strings.Join(args, " ")
		err := taskManager.AddTask(title)

		if err != nil {
			fmt.Printf("error adding task: %v\n", err)
			return
		}

		fmt.Printf("Task successfully added: \"%s\"\n", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
