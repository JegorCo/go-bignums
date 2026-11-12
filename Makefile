.PHONY: all build test test-max test-chunks run

all: build run

build:
	mkdir -p bin
	go build -o bin/go-bignums .

test:
	go test ./...

test-max:
	go test -v -run '^TestMaximum' ./...

test-chunks:
	go test -v -run '^TestMaxChunks' ./...

run: build
	./bin/go-bignums
