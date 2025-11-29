# port in .env file
air:
	@echo "[ Air start... ]"
	air
	@echo "[ Final ]"

run:
	@echo "[ Runing... ]"
	@go run ./cmd/shortener/main.go
	@echo "[ Final ]"

# custom port

run-p:
	@echo "[ Runing on costom port $(PORT)... ]"
	@go run ./cmd/shortener/main.go --port=$(PORT)
	@echo "[ Final ]"

# build

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

# docker-dev

docker-b-dev:
	@echo "[ Docker build... ]"
	docker build -t shortener-service_dev -f ./docker/Dockerfile.dev .
	@echo "[ Final ]"

docker-r-dev:
	@echo "[ Docker start... ]"
	docker run -p 8080:8080 --name shortener-service_dev -d shortener-service_dev
	@echo "[ Final ]"

# docker-watch

docker-b-watch:
	@echo "[ Docker build... ]"
	docker build -t shortener-service_watch -f ./docker/Dockerfile.watch . 
	@echo "[ Final ]"

docker-r-watch:
	@echo "[ Docker run... ]"
	docker run -p 8080:8080 --name shortener-service_watch -d shortener-service_watch
	@echo "[ Final ]"

# docker-prod

docker-b-prod:
	@echo "[ Docker build... ]"
	docker build -t shortener-service_prod -f ./docker/Dockerfile.prod . 
	@echo "[ Final ]"

docker-r-prod:
	@echo "[ Docker run... ]"
	docker run -p 8080:8080 --name shortener-service_prod -d shortener-service_prod
	@echo "[ Final ]"

# docker compose

up:
	@echo "[ Docker run... ]"
	docker compose up --build
	@echo "[ Docker compose down... ]"
	docker compose down
	@echo "[ Final ]"

delete:
	@echo "[ Docker compose down volumes ]"
	docker compose down --volumes
	@echo "[ Final ]"

# tests

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

logger-clean:
	@echo "[ Logger clean... ]"
	@rm -rf ./logs/*
	@echo "[ Final ]"

clean-mod-cache:
	@echo "[ Clean mod cache... ]"
	@go clean -modcache
	@echo "[ Final ]"

