package errors

import "fmt"

var (
	ErrTaskNotFound        = fmt.Errorf("task not found")
	ErrHomeDirNotFound     = fmt.Errorf("user's home directory could not be found")
	ErrTaskManagerNotInit  = fmt.Errorf("task manager not initialized")
	ErrInvalidID           = fmt.Errorf("invalid ID, must be a number")
)

func TaskNotFoundByID(id int) error {
	return fmt.Errorf("task with ID: %d not found", id)
}

func TaskNotFoundForRemoval(id int) error {
	return fmt.Errorf("task with ID: %d not found for removal", id)
}

func HomeDirError(err error) error {
	return fmt.Errorf("%w: %v", ErrHomeDirNotFound, err)
}

func InvalidIDError(id string) error {
	return fmt.Errorf("%w: '%s'", ErrInvalidID, id)
}
