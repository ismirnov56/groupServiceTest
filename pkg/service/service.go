package service

import (
	"app/models"
	"app/pkg/repository"
)

type Group interface {
	CreateGroup(group models.Group) (models.Group, error)
	GetAllGroups() ([]models.Group, error)
	GetGroupById(groupId int) (models.Group, error)
	DeleteGroup(groupId int) error
	UpdateGroup(groupId int, group models.Group) (models.Group, error)
}

type GroupInfo interface {
	GetGroupInfoById(groupId int) (models.GroupInfo, error)
}

type User interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(userId int) (models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user models.User) (models.User, error)
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
