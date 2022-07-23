package service

import (
	"app/models"
	"app/pkg/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(group models.Group) (models.Group, error) {
	return s.repo.CreateGroup(group)
}

func (s *GroupService) GetAllGroups() ([]models.Group, error) {
	return s.repo.GetAllGroups()
}

func (s *GroupService) GetGroupByID(groupID int) (models.Group, error) {
	return s.repo.GetGroupByID(groupID)
}

func (s *GroupService) DeleteGroup(groupID int) error {
	return s.repo.DeleteGroup(groupID)
}

func (s *GroupService) UpdateGroup(groupID int, group models.Group) (models.Group, error) {
	return s.repo.UpdateGroup(groupID, group)
}
