package cmd

import (
	"fmt"
	"strings"

	"github.com/mrtokyo33/gotodo/pkg/errors"
	"github.com/mrtokyo33/gotodo/pkg/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Quickly create a new task",
	Long:  `Quickly create a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println(utils.ErrorMessage(errors.ErrTaskManagerNotInit.Error()))
			return
		}

		title := strings.Join(args, " ")
		err := taskManager.AddTask(title)

		if err != nil {
			fmt.Println(utils.ErrorMessage(fmt.Sprintf("error adding task: %v", err)))
			return
		}

		fmt.Println(utils.TaskAdded(title))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
