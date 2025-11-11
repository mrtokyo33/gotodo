package utils

import (
	"fmt"

	"github.com/mrtokyo33/todo/pkg/models"
)

func TaskStatus(completed bool) string {
	if completed {
		return Highlight("[✓]")
	}
	return Primary("[ ]")
}

func TaskLine(task models.Task) string {
	status := TaskStatus(task.Completed)
	id := Primary(fmt.Sprintf("%d", task.ID))
	return fmt.Sprintf("%s %s: %s", status, id, task.Title)
}

func SuccessMessage(msg string) string {
	return Primary(msg)
}

func ErrorMessage(msg string) string {
	return Error(msg)
}

func WarningMessage(msg string) string {
	return Warning(msg)
}

func TaskAdded(title string) string {
	return SuccessMessage(fmt.Sprintf("Task successfully added: \"%s\"", title))
}

func TaskRemoved(id int) string {
	return SuccessMessage(fmt.Sprintf("Task %d removed successfully", id))
}

func TaskMarkedDone(id int) string {
	return SuccessMessage(fmt.Sprintf("Task %d marked as done", id))
}

func TaskMarkedPending(id int) string {
	return SuccessMessage(fmt.Sprintf("Task %d marked as pending", id))
}

func EmptyListMessage() string {
	return Warning("\nYour to-do list is empty! Go add some tasks.")
}

func ListTitle() string {
	return Title("\nCurrent To-Do List:")
}
