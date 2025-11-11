package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Long:  `Mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println("ERROR: Task Manager not initialized")
			return
		}

		for _, argID := range args {
			id, err := strconv.Atoi(argID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid ID '%s'. Must be a number.\n", argID)
				continue
			}

			err = taskManager.SetTaskStatus(id, true)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error completing task %d: %v\n", id, err)
				continue
			}

			fmt.Printf("Task %d marked as done.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

}
