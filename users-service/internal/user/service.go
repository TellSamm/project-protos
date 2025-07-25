package user

import (
	"github.com/TellSamm/users-service/internal/models"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserByID(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUserByID(id string) error {
	return s.repo.DeleteUserByID(id)
}
