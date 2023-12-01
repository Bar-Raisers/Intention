package main

import (
	"context"

	tasks "github.com/bar-raisers/intention/scheduling/contracts/tasks/v1"
)

type TasksService struct{}

func NewTasksService() *TasksService {
	return &TasksService{}
}

func (service *TasksService) CreateTask(ctx context.Context, request *tasks.CreateTaskRequest) (*tasks.CreateTaskResponse, error) {
	return &tasks.CreateTaskResponse{}, nil
}
