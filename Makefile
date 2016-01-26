BINARY=saberlight
VERSION=$(shell git rev-parse --short HEAD)

build:
	go build -o app/${BINARY} -ldflags "-X github.com/madhead/saberlight/app/commands.version=${VERSION}" app/saberlight.go
