package postgres

import (
	"context"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
)

type DBService interface {
	Close() error
	DBTaskService
	DBUserService
}

type DBTaskService interface {
	CreateTaskTable(ctx context.Context) error
	CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskResponse, error)
	GetTask(ctx context.Context, req *v1.GetTaskRequest) (*v1.Task, error)
	UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*v1.Task, error)
	ListTask(ctx context.Context, req *v1.ListTasksRequest) (*v1.ListTasksResponse, error)
	DeleteTask(ctx context.Context, req *v1.DeleteTaskRequest) error
}

type DBUserService interface {
	CreateUserTable(ctx context.Context) error
	CreateUser(ctx context.Context, u *models.User) (*models.User, error)
	GetUser(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, username string) error
	ListUser(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, u *models.User) (*models.User, error)
}
