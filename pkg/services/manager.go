package services

import (
	"github.com/mrtokyo33/gotodo/pkg/errors"
	"github.com/mrtokyo33/gotodo/pkg/models"
	"github.com/mrtokyo33/gotodo/pkg/repositories"
)

type TaskManager struct {
	Repo repositories.TaskRepository
}

func NewTaskManager(repo repositories.TaskRepository) *TaskManager {
	return &TaskManager{Repo: repo}
}

func (m *TaskManager) AddTask(title string) error {
	tasks, err := m.Repo.GetAll()
	if err != nil {
		return err
	}

	nextID := 1
	if len(tasks) > 0 {
		for _, task := range tasks {
			if task.ID >= nextID {
				nextID = task.ID + 1
			}
		}
	}

	newTask := models.Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}

	tasks = append(tasks, newTask)

	return m.Repo.SaveAll(tasks)
}

func (m *TaskManager) ListTasks() ([]models.Task, error) {
	return m.Repo.GetAll()
}

func (m *TaskManager) SetTaskStatus(id int, status bool) error {
	tasks, err := m.Repo.GetAll()
	if err != nil {
		return err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			if task.Completed == status {
				return nil
			}

			tasks[i].Completed = status
			found = true
			break
		}
	}

	if !found {
		return errors.TaskNotFoundByID(id)
	}

	return m.Repo.SaveAll(tasks)
}

func (m *TaskManager) RemoveTask(id int) error {
	tasks, err := m.Repo.GetAll()
	if err != nil {
		return err
	}

	newTasks := make([]models.Task, 0, len(tasks))
	removed := false
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			removed = true
		}
	}

	if !removed {
		return errors.TaskNotFoundForRemoval(id)
	}

	return m.Repo.SaveAll(newTasks)
}
