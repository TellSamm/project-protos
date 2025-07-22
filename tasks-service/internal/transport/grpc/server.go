package grpc

import (
	"net"

	taskpb "github.com/TellSamm/project-protos/proto/task"
	userpb "github.com/TellSamm/project-protos/proto/user"
	"github.com/TellSamm/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.TaskService, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)

	return grpcSrv.Serve(lis)
}
