package cmd

import (
	"fmt"
	"strconv"

	"github.com/mrtokyo33/todo/pkg/errors"
	"github.com/mrtokyo33/todo/pkg/utils"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a task permanently",
	Long:  `Delete a task permanently`,
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

			err = taskManager.RemoveTask(id)

			if err != nil {
				fmt.Println(utils.ErrorMessage(fmt.Sprintf("error removing task %d: %v", id, err)))
				continue
			}

			fmt.Println(utils.TaskRemoved(id))
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
