package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Mark a task as pending",
	Long:  `Mark a task as pending`,
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

			err = taskManager.SetTaskStatus(id, false)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error marking task %d as pending: %v\n", id, err)
				continue
			}

			fmt.Printf("Task %d marked as pending.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)

}
