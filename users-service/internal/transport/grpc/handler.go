package grpc

import (
	"context"

	userpb "github.com/TellSamm/project-protos/proto/user"
	"github.com/TellSamm/users-service/internal/models"
	"github.com/TellSamm/users-service/internal/user"
)

type Handler struct {
	svc user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	modelUser := &models.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	if err := h.svc.CreateUser(modelUser); err != nil {
		return nil, err
	}

	protoUser := &userpb.User{
		Id:    modelUser.ID.String(),
		Email: modelUser.Email,
	}

	resp := &userpb.CreateUserResponse{
		User: protoUser,
	}

	return resp, nil
}

func (h *Handler) ListUsers(ctx context.Context, _ *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	resp := &userpb.ListUsersResponse{}
	for _, u := range users {
		protoUser := &userpb.User{
			Id:    u.ID.String(),
			Email: u.Email,
		}
		resp.Users = append(resp.Users, protoUser)
	}

	return resp, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	updateData := &models.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	id := req.GetId()

	dbUser, err := h.svc.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	dbUser.Email = updateData.Email
	dbUser.Password = updateData.Password

	if err := h.svc.UpdateUser(dbUser); err != nil {
		return nil, err
	}

	protoUser := &userpb.User{
		Id:    dbUser.ID.String(),
		Email: dbUser.Email,
	}

	resp := &userpb.UpdateUserResponse{
		User: protoUser,
	}

	return resp, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id := req.GetId()

	userModel, err := h.svc.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	protoUser := &userpb.User{
		Id:    userModel.ID.String(),
		Email: userModel.Email,
	}

	return &userpb.GetUserResponse{
		User: protoUser,
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.GetId()

	if err := h.svc.DeleteUserByID(id); err != nil {
		return nil, err
	}

	resp := &userpb.DeleteUserResponse{}
	return resp, nil
}
