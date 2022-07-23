package service

import (
	"app/models"
	"app/pkg/repository"
	"context"
)

type Group interface {
	CreateGroup(group models.Group) (models.Group, error)
	GetAllGroups() ([]models.Group, error)
	GetGroupByID(groupID int) (models.Group, error)
	DeleteGroup(groupID int) error
	UpdateGroup(groupID int, group models.Group) (models.Group, error)
}

type GroupInfo interface {
	GetGroupInfoByID(ctx context.Context, groupID int) (models.GroupInfo, error)
}

type User interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID int) (models.User, error)
	DeleteUser(userID int) error
	UpdateUser(userID int, user models.User) (models.User, error)
}

type Service struct {
	Group
	GroupInfo
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Group:     NewGroupService(repos.Group),
		GroupInfo: NewGroupInfoService(repos.GroupInfo),
		User:      NewUserService(repos.User),
	}
}
