.PHONY: build

APP_NAME = dbTable2Struct

all: build

build:
	make build_linux

build_linux:
	mkdir -p build/${APP_NAME}_linux
	go mod tidy -compat=1.17
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o ./build/${APP_NAME}_linux/dbTable2Struct main.go
	cp ./conf.toml ./build/${APP_NAME}_linux/
	cp -r ./template ./build/${APP_NAME}_linux/