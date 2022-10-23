all: dependencies format mocking test
dependencies:
	@echo "Sync dependencies with go mod tidy"
	@go mod tidy
format:
	@echo "Formatting Go code recursively"
	@go fmt ./...
mocking:
	@echo "Generatting mocks recursively"
	@go generate ./...
test:
	@echo "Making test coverage report..."
	@go test ./... -coverprofile=coverage.out -race
	@go tool cover -html=coverage.out
