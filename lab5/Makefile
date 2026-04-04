.PHONY: run run-debug run-minsk build clean

run:
	go run ./cmd/linux/cli/main.go

run-debug:
	go run ./cmd/linux/cli/main.go -config ./config/config.yaml

run-minsk:
	go run ./cmd/linux/cli/main.go -config ./config/config_minsk.yaml

build:
	go build -o ./build/weather_app ./cmd/linux/cli/main.go

clean:
	rm -rf ./build/*

# CI/CD commands
ci-test:
	go test -v -race -cover ./...

ci-lint:
	golangci-lint run

ci-build:
	go build -o /dev/null ./cmd/linux/cli/main.go

ci-all: ci-lint ci-test ci-build

ci-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
