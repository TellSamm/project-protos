package task

import (
	"github.com/TellSamm/tasks-service/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (repo *taskRepository) CreateTask(task *models.Task) error {
	return repo.db.Create(task).Error
}

func (repo *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := repo.db.Find(&tasks).Error
	return tasks, err
}

func (repo *taskRepository) GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	err := repo.db.First(&task, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &task, err
}

func (repo *taskRepository) UpdateTask(task *models.Task) error {
	return repo.db.Save(task).Error
}

func (repo *taskRepository) DeleteTaskByID(id string) error {
	return repo.db.Delete(&models.Task{}, "id = ?", id).Error
}
