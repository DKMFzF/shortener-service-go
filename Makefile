.PHONY: help air run run-p run-l run-p-l run-p-l-d run-debug \
	build build-win build-clean \
	docker-b-dev docker-r-dev docker-b-watch docker-r-watch \
	docker-b-prod docker-r-prod up delete \
	test-all test-all-cache test-all-full-info \
	logger-clean clean-mod-cache swagger \
	tag-register push-register observe-run

APP_NAME      := shortener-service
CMD_DIR       := ./cmd/shortener
MAIN_PKG      := $(CMD_DIR)/main.go
BIN_DIR       := ./bin
BIN           := $(BIN_DIR)/main
BIN_WIN       := $(BIN_DIR)/main.exe

GO            ?= go
DOCKER        ?= docker
COMPOSE       ?= docker compose

PORT          ?= 8080
CODE_REGISTER ?= your-registry-id

IMAGE_BASE    := $(APP_NAME)
IMAGE_DEV     := $(IMAGE_BASE)_dev
IMAGE_WATCH   := $(IMAGE_BASE)_watch
IMAGE_PROD    := $(IMAGE_BASE)_prod

.DEFAULT_GOAL := help

# app run

air:
	@echo "[ Air start... ]"
	@air
	@echo "[ Final ]"

run:
	@echo "[ Running... ]"
	@$(GO) run $(MAIN_PKG)
	@echo "[ Final ]"

run-p:
	@echo "[ Running on custom port $(PORT)... ]"
	@$(GO) run $(MAIN_PKG) -p=$(PORT)
	@echo "[ Final ]"

run-l:
	@echo "[ Running service with writing logs in file ]"
	@$(GO) run $(MAIN_PKG) -l
	@echo "[ Final ]"

run-p-l:
	@echo "[ Running service on port $(PORT) with writing logs in file ]"
	@$(GO) run $(MAIN_PKG) --p=$(PORT) --writingLogs
	@echo "[ Final ]"

run-p-l-d:
	@echo "[ Running service on port $(PORT) with writing logs in file ]"
	@$(GO) run $(MAIN_PKG) -p=$(PORT) -l -dl
	@echo "[ Final ]"

run-debug:
	@echo "[ Running service on port $(PORT) with debug mode ]"
	@$(GO) run $(MAIN_PKG) -p=$(PORT) -l -dl
	@echo "[ Final]"

build: $(BIN)

$(BIN):
	@echo "[ Building... ]"
	@mkdir -p $(BIN_DIR)
	@$(GO) build -o $@ $(MAIN_PKG)
	@echo "[ Final ]"

build-win: $(BIN_WIN)

$(BIN_WIN):
	@echo "[ Building (windows)... ]"
	@mkdir -p $(BIN_DIR)
	@GOOS=windows GOARCH=amd64 $(GO) build -o $@ $(MAIN_PKG)
	@echo "[ Final ]"

build-clean:
	@echo "[ Builds file clean... ]"
	@rm -rf $(BIN_DIR)/*
	@echo "[ Final ]"

# docker

docker-b-dev:
	@echo "[ Docker build (dev)... ]"
	@$(DOCKER) build -t $(IMAGE_DEV) -f ./docker/Dockerfile.dev .
	@echo "[ Final ]"

docker-r-dev:
	@echo "[ Docker run (dev)... ]"
	@$(DOCKER) run -p $(PORT):8080 --name $(IMAGE_DEV) -d $(IMAGE_DEV)
	@echo "[ Final ]"

docker-b-watch:
	@echo "[ Docker build (watch)... ]"
	@$(DOCKER) build -t $(IMAGE_WATCH) -f ./docker/Dockerfile.watch .
	@echo "[ Final ]"

docker-r-watch:
	@echo "[ Docker run (watch)... ]"
	@$(DOCKER) run -p $(PORT):8080 --name $(IMAGE_WATCH) -d $(IMAGE_WATCH)
	@echo "[ Final ]"

docker-b-prod:
	@echo "[ Docker build (prod)... ]"
	@$(DOCKER) build -t $(IMAGE_PROD) -f ./docker/Dockerfile.prod .
	@echo "[ Final ]"

docker-r-prod:
	@echo "[ Docker run (prod)... ]"
	@$(DOCKER) run -p $(PORT):8080 --name $(IMAGE_PROD) -d $(IMAGE_PROD)
	@echo "[ Final ]"

# docker compose

up:
	@echo "[ Docker compose up... ]"
	@$(COMPOSE) up --build
	@echo "[ Docker compose down... ]"
	@$(COMPOSE) down
	@echo "[ Final ]"

delete:
	@echo "[ Docker compose down volumes ]"
	@$(COMPOSE) down --volumes
	@echo "[ Final ]"

# tests

test-all:
	@echo "[ Testing... ]"
	@$(GO) test -count=1 ./...
	@echo "[ Final ]"

test-all-cache:
	@echo "[ Testing + cache... ]"
	@$(GO) test ./...
	@echo "[ Final ]"

test-all-full-info:
	@echo "[ Testing full info... ]"
	@$(GO) test -v ./...
	@echo "[ Final ]"

# utils

logger-clean:
	@echo "[ Logger clean... ]"
	@rm -rf ./logs/*
	@echo "[ Final ]"

clean-mod-cache:
	@echo "[ Clean mod cache... ]"
	@$(GO) clean -modcache
	@echo "[ Final ]"

swagger:
	swag init -g shortener/main.go -d ./cmd,./internal

# yandex-cr

tag-register:
	@$(DOCKER) tag $(IMAGE_PROD):latest cr.yandex/$(CODE_REGISTER)/$(IMAGE_PROD):latest

push-register:
	@$(DOCKER) push cr.yandex/$(CODE_REGISTER)/$(IMAGE_PROD):latest

# promtail

observe-run:
	@$(DOCKER) run -d \
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

# help

help:
	@echo ""
	@echo "Scripts in Makefile"
	@echo "  make air              - run air (live reload)"
	@echo "  make run              - go run main"
	@echo "  make run-p PORT=8081  - run on custom port"
	@echo "  make run-l            - run with file logging"
	@echo "  make build            - build binary ($(BIN))"
	@echo "  make build-win        - build Windows binary ($(BIN_WIN))"
	@echo "  make test-all         - run tests (no cache)"
	@echo "  make docker-b-dev     - build dev image"
	@echo "  make docker-r-dev     - run dev container"
	@echo "  make up               - docker compose up+down"
	@echo "  make logger-clean     - clean logs"
	@echo ""