.PHONY:deploy

TARGET=giniladmin

GOOS=
GOARCH=amd64
SWAG=swag

CGO_ENABLED=0
CC=x86_64-linux-musl-gcc
CXX=x86_64-linux-musl-g++

build: swag
	mkdir -p bin
	go mod tidy
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) CGO_ENABLED=$(CGO_ENABLED) CC=$(CC) CXX=$(CXX) go build -gcflags "all=-N -l" -o bin/$(TARGET) cmd/main.go

swag:
	$(SWAG) init -g cmd/main.go