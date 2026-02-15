# port in .env file

air:
	@echo "[ Air start... ]"
	air
	@echo "[ Final ]"

run:
	@echo "[ Runing... ]"
	@go run ./cmd/shortener/main.go
	@echo "[ Final ]"

# custom flags

run-p:
	@echo "[ Runing on costom port $(PORT)... ]"
	@go run ./cmd/shortener/main.go --port=$(PORT)
	@echo "[ Final ]"

run-l:
	@echo "[ Runing service with writing logs in file ]"
	@go run ./cmd/shortener/main.go --writingLogs
	@echo "[ Final ]"

run-p-l:
	@echo "[ Runing service on port $(PORT) with writing logs in file]"
	@go run ./cmd/shortener/main.go --port=$(PORT) --writingLogs
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

# utils

logger-clean:
	@echo "[ Logger clean... ]"
	@rm -rf ./logs/*
	@echo "[ Final ]"

clean-mod-cache:
	@echo "[ Clean mod cache... ]"
	@go clean -modcache
	@echo "[ Final ]"

swagger:
	swag init -g shortener/main.go -d ./cmd,./internal

# yndx container register

tag-register:
	docker tag shortener-service_prod:latest cr.yandex/${CODE_REGISTER}/shortener-service_prod:latest

push-register:
	docker push cr.yandex/${CODE_REGISTER}/shortener-service_prod:latest

# grafana

observe-run:
	docker run -d \
	    --name promtail \
	    --restart unless-stopped \
	    -v /var/log:/var/log:ro \
	    -v /var/lib/docker/containers:/var/lib/docker/containers:ro \
	    -v /var/lib/promtail:/var/lib/promtail \
	    -v $$(pwd):/etc/promtail:ro \
	    -v /var/run/docker.sock:/var/run/docker.sock \
	    --network host \
	    grafana/promtail:3.4.1 \
	    -config.file=/etc/promtail/promtail-config.yaml \
	    -log.level=info

