package service

import (
	"app/models"
	"app/pkg/repository"
	"context"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(ctx context.Context, group models.Group) (models.Group, error) {
	return s.repo.CreateGroup(ctx, group)
}

func (s *GroupService) GetAllGroups(ctx context.Context) ([]models.Group, error) {
	return s.repo.GetAllGroups(ctx)
}

func (s *GroupService) GetGroupByID(ctx context.Context, groupID int) (models.Group, error) {
	return s.repo.GetGroupByID(ctx, groupID)
}

func (s *GroupService) DeleteGroup(ctx context.Context, groupID int) error {
	return s.repo.DeleteGroup(ctx, groupID)
}

func (s *GroupService) UpdateGroup(ctx context.Context, groupID int, group models.Group) (models.Group, error) {
	return s.repo.UpdateGroup(ctx, groupID, group)
}
