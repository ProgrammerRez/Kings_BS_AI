SRC_DIR = ./main

build:
	go build -o bin/app $(SRC_DIR)

run: build
	./bin/app

test:
	go test ./... -v