package postgres

import (
	"context"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
)

type DBService interface {
	Close() error
	DBTaskService
	DBUserService
}

type DBTaskService interface {
}
type DBUserService interface {
	CreateUserTable(ctx context.Context) error
	CreateUser(ctx context.Context, u *models.User) (*models.User, error)
	GetUser(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, username string) error
	ListUser(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, u *models.User) (*models.User, error)
}
