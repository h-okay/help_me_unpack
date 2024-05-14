.PHONY: fmt lint vet build run clean ci
.DEFAULT_GOAL := build

fmt:
	go fmt ./...

lint: fmt
	golint ./...

vet: fmt
	go vet ./...
	shadow ./...

ci: fmt lint vet
	golangci-lint run ./...

build: ci
	go build -o bin .

run: build
	./bin; $(MAKE) clean
	
clean: 
	rm ./bin
