package service

import (
	"app/models"
	"app/pkg/repository"
	"context"
)

type GroupInfoService struct {
	repo repository.GroupInfo
}

func NewGroupInfoService(repo repository.GroupInfo) *GroupInfoService {
	return &GroupInfoService{repo: repo}
}

func (s *GroupInfoService) GetGroupInfoByID(ctx context.Context, groupID int) (models.GroupInfo, error) {
	return s.repo.GetGroupInfoByID(ctx, groupID)
}
