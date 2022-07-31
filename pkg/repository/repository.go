package repository

import (
	"app/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Group interface {
	CreateGroup(ctx context.Context, group models.Group) (models.Group, error)
	GetAllGroups(ctx context.Context) ([]models.Group, error)
	GetGroupByID(ctx context.Context, groupID int) (models.Group, error)
	DeleteGroup(ctx context.Context, groupID int) error
	UpdateGroup(ctx context.Context, groupID int, group models.Group) (models.Group, error)
}

type GroupInfo interface {
	GetGroupInfoByID(ctx context.Context, groupID int) (models.GroupInfo, error)
}

type User interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, userID int) (models.User, error)
	DeleteUser(ctx context.Context, userID int) error
	UpdateUser(ctx context.Context, userID int, user models.User) (models.User, error)
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
