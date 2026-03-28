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
