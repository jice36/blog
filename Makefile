.PHONY: build
build:
	go build -v ./cmd/blog/main.go
	go build -v ./cmd/db/db.go
.DEFAULT_GOAL := build