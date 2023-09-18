package task

import "rest_api/internal/app/model"

type InMemoryTaskRepository struct {
	tasks []model.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{}
}

func (r *InMemoryTaskRepository) GetAllTasks() ([]model.Task, error) {
	return r.tasks, nil
}
func (r *InMemoryTaskRepository) AddTask(task model.Task) error {
	r.tasks = append(r.tasks, task)
	return nil
}
