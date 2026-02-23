package service

import (
	"github.com/Pashhhka/task-manager/internal/models"
	"github.com/Pashhhka/task-manager/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task *models.Task) error {
	task.Status = "pending"
	return s.repo.Create(task)
}

func (s *TaskService) GetByUser(userID int) ([]models.Task, error) {
	return s.repo.GetByUser(userID)
}

func (s *TaskService) Update(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id, userID int) error {
	return s.repo.Delete(id, userID)
}
