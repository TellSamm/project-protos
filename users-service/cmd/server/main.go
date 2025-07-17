package main

import (
	"log"

	"github.com/TellSamm/users-service/internal/database"
	transportgrpc "github.com/TellSamm/users-service/internal/transport/grpc"
	"github.com/TellSamm/users-service/internal/user"
)

func main() {

	database.InitDB()

	repo := user.NewUserRepository(database.DB)
	svc := user.NewUSerService(repo)

	if err := transportgrpc.RunGRPC(svc); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
