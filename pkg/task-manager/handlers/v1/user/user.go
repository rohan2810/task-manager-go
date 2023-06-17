package user

import (
	"context"
	"fmt"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
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

func (s *Service) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}

	userT := &models.User{
		Username: req.GetUser().Username,
		Fullname: req.GetUser().Fullname,
		Email:    req.GetUser().Email,
		Role:     req.GetUser().Role,
	}

	user, err := s.DB.CreateUser(ctx, userT)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &v1.User{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *Service) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	var err error
	// TODO: implement request fields

	users, err := s.DB.ListUser(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	usersA := make([]*v1.User, len(users))
	for _, user := range users {
		usersA = append(usersA, &v1.User{
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email,
			Role:     user.Role,
		})
	}
	return &v1.ListUsersResponse{
		Users: usersA,
	}, nil
}
