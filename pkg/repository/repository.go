package repository

import (
	"app/models"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Group
	GroupInfo
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Group:     NewGroupPostgres(db),
		GroupInfo: NewGroupInfoPostgres(db),
		User:      NewUserPostgres(db),
	}
}
