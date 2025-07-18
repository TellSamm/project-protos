package grpc

import (
	"fmt"
	"net"

	userpb "github.com/TellSamm/project-protos/proto/user"
	"github.com/TellSamm/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.UserService) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("не удалось начать прослушивание: %w", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	fmt.Println("gRPC сервер запущен на порту :50051")
	return grpcServer.Serve(lis)
}
