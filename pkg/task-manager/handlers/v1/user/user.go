package user

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Service) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}
	requestedUsername := req.Username

	user, err := s.DB.GetUser(ctx, requestedUsername)
	if err != nil {
		return nil, err
	}

	return &v1.User{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}
	requestedUsername := req.Username

	err = s.DB.DeleteUser(ctx, requestedUsername)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
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
	user, err := s.DB.UpdateUser(ctx, userT)
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
