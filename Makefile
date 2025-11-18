run:
	@echo "[ Runing... ]"
	@go run ./cmd/shortener/main.go
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
