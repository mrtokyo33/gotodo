package cmd

import (
	"fmt"
	"strconv"

	"github.com/mrtokyo33/gotodo/pkg/errors"
	"github.com/mrtokyo33/gotodo/pkg/utils"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Mark a task as pending",
	Long:  `Mark a task as pending`,
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

			err = taskManager.SetTaskStatus(id, false)

			if err != nil {
				fmt.Println(utils.ErrorMessage(fmt.Sprintf("error marking task %d as pending: %v", id, err)))
				continue
			}

			fmt.Println(utils.TaskMarkedPending(id))
		}
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
