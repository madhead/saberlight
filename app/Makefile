BINARY=saberlight
VERSION=$(shell git rev-parse --short HEAD)

build:
	go build -o ${BINARY} -ldflags "-X github.com/madhead/saberlight/app/commands.version=${VERSION}"
