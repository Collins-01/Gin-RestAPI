package task

import "rest_api/internal/app/model"

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	AddTask(task model.Task) error
}
