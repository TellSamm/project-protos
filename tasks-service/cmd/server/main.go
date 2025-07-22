package server

import (
	"log"

	"github.com/TellSamm/tasks-service/internal/database"
	"github.com/TellSamm/tasks-service/internal/task"
	transportgrpc "github.com/TellSamm/tasks-service/internal/transport/grpc"
)

func main() {
	//Инициализация БД
	database.InitDB()

	//Репозиторий и сервис задач
	repo := task.NewTaskRepository(database.DB)
	svc := task.NewTaskService(repo)

	//Клиент к Users-сервису
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	//Запуск gRPC Tasks-сервиса
	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
