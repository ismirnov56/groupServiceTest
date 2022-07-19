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

func (s *GroupService) GetGroupById(groupId int) (models.Group, error) {
	return s.repo.GetGroupById(groupId)
}

func (s *GroupService) DeleteGroup(groupId int) error {
	return s.repo.DeleteGroup(groupId)
}
