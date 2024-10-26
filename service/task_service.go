package service

import (
	"go-todo-app/model"
	"go-todo-app/repository"
)

type TaskService struct {
	taskRepo repository.ITaskRepository
}
func NewTaskService(taskRepo repository.ITaskRepository) ITaskService {
	return &TaskService{taskRepo}
}
// CreateTask implements ITaskService.
func (t *TaskService) CreateTask(task model.Task) error {
	return t.taskRepo.CreateTask(task)
}

// DeleteTask implements ITaskService.
func (t *TaskService) DeleteTask(id int) error {
	return t.taskRepo.DeleteTask(id)
}

// GetTaskById implements ITaskService.
func (t *TaskService) GetTaskById(id int) (model.Task, error) {
	return t.taskRepo.GetTaskById(id)
}

// GetTasks implements ITaskService.
func (t *TaskService) GetTasks() ([]model.Task, error) {
	return t.taskRepo.GetTasks()
}
// UpdateTask implements ITaskService.
func (t *TaskService) UpdateTask(task model.Task) error {
	return t.taskRepo.UpdateTask(task)
}


