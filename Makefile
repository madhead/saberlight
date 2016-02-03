BINARY=saberlight
VERSION=$(shell printf "r%s.%s" "$(shell git rev-list --count HEAD)" "$(shell git rev-parse --short HEAD)")
SRC_DIR=app

build:
	cd ${SRC_DIR} && go build -o ${BINARY} -ldflags "-X github.com/madhead/saberlight/app/commands.version=${VERSION}" saberlight.go

test:
	cd ${SRC_DIR} && go test $(shell go list ./... | grep -v vendor)
