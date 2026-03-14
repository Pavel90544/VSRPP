.PHONY: run run-debug run-minsk build clean

run:
	go run ./cmd/linux/cli/main.go

run-debug:
	go run ./cmd/linux/cli/main.go -debug=true

run-minsk:
	go run ./cmd/linux/cli/main.go -lat=53.9045 -lon=27.5615

build:
	go build -o ./build/weather_app ./cmd/linux/cli/main.go

clean:
	rm -rf ./build/*
