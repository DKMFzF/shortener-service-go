run:
	@echo "[ Runing... ]"
	@go run ./cmd/shortener/main.go
	@echo "[ Final ]"

build:
	@echo "[ Building... ]"
	@go build -o ./bin/main ./cmd/shortener/main.go
	@echo "[ Final ]"

build-win:
	@echo "[ Building... ]"
	@go build -o ./bin/main.exe ./cmd/shortener/main.go
	@echo "[ Final ]"

build-clean:
	@echo "[ Builds file clean... ]"
	@rm -rf ./bin/*
	@echo "[ Final ]"

docker-build:
	@echo "[ Docker build... ]"
	docker build -t shortener-service .
	@echo "[ Final ]"

docker-run:
	@echo "[ Docker start... ]"
	docker run -p 8080:8080 --name shortener-service -d shortener-service
	@echo "[ Final ]"

docker-compose-build:
	@echo "[ Docker compose build... ]"
	docker compose --build .
	@echo "[ Final ]"

docker-compose-up:
	@echo "[ Docker compose run... ]"
	docker compose up --build -d
	@echo "[ Final ]"

docker-compose-down:
	@echo "[ Docker compose down... ]"
	docker compose down
	@echo "[ Final ]"

test-all:
	@echo "[ Testing... ]"
	@go test count 1 ./...
	@echo "[ Final ]"

test-all-cache:
	@echo "[ Testing + cache... ]"
	@go test ./...
	@echo "[ Final ]"

test-all-full-info:
	@echo "[ Testing full info... ]"
	@go test -v ./...
	@echo "[ Final ]"

