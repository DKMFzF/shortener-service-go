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
