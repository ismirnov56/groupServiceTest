package service

import (
	"app/models"
	"app/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(userID int) (models.User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *UserService) DeleteUser(userID int) error {
	return s.repo.DeleteUser(userID)
}

func (s *UserService) UpdateUser(userID int, user models.User) (models.User, error) {
	return s.repo.UpdateUser(userID, user)
}
