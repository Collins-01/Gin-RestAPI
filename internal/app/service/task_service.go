package service

import (
	"rest_api/internal/app/model"
	repository "rest_api/internal/app/repository/task"
)

type TaskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{taskRepository: repo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.taskRepository.GetAllTasks()
}

func (s *TaskService) AddTask(task model.Task) error {
	return s.taskRepository.AddTask(task)
}
