package task

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/repo/postgres"
)

type Service struct {
	v1.UnimplementedTaskServiceServer
	DB  postgres.DBService
	ctx context.Context
}

func NewService(db postgres.DBService, ctx context.Context) *Service {
	return &Service{
		DB:  db,
		ctx: ctx,
	}
}

func (s *Service) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskResponse, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}

	task, err := s.DB.CreateTask(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &v1.CreateTaskResponse{
		Done:   true,
		TaskId: task.TaskId,
	}, nil
}

func (s *Service) GetTask(ctx context.Context, req *v1.GetTaskRequest) (*v1.Task, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}
	task, err := s.DB.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Service) UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*v1.Task, error) {
	var err error
	if req == nil {
		fmt.Println("empty request")
	}
	task, err := s.DB.UpdateTask(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return task, nil
}

func (s *Service) DeleteTask(ctx context.Context, request *v1.DeleteTaskRequest) (*emptypb.Empty, error) {
	var err error
	if request == nil {
		fmt.Println("empty request")
	}

	err = s.DB.DeleteTask(ctx, request)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) ListTasks(ctx context.Context, request *v1.ListTasksRequest) (*v1.ListTasksResponse, error) {
	var err error
	tasks, err := s.DB.ListTask(ctx, request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return tasks, nil
}
