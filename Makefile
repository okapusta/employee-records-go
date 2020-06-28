all: deps build

deps:
	dep ensure

build:
	go build -o bin/employees ./cmd/employees

test:
	go test ./...

container:
	docker build . -t employees
