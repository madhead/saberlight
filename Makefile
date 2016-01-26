BINARY=saberlight
VERSION=$(shell git rev-parse --short HEAD)

build:
	cd app && go build -o ${BINARY} -ldflags "-X github.com/madhead/saberlight/app/commands.version=${VERSION}" saberlight.go

test:
	cd app && go test $(shell go list ./... | grep -v vendor)
