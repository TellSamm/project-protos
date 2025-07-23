package grpc

import (
	"context"
	"fmt"
	taskpb "github.com/TellSamm/project-protos/proto/task"
	userpb "github.com/TellSamm/project-protos/proto/user"
	"github.com/TellSamm/tasks-service/internal/models"
	"github.com/TellSamm/tasks-service/internal/task"
	"github.com/google/uuid"
)

type Handler struct {
	svc        task.TaskService
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.TaskService, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %s not found: %w", req.UserId, err)
	}

	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	taskModel := &models.Task{
		UserID: userUUID,
		Title:  req.Title,
		IsDone: false,
	}

	if err := h.svc.CreateTask(taskModel); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &taskpb.CreateTaskResponse{
		Task: &taskpb.Task{
			Id:     taskModel.ID,
			UserId: taskModel.UserID.String(),
			Title:  taskModel.Title,
			IsDone: taskModel.IsDone,
		},
	}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	taskModel, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &taskpb.GetTaskResponse{
		Task: &taskpb.Task{
			Id:     taskModel.ID,
			UserId: taskModel.UserID.String(),
			Title:  taskModel.Title,
			IsDone: taskModel.IsDone,
		},
	}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{
			Id:     t.ID,
			UserId: t.UserID.String(),
			Title:  t.Title,
			IsDone: t.IsDone,
		})
	}

	return &taskpb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	allTasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	var pbTasks []*taskpb.Task
	for _, t := range allTasks {
		if t.UserID == userID {
			pbTasks = append(pbTasks, &taskpb.Task{
				Id:     t.ID,
				UserId: t.UserID.String(),
				Title:  t.Title,
				IsDone: t.IsDone,
			})
		}
	}

	return &taskpb.ListTasksByUserResponse{Tasks: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.Task.UserId}); err != nil {
		return nil, fmt.Errorf("user %s not found: %w", req.Task.UserId, err)
	}

	userUUID, err := uuid.Parse(req.Task.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	taskModel := &models.Task{
		ID:     req.Task.Id,
		UserID: userUUID,
		Title:  req.Task.Title,
		IsDone: req.Task.IsDone,
	}

	if err := h.svc.UpdateTask(taskModel); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return &taskpb.UpdateTaskResponse{
		Task: &taskpb.Task{
			Id:     taskModel.ID,
			UserId: taskModel.UserID.String(),
			Title:  taskModel.Title,
			IsDone: taskModel.IsDone,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTaskByID(req.Id); err != nil {
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}

	return &taskpb.DeleteTaskResponse{}, nil
}
