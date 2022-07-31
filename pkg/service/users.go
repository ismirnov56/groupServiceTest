package service

import (
	"app/models"
	"app/pkg/repository"
	"context"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, userID int) (models.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	return s.repo.DeleteUser(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, userID int, user models.User) (models.User, error) {
	return s.repo.UpdateUser(ctx, userID, user)
}
