# Сборка protobuf
PROTOS := $(shell find proto -name "*.proto")
OUT_DIR := .

generate:
	@protoc \
	  --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	  $(PROTOS)

clean:
	find . -name "*.pb.go" -delete

# ===============================
# MIGRATIONS!!!!
# ===============================

# Путь до бинарника migrate
DB_URL=postgres://postgres:postgres@localhost:5436/protodb?sslmode=disable
MIGRATE=migrate -path ./migrations -database "$(DB_URL)"

# Создание новой миграции: make migrate-new NAME=create_users
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# ===============================
# APP!!!!!
# ===============================

# Запуск приложения
run:
	go run cmd/server/main.go