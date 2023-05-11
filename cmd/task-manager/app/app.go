package app

import (
	"context"
	"google.golang.org/grpc"
	"rohan.com/task-manager/pkg/task-manager/repo/postgres"
)

type app struct {
	dbService  postgres.DBService
	grpcServer *grpc.Server
	ctx        context.Context
}

func (a *app) createDBtables() error {
	err := a.dbService.CreateUserTable(context.Background())
	if err != nil {
		return err
	}
	return nil
}
