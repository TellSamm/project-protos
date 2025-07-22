package task

import (
	"errors"
	"github.com/TellSamm/tasks-service/internal/models"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTaskByID(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	if task.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return s.repo.CreateTask(task)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskByID(id string) (*models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *taskService) DeleteTaskByID(id string) error {
	return s.repo.DeleteTaskByID(id)
}
