package cmd

import (
	"fmt"
	"strconv"

	"github.com/mrtokyo33/gotodo/pkg/errors"
	"github.com/mrtokyo33/gotodo/pkg/utils"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as done",
	Long:  `Mark a task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskManager == nil {
			fmt.Println(utils.ErrorMessage(errors.ErrTaskManagerNotInit.Error()))
			return
		}

		for _, argID := range args {
			id, err := strconv.Atoi(argID)
			if err != nil {
				fmt.Println(utils.ErrorMessage(errors.InvalidIDError(argID).Error()))
				continue
			}

			err = taskManager.SetTaskStatus(id, true)

			if err != nil {
				fmt.Println(utils.ErrorMessage(fmt.Sprintf("error completing task %d: %v", id, err)))
				continue
			}

			fmt.Println(utils.TaskMarkedDone(id))
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
