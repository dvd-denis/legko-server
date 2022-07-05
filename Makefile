.PHONY: build
build:
	go build -o build/ -v ./cmd/apiserver

.DEFAULT_GOAL := build