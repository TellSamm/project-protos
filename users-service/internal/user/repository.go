package user

import (
	"github.com/TellSamm/users-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserByID(id string) error
	GetTasksByUserID(userID uuid.UUID) ([]models.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *userRepository) GetTasksByUserID(userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task
	err := repo.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (repo *userRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) UpdateUser(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *userRepository) DeleteUserByID(id string) error {
	return repo.db.Delete(&models.User{}, "id = ?", id).Error
}
