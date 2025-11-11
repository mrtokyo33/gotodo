package repositories

import (
	"encoding/json"
	stderrors "errors"
	"os"
	"path/filepath"

	"github.com/mrtokyo33/todo/pkg/errors"
	"github.com/mrtokyo33/todo/pkg/models"
)

type JSONRepository struct {
	filePath string
}

func NewJSONRepository(fileName string) *JSONRepository {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(errors.HomeDirError(err))
	}

	filePath := filepath.Join(homeDir, fileName)

	return &JSONRepository{
		filePath: filePath,
	}
}

func (j *JSONRepository) GetAll() ([]models.Task, error) {
	data, err := os.ReadFile(j.filePath)

	if stderrors.Is(err, os.ErrNotExist) || len(data) == 0 {
		return []models.Task{}, nil
	}

	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (j *JSONRepository) SaveAll(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.filePath, data, 0600)
}

func (j *JSONRepository) GetByID(id int) (*models.Task, error) {
	tasks, err := j.GetAll()
	if err != nil {
		return nil, err
	}

	for i, task := range tasks {
		if task.ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.ErrTaskNotFound
}
