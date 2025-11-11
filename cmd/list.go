package cmd

import (
	"fmt"

	"github.com/mrtokyo33/gotodo/pkg/errors"
	"github.com/mrtokyo33/gotodo/pkg/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all tasks",
	Long:  `View all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println(utils.ErrorMessage(errors.ErrTaskManagerNotInit.Error()))
			return
		}

		tasks, err := taskManager.ListTasks()

		if err != nil {
			fmt.Println(utils.ErrorMessage(fmt.Sprintf("error loading tasks: %v", err)))
			return
		}

		if len(tasks) == 0 {
			fmt.Println(utils.EmptyListMessage())
			return
		}

		fmt.Println(utils.ListTitle())
		for _, task := range tasks {
			fmt.Println(utils.TaskLine(task))
		}
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
