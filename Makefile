run:
	@go run main.go

build:
	@go build main.go

start-api: build
	./main