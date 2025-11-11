package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task permanently",
	Long:  `Delete a task permanently`,
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

			err = taskManager.RemoveTask(id)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error removing task %d: %v\n", id, err)
				continue
			}

			fmt.Printf("Task %d removed successfully.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

}
