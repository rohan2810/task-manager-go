package user

import (
	"context"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/repo/postgres"
)

type Service struct {
	v1.UnimplementedUserServiceServer
	DB  postgres.DBService
	ctx context.Context
}

func NewService(db postgres.DBService, ctx context.Context) *Service {
	return &Service{
		DB:  db,
		ctx: ctx,
	}
}
