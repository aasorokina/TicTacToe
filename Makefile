SHELL = /bin/bash

all: init
.PHONY: all run clean fmt lint init

run:
	@go run cmd/main.go

init: clean
	@go mod init tictactoe
	@go mod tidy

clean:
	@rm -rf go.mod go.sum

fmt:
	@go fmt ./...

lint:
	@golint ./...
