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

func (s *UserService) GetUserById(userId int) (models.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func (s *UserService) UpdateUser(userId int, user models.User) (models.User, error) {
	return s.repo.UpdateUser(userId, user)
}
