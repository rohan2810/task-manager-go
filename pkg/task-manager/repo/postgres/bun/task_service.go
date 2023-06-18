package bun

import (
	"context"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/internal/utils"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
)

func (db *DBClient) CreateTaskTable(ctx context.Context) error {
	_, err := db.DB.NewCreateTable().
		IfNotExists().Model(&models.Task{}).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBClient) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.CreateTaskResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (db *DBClient) GetTask(ctx context.Context, req *v1.GetTaskRequest) (*v1.Task, error) {
	task := &models.Task{}

	err := db.DB.NewSelect().
		Model(task).
		Where("ID = ?", req.Id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListTasksResponse{
		Tasks: utils.TransformMessage(task),
	}, nil
}

func (db *DBClient) UpdateTask(ctx context.Context, req *v1.UpdateTaskRequest) (*v1.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (db *DBClient) ListTask(ctx context.Context, req *v1.ListTasksRequest) (*v1.ListTasksResponse, error) {
	var tasks []*models.Task

	err := db.DB.NewSelect().
		Model(&tasks).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.ListTasksResponse{
		Tasks: utils.TransformMessage(),
	}, nil
}

func (db *DBClient) DeleteTask(ctx context.Context, req *v1.DeleteTaskRequest) error {
	//TODO implement me
	panic("implement me")
}
