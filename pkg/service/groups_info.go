package service

import (
	"app/models"
	"app/pkg/repository"
)

type GroupInfoService struct {
	repo repository.GroupInfo
}

func NewGroupInfoService(repo repository.GroupInfo) *GroupInfoService {
	return &GroupInfoService{repo: repo}
}

func (s *GroupInfoService) GetGroupInfoById(groupId int) (models.GroupInfo, error) {
	return s.repo.GetGroupInfoById(groupId)
}
