package repositories

import "github.com/mrtokyo33/gotodo/pkg/models"

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	SaveAll([]models.Task) error
	GetByID(id int) (*models.Task, error)
}
