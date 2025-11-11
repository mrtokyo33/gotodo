package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all tasks",
	Long:  `View all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println("ERROR: Task Manager not initialized")
			return
		}

		tasks, err := taskManager.ListTasks()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("\nYour to-do list is empty! Go add some tasks.")
			return
		}

		fmt.Println("\nCurrent To-Do List:")
		for _, task := range tasks {
			status := "[ ]"
			if task.Completed {
				status = "[x]"
			}

			fmt.Printf("%s %d: %s\n", status, task.ID, task.Title)
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
