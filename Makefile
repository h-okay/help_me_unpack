.PHONY: test cover fmt lint vet build run clean ci
.DEFAULT_GOAL := build

test:
	go test ./test

cover:
	go test -coverpkg=./src -coverprofile=c.out ./test 
	go tool cover -func=c.out
	rm c.out

fmt:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...
	shadow ./...

ci: fmt lint vet
	golangci-lint run ./...

build: ci
	go build -o bin .

run: build
	./bin
	
clean: 
	rm ./bin
